package Testing

import (
	"Service/App/Model/Testing"
	Repository "Service/App/Repository"
	"Service/App/Service/Helper"
	"Service/Config"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"net/url"
)

type TestingRepository struct{}

func (repo TestingRepository) Get(paramters url.Values) ([]Testing.Testing, interface{}, error) {
	var testings []Testing.Testing

	fromDate, toDate := Helper.SetDateRange(paramters)

	query := Config.PgSQL.Where("\"createdAt\" BETWEEN ? AND ?", fromDate, toDate)

	if search := paramters.Get("search"); len(search) > 3 {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	query, pagination := Repository.Paginate(paramters, query, Testing.Testing{})
	err := query.Preload("Subs").Order("id DESC").Find(&testings).Error
	if err != nil {
		xtremelog.Error(err)
		return nil, nil, err
	}

	return testings, pagination.ParsePagination(), nil
}
