package repository

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	"gorm.io/gorm"

	"service/internal/pkg/config"
	"service/internal/pkg/constant"
	"service/internal/pkg/core"
	error2 "service/internal/pkg/error"
	"service/internal/pkg/form"
	"service/internal/pkg/model"
)

type CustomerRepository interface {
	core.TransactionInterface
	core.FirstRepository[form.CustomerFilterForm, model.Customer]
	core.FindRepository[form.CustomerFilterForm, model.Customer]
	core.PaginateRepository[form.CustomerFilterForm, model.Customer]

	Create(form form.CustomerForm, file *xtrememodel.MapInterfaceColumn) model.Customer
	Update(customer model.Customer, form form.CustomerForm, file *xtrememodel.MapInterfaceColumn) model.Customer
	UpdateStatus(customer model.Customer, form form.CustomerStatusForm) model.Customer
	Delete(customer model.Customer)
	CountIDandSIMNumber(form form.CustomerFilterForm) int64
}

func NewCustomerRepository(args ...*gorm.DB) CustomerRepository {
	repository := customerRepository{}
	if len(args) > 0 {
		repository.Transaction = args[0]
	}

	return &repository
}

type customerRepository struct {
	Transaction *gorm.DB
}

func (repo *customerRepository) SetTransaction(tx *gorm.DB) {
	repo.Transaction = tx
}

func (repo *customerRepository) FirstByForm(form form.CustomerFilterForm, args ...func(query *gorm.DB) *gorm.DB) model.Customer {
	query := repo.prepareAndFilter(form)

	if len(args) > 0 {
		query = args[0](query)
	}

	var customer model.Customer
	err := query.First(&customer).Error
	if err != nil {
		error2.ErrXtremeCustomerGet(err.Error())
	}

	return customer
}

func (repo *customerRepository) FindByForm(form form.CustomerFilterForm) []model.Customer {
	query := repo.prepareAndFilter(form)

	var customers []model.Customer
	err := query.Find(&customers).Error
	if err != nil {
		error2.ErrXtremeCustomerGet(err.Error())
	}

	return customers
}

func (repo *customerRepository) PaginateByForm(form form.CustomerFilterForm) ([]model.Customer, interface{}) {
	parameter := url.Values{}
	parameter.Set("page", strconv.Itoa(form.Page))
	parameter.Set("limit", strconv.Itoa(form.Limit))

	query := repo.prepareAndFilter(form)
	customers, pagination, err := xtrememodel.Paginate(query, parameter, model.Customer{})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		error2.ErrXtremeCustomerGet(err.Error())
	}

	return customers, pagination
}

func (repo *customerRepository) Create(form form.CustomerForm, file *xtrememodel.MapInterfaceColumn) model.Customer {
	customer := model.Customer{
		Name:      form.Name,
		IDNumber:  form.IDNumber,
		SIMNumber: form.SIMNumber,
		Phone:     form.Phone,
		Address:   form.Address,
		StatusId:  constant.CUSTOMER_STATUS_ACTIVE_ID,
	}

	if file != nil {
		customer.IDPhoto = file
	}

	err := repo.Transaction.Create(&customer).Error
	if err != nil {
		error2.ErrXtremeCustomerSave(err.Error())
	}

	return customer
}

func (repo *customerRepository) Update(customer model.Customer, form form.CustomerForm, file *xtrememodel.MapInterfaceColumn) model.Customer {
	customer.Name = form.Name
	customer.IDNumber = form.IDNumber
	customer.SIMNumber = form.SIMNumber
	customer.Phone = form.Phone
	customer.Address = form.Address
	customer.StatusId = form.StatusId
	customer.BlacklistReason = form.BlacklistReason

	if file != nil {
		customer.IDPhoto = file
	}

	err := repo.Transaction.Updates(&customer).Error
	if err != nil {
		error2.ErrXtremeCustomerUpdate(err.Error())
	}
	return customer
}

func (repo *customerRepository) UpdateStatus(customer model.Customer, form form.CustomerStatusForm) model.Customer {
	err := repo.Transaction.Model(&customer).Update("status", form.StatusId).Error
	if err != nil {
		error2.ErrXtremeCustomerUpdate(err.Error())
	}

	return customer
}

func (repo *customerRepository) Delete(customer model.Customer) {
	err := repo.Transaction.Delete(&customer).Error
	if err != nil {
		error2.ErrXtremeCustomerDelete(err.Error())
	}
}

func (repo *customerRepository) CountIDandSIMNumber(form form.CustomerFilterForm) int64 {
	var count int64

	err := config.PgSQL.Model(&model.Customer{}).Where(`customers."IDNumber" = ? OR customers."SIMNumber" = ?`, form.IDNumber, form.SIMNumber).Count(&count).Error
	if err != nil {
		error2.ErrXtremeCustomerGet(err.Error())
	}

	return count

}

/** --- UNEXPORTED FUNCTIONS --- */

func (repo *customerRepository) prepareAndFilter(form form.CustomerFilterForm) *gorm.DB {
	query := config.PgSQL

	if form.ID > 0 {
		query = query.Where("customers.id = ?", form.ID)
	}

	if form.UUID != "" {
		query = query.Where("customers.uuid = ?", form.UUID)
	}

	if form.IDNumber != "" {
		query = query.Where(`customers."IDNumber" = ?`, form.IDNumber)
	}

	if form.SIMNumber != "" {
		query = query.Where(`customers."SIMNumber" = ?`, form.SIMNumber)
	}

	if search := form.Search; len(search) > 3 {
		searchVal := "%" + search + "%"
		query = query.Where(`customers."name" ILIKE ?`, searchVal)
	}

	if len(form.Orders) > 0 {
		for key, value := range form.Orders {
			query = query.Order(fmt.Sprintf("%s %s", key, value))
		}
	} else {
		query = query.Order("id DESC")
	}

	return query
}
