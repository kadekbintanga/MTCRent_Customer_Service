package service

import (
	"fmt"
	"strings"

	"github.com/globalxtreme/go-identifier/data"
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
	SetActivityRepository(repo port.ActivityRepository)

	Create(form form2.CustomerForm) model.Customer
	Update(filterForm form2.CustomerFilterForm, form form2.CustomerForm) model.Customer
	UpdateStatus(filterForm form2.CustomerFilterForm, form form2.CustomerStatusForm) model.Customer
	Delete(filterForm form2.CustomerFilterForm)
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

func (srv *customerService) SetActivityRepository(repo port.ActivityRepository) {
	srv.activityRepo = repo
}

func (srv *customerService) Create(form form2.CustomerForm) model.Customer {
	srv.repository = repository.NewCustomerRepository()
	checkIdSimNumber := srv.repository.FindByForm(form2.CustomerFilterForm{IDNumber: form.IDNumber, SIMNumber: form.SIMNumber})
	if len(checkIdSimNumber) > 0 {
		error2.ErrXtremeCustomerSave("Your ID or SIM Number has been registered")
	}
	if strings.HasPrefix(form.Phone, "0") {
		form.Phone = "62" + form.Phone[1:]
	}

	var customer model.Customer

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository = repository.NewCustomerRepository(tx)

		customer = srv.repository.Create(form)

		parser := parser.CustomerParser{Object: customer}
		activity.UseActivity{Employee: srv.employee}.SetReference(&customer).SetParser(&parser).SetNewProperty(constant.ACTION_CREATE).
			Save(fmt.Sprintf("Create new Customer: %s [%d]", customer.Name, customer.ID))

		return nil
	})

	return customer

}

func (srv *customerService) Update(filterForm form2.CustomerFilterForm, form form2.CustomerForm) model.Customer {
	srv.repository = repository.NewCustomerRepository()
	customer := srv.repository.FirstByForm(filterForm)

	parser := parser.CustomerParser{Object: customer}

	if strings.HasPrefix(form.Phone, "0") {
		form.Phone = "62" + form.Phone[1:]
	}

	if form.StatusId == 0 {
		form.StatusId = customer.StatusId
	}

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)
		useActivity := activity.UseActivity{Employee: srv.employee}.SetReference(&customer).SetParser(&parser).SetOldProperty(constant.ACTION_UPDATE)

		customer = srv.repository.Update(customer, form)

		parser.Object = customer
		useActivity.SetReference(&customer).SetParser(&parser).SetNewProperty(constant.ACTION_UPDATE).
			Save(fmt.Sprintf("Update Customer: %s [%d]", customer.Name, customer.ID))
		return nil
	})
	return customer
}

func (srv *customerService) UpdateStatus(filterForm form2.CustomerFilterForm, form form2.CustomerStatusForm) model.Customer {
	srv.repository = repository.NewCustomerRepository()
	customer := srv.repository.FirstByForm(filterForm)

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

func (srv *customerService) Delete(filterForm form2.CustomerFilterForm) {
	srv.repository = repository.NewCustomerRepository()
	customer := srv.repository.FirstByForm(filterForm)

	config.PgSQL.Transaction(func(tx *gorm.DB) error {
		srv.repository.SetTransaction(tx)

		parser := parser.CustomerParser{Object: customer}
		srv.repository.Delete(customer)

		activity.UseActivity{Employee: srv.employee}.SetReference(customer).SetParser(&parser).SetOldProperty(constant.ACTION_DELETE).
			Save(fmt.Sprintf("Delete customer %s [%d]", customer.Name, customer.ID))
		return nil
	})
}
