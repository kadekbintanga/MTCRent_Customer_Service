package Activity

import (
	"Service/App/Model/Activity"
	Repository "Service/App/Repository"
	"Service/App/Service/Helper"
	"Service/Config"
	"fmt"
	"gorm.io/gorm"
	"net/url"
)

type ActivityRepository struct{}

func (repo ActivityRepository) Get(parameters url.Values) ([]Activity.Activity, interface{}, error) {
	var activities []Activity.Activity

	query := repo.Query(parameters)
	query, pagination := Repository.Paginate(parameters, query, &Activity.Activity{})
	err := query.Order("id DESC").Find(&activities).Error
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return activities, pagination.ParsePagination(), nil
}

func (repo ActivityRepository) Query(parameters url.Values) *gorm.DB {
	fromDate, toDate := Helper.SetDateRange(parameters)

	query := Config.PgSQL.Where("\"createdAt\" BETWEEN ? AND ?", fromDate, toDate)

	if feature := parameters.Get("feature"); len(feature) > 0 {
		query = query.Where("feature = ?", feature)
	}

	if action := parameters.Get("action"); len(action) > 0 {
		query = query.Where("action = ?", action)
	}

	if search := parameters.Get("search"); len(search) > 3 {
		query = query.Where("description LIKE ? OR subFeature LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	return query
}
