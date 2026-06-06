package service

import (
	"fmt"
	"strings"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"github.com/globalxtreme/go-identifier/data"
	gxstorage "github.com/globalxtreme/go-storage/v2"
	"gorm.io/gorm"

	"service/internal/customer/repository"
	"service/internal/pkg/activity"
	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	error2 "service/internal/pkg/error"
	form2 "service/internal/pkg/form"
	"service/internal/pkg/model"
	"service/internal/pkg/parser"
	"service/internal/pkg/port"
)

type CustomerService interface {
	SetTransaction(tx *gorm.DB)
	SetEmployeeIdentifier(employee data.EmployeeIdentifierData)

	Create(form form2.CustomerForm) model.Customer
	Update(uuid string, form form2.CustomerForm) model.Customer
	UpdateStatus(uuid string, form form2.CustomerStatusForm) model.Customer
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
}

func (srv *customerService) SetTransaction(tx *gorm.DB) {
	srv.tx = tx
}

func (srv *customerService) SetEmployeeIdentifier(employee data.EmployeeIdentifierData) {
	srv.employee = employee
}

func (srv *customerService) Create(form form2.CustomerForm) model.Customer {
	customer := srv.prepare(nil)
	if srv.checkIDandSIMNumberExist(form.IDNumber, form.SIMNumber) {
		error2.ErrXtremeCustomerSave("Your ID or SIM Number has been registered")
	}

	form.Phone = srv.adjustmentPhone(form.Phone)

	uploadIDFile := srv.uploadIDPhoto(form)
	uploadID := xtrememodel.MapInterfaceColumn(uploadIDFile)

	var err error
	if file, ok := (uploadID)["file"].(string); ok {
		defer srv.rollbackStorage(&err, file)
	}

	err = config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)

		customer = srv.repository.Create(form, &uploadID)

		parser := parser.CustomerParser{Object: customer}
		activity.UseActivity{Employee: srv.employee}.SetReference(&customer).SetParser(&parser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Customer: %s [%d]", customer.Name, customer.ID))

		return nil
	})

	return customer
}

func (srv *customerService) Update(uuid string, form form2.CustomerForm) model.Customer {
	customer := srv.prepare(&uuid)
	parser := parser.CustomerParser{Object: customer}

	if form.IDNumber != customer.IDNumber || form.SIMNumber != customer.SIMNumber {
		if srv.checkIDandSIMNumberExist(form.IDNumber, form.SIMNumber) {
			error2.ErrXtremeCustomerUpdate("Your ID or SIM Number has been registered")
		}
	}

	form.Phone = srv.adjustmentPhone(form.Phone)

	var (
		err            error
		uploadID       *xtrememodel.MapInterfaceColumn
		oldIDPhotoPath string
	)
	if form.IDPhoto != nil {
		temp := xtrememodel.MapInterfaceColumn(srv.uploadIDPhoto(form))
		uploadID = &temp

		if customer.IDPhoto != nil {
			if file, ok := (*customer.IDPhoto)["file"].(string); ok {
				oldIDPhotoPath = file
			}
		}

		defer srv.rollbackStorage(&err, temp["file"].(string))
	}

	err = config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)
		useActivity := activity.UseActivity{Employee: srv.employee}.SetReference(&customer).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		customer = srv.repository.Update(customer, form, (*xtrememodel.MapInterfaceColumn)(uploadID))

		parser.Object = customer
		useActivity.SetReference(&customer).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Update Customer: %s [%d]", customer.Name, customer.ID))
		return nil
	})
	if err == nil && oldIDPhotoPath != "" {
		srv.deleteIDPhoto(oldIDPhotoPath)
	}
	return customer
}

func (srv *customerService) UpdateStatus(uuid string, form form2.CustomerStatusForm) model.Customer {
	srv.repository = repository.NewCustomerRepository()
	customer := srv.prepare(&uuid)

	parser := parser.CustomerParser{Object: customer}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)
		useActivity := activity.UseActivity{Employee: srv.employee}.SetReference(&customer).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		customer := srv.repository.UpdateStatus(customer, form)

		parser.Object = customer
		useActivity.SetReference(&customer).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Update Customer status: %s [%d]", customer.Name, customer.ID))
		return nil
	})
	return customer
}

func (srv *customerService) Delete(uuid string) {
	customer := srv.prepare(&uuid)
	var oldIDPhotoPath string
	if file, ok := (*customer.IDPhoto)["file"].(string); ok {
		oldIDPhotoPath = file
	}

	err := config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)
		srv.repository.Delete(customer)

		activity.UseActivity{Employee: srv.employee, Action: constant.ACTION_DELETE}.SetReference(customer).
			Save(fmt.Sprintf("Delete customer %s [%d]", customer.Name, customer.ID))
		return nil
	})
	if err == nil && oldIDPhotoPath != "" {
		fmt.Println(oldIDPhotoPath)
		srv.deleteIDPhoto(oldIDPhotoPath)
	}
}

/** --- UNEXPORTED FUNCTIONS --- */

func (srv *customerService) prepare(uuid *string) model.Customer {
	srv.repository = repository.NewCustomerRepository()

	var customer model.Customer
	if uuid != nil {
		customer = srv.repository.FirstByForm(form2.CustomerFilterForm{UUID: *uuid})
	}
	return customer
}

func (srv *customerService) checkIDandSIMNumberExist(IDNumber, SIMNumber string) bool {
	count := srv.repository.CountIDandSIMNumber(form2.CustomerFilterForm{IDNumber: IDNumber, SIMNumber: SIMNumber})
	if count > 0 {
		return true
	}
	return false
}

func (srv *customerService) adjustmentPhone(phone string) string {
	if strings.HasPrefix(phone, "0") {
		return "62" + phone[1:]
	}
	return phone
}

func (srv *customerService) uploadIDPhoto(form form2.CustomerForm) map[string]interface{} {
	file, fileHandler, err := form.Request.FormFile("IDPhoto[file]")
	if err != nil {
		error2.ErrXtremeFileUpload(err.Error())
	}

	attachment := make(map[string]interface{})

	upload, err := gxstorage.UploadFile(gxstorage.PublicStorageUpload{
		File:          file,
		Path:          constant.PathImageCustomerID(),
		Name:          fileHandler.Filename,
		MimeType:      form.IDPhoto.MimeType,
		CreatedBy:     srv.employee.ID,
		CreatedByName: srv.employee.FullName,
	})

	if err != nil {
		fmt.Println(err.Error())
		error2.ErrXtremeFileUpload(err.Error())
	}

	fullPath := upload.GetResult().GetFullPath()

	attachment = map[string]interface{}{
		"file":     fullPath,
		"mimeType": form.IDPhoto.MimeType,
	}

	return attachment
}

func (srv *customerService) deleteIDPhoto(path string) {
	_, err := gxstorage.Delete(path)
	if err != nil {
		error2.ErrXtremeFileDelete(err.Error())
	}
}

func (srv *customerService) rollbackStorage(err *error, customerIDPhotoPath string) {
	if r := recover(); r != nil {
		srv.deleteIDPhoto(customerIDPhotoPath)
		panic(r)
	}
	if err != nil && *err != nil {
		srv.deleteIDPhoto(customerIDPhotoPath)
	}
}
