package repository

import (
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"gorm.io/gorm"
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	error2 "service/internal/pkg/error"
	"service/internal/pkg/model"
	"service/internal/pkg/request"
)

/** --- INTERFACE --- */

type TestingRepository interface {
	core.TransactionRepository

	FirstById(id string, args ...func(query *gorm.DB) *gorm.DB) model.Testing
	Find(parameter url.Values) ([]model.Testing, interface{}, error)

	Store(request request.TestingRequest) model.Testing
	Delete(testing model.Testing)

	AddSub(testing model.Testing, sub string) model.TestingSub
	DeleteSub(testingSub model.TestingSub)
}

func NewTestingRepository(args ...*gorm.DB) TestingRepository {
	repository := testingRepository{}
	if len(args) > 0 {
		repository.transaction = args[0]
	}

	return &repository
}

/** --- MAIN REPOSITORY --- */

type testingRepository struct {
	transaction *gorm.DB
}

func (repo *testingRepository) SetTransaction(tx *gorm.DB) {
	repo.transaction = tx
}

func (repo *testingRepository) FirstById(id string, args ...func(query *gorm.DB) *gorm.DB) model.Testing {
	var testing model.Testing

	query := config.PgSQL
	if len(args) > 0 {
		query = args[0](query)
	}

	err := query.First(&testing, "id = ?", id).Error
	if err != nil {
		error2.ErrXtremeTestingGet(err.Error())
	}

	return testing
}

func (repo *testingRepository) Find(parameter url.Values) ([]model.Testing, interface{}, error) {
	var testings []model.Testing

	fromDate, toDate := core.SetDateRange(parameter)

	query := config.PgSQL.Where("\"createdAt\" BETWEEN ? AND ?", fromDate, toDate)

	if search := parameter.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query, pagination := core.Paginate(parameter, query, model.Testing{})
	err := query.Preload("Subs").Order("id DESC").Find(&testings).Error
	if err != nil {
		xtremelog.Error(err)
		return nil, nil, err
	}

	return testings, pagination.ParsePagination(), nil
}

func (repo *testingRepository) Store(req request.TestingRequest) model.Testing {
	testing := model.Testing{
		Name: req.Name,
	}

	err := repo.transaction.Create(&testing).Error
	if err != nil {
		error2.ErrXtremeTestingSave(err.Error())
	}

	return testing
}

func (repo *testingRepository) Delete(testing model.Testing) {
	err := repo.transaction.Delete(&testing).Error
	if err != nil {
		error2.ErrXtremeTestingDelete(err.Error())
	}
}

func (repo *testingRepository) AddSub(testing model.Testing, sub string) model.TestingSub {
	testingSub := model.TestingSub{
		TestingId: testing.ID,
		Name:      sub,
	}

	err := repo.transaction.Create(&testingSub).Error
	if err != nil {
		error2.ErrXtremeTestingSubSave(err.Error())
	}

	return testingSub
}

func (repo *testingRepository) DeleteSub(testingSub model.TestingSub) {
	err := repo.transaction.Delete(&testingSub).Error
	if err != nil {
		error2.ErrXtremeTestingSubDelete(err.Error())
	}
}
