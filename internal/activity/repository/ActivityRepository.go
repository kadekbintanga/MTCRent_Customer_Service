package repository

import (
	"fmt"
	"gorm.io/gorm"
	"net/url"
	"service/internal/pkg/config"
	"service/internal/pkg/core"
	"service/internal/pkg/model"
)

/** --- INTERFACE --- */

type ActivityRepository interface {
	Find(parameters url.Values) ([]model.Activity, interface{}, error)
}

func NewActivityRepository() ActivityRepository {
	return activityRepository{}
}

/** --- MAIN REPOSITORY --- */

type activityRepository struct {
	Transaction *gorm.DB
}

func (repo activityRepository) Find(parameters url.Values) ([]model.Activity, interface{}, error) {
	var activities []model.Activity

	query := repo.filterByParam(parameters)
	query, pagination := core.Paginate(parameters, query, &model.Activity{})
	err := query.Order("id DESC").Find(&activities).Error
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return activities, pagination.ParsePagination(), nil
}

/** --- UNEXPORTED FUNCTIONS --- */

func (repo activityRepository) filterByParam(parameters url.Values) *gorm.DB {
	fromDate, toDate := core.SetDateRange(parameters)

	query := config.PgSQL.Where("\"createdAt\" BETWEEN ? AND ?", fromDate, toDate)

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
