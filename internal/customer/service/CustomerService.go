package service

import (
	"fmt"
	"net/http"

	"github.com/globalxtreme/go-identifier/data"
	gxstorage "github.com/globalxtreme/go-storage/v2"
	"gorm.io/gorm"

	"service/internal/customer/repository"
	"service/internal/pkg/activity"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	"service/internal/pkg/core"
	error2 "service/internal/pkg/error"
	form2 "service/internal/pkg/form"
	"service/internal/pkg/model"
	"service/internal/pkg/parser"
	"service/internal/pkg/port"
	"service/internal/pkg/saga"
)

type CustomerService interface {
	SetTransaction(tx *gorm.DB)
	SetEmployeeIdentifier(employee data.EmployeeIdentifierData)

	Create(form form2.CustomerForm) model.Customer
	Update(uuid string, form form2.CustomerForm) model.Customer
	UpdateStatus(uuid string, form form2.CustomerStatusForm) (model.Customer, model.Customer)
	Delete(uuid string)
}

func NewCustomerService() CustomerService {
	return &customerService{}
}

type customerService struct {
	tx           *gorm.DB
	repository   repository.CustomerRepository
	activityRepo port.ActivityRepository
	employee     data.EmployeeIdentifierData
	saga         saga.StorageSaga
}

func (srv *customerService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *customerService) SetEmployeeIdentifier(employee data.EmployeeIdentifierData) {
	srv.employee = employee
}

func (srv *customerService) Create(form form2.CustomerForm) model.Customer {
	srv.saga = saga.StorageSaga{}
	defer srv.saga.Close()

	customer := srv.prepare(nil, &form)
	srv.validateData(nil, form)

	uploadIdetityFhoto := srv.uploadIdentityPhoto(form)

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)

		customer = srv.repository.Create(form, &uploadIdetityFhoto)

		parser := parser.CustomerParser{Object: customer}
		activity.UseActivity{Employee: srv.employee}.SetReference(&customer).SetParser(&parser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Customer: %s [%d]", customer.Name, customer.ID))

		return nil
	})

	return customer
}

func (srv *customerService) Update(uuid string, form form2.CustomerForm) model.Customer {
	srv.saga = saga.StorageSaga{}
	defer srv.saga.Close()

	customer := srv.prepare(&uuid, &form)
	srv.validateData(&customer, form)

	parser := parser.CustomerParser{Object: customer}

	identityPhoto := srv.uploadIdentityPhoto(form)
	if identityPhoto != nil && customer.IdentityPhoto != nil {
		if file, ok := (*customer.IdentityPhoto)["file"].(string); ok {
			srv.saga.DeleteStoragePaths = append(srv.saga.DeleteStoragePaths, file)
		}
	}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)

		useActivity := activity.UseActivity{Employee: srv.employee}.SetReference(&customer).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		customer = srv.repository.Update(customer, form, &identityPhoto)

		parser.Object = customer
		useActivity.SetReference(&customer).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Update Customer: %s [%d]", customer.Name, customer.ID))

		return nil
	})

	return customer
}

func (srv *customerService) UpdateStatus(uuid string, form form2.CustomerStatusForm) (model.Customer, model.Customer) {
	srv.repository = repository.NewCustomerRepository()
	customer := srv.prepare(&uuid, nil)

	parser := parser.CustomerParser{Object: customer}
	oldCustomer := customer

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)

		useActivity := activity.UseActivity{Employee: srv.employee}.SetReference(&customer).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE, constant.ACTIVITY_CUSTOMER_STATUS)

		customer = srv.repository.UpdateStatus(customer, form)

		parser.Object = customer
		useActivity.SetReference(&customer).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE, constant.ACTIVITY_CUSTOMER_STATUS).
			Save(fmt.Sprintf("Update Customer status: %s [%d]", customer.Name, customer.ID))

		return nil
	})
	return customer, oldCustomer
}

func (srv *customerService) Delete(uuid string) {
	srv.saga = saga.StorageSaga{}
	defer srv.saga.Close()
	customer := srv.prepare(&uuid, nil)
	if file, ok := (*customer.IdentityPhoto)["file"].(string); ok {
		srv.saga.DeleteStoragePaths = append(srv.saga.DeleteStoragePaths, file)
	}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)
		srv.repository.Delete(customer)

		activity.UseActivity{Employee: srv.employee, Action: constant.ACTION_DELETE}.SetReference(customer).
			Save(fmt.Sprintf("Delete customer %s [%d]", customer.Name, customer.ID))
		return nil
	})
}

/** --- UNEXPORTED FUNCTIONS --- */

func (srv *customerService) prepare(uuid *string, form *form2.CustomerForm) model.Customer {
	srv.repository = repository.NewCustomerRepository()
	srv.repository.SetEmployeeIdentifier(srv.employee)

	var customer model.Customer
	if uuid != nil {
		customer = srv.repository.FirstByForm(form2.CustomerFilterForm{UUID: *uuid})
	}

	if form != nil {
		if form.Phone != "" {
			form.Phone = core.AdjustmentPhone(form.Phone)
		}
	}
	return customer
}

func (srv *customerService) validateData(customer *model.Customer, form form2.CustomerForm) {
	var customerId uint
	if customer != nil {
		customerId = customer.ID
	}

	srv.repository.CheckDuplicateIDorSIMNumber(form2.CustomerFilterForm{IDNumber: form.IDNumber, SIMNumber: form.SIMNumber, ID: customerId})
}

func (srv *customerService) uploadIdentityPhoto(form form2.CustomerForm) map[string]interface{} {
	file, fileHandler, err := form.Request.FormFile("identityPhoto[file]")
	if err != nil {
		if err == http.ErrMissingFile {
			return nil
		}
		error2.ErrXtremeFileUpload(err.Error())
	}

	defer file.Close()

	attachment := make(map[string]interface{})

	upload, err := gxstorage.UploadFile(gxstorage.PublicStorageUpload{
		File:          file,
		Path:          constant.PathImageCustomerID(),
		Name:          fileHandler.Filename,
		MimeType:      form.IdentityPhoto.MimeType,
		CreatedBy:     srv.employee.ID,
		CreatedByName: srv.employee.FullName,
	})

	if err != nil {
		error2.ErrXtremeFileUpload(err.Error())
	}

	fullPath := upload.GetResult().GetFullPath()

	attachment = map[string]interface{}{
		"file":     fullPath,
		"mimeType": form.IdentityPhoto.MimeType,
	}

	srv.saga.StoragePaths = append(srv.saga.StoragePaths, fullPath)

	return attachment
}
