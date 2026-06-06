package parser

import (
	"service/internal/pkg/constant"
	"service/internal/pkg/model"
)

type CustomerParser struct {
	Array  []model.Customer
	Object model.Customer
}

func (parser CustomerParser) Get() []interface{} {
	var result []interface{}

	for _, customer := range parser.Array {
		firstParser := CustomerParser{Object: customer}
		result = append(result, firstParser.First())
	}
	return result
}

func (parser CustomerParser) First() interface{} {
	customer := parser.Object

	return map[string]interface{}{
		"uuid":            customer.UUID,
		"name":            customer.Name,
		"IDNumber":        customer.IDNumber,
		"SIMNumber":       customer.SIMNumber,
		"phone":           customer.Phone,
		"address":         customer.Address,
		"status":          constant.CustomerStatus{}.IDAndName(customer.StatusId),
		"blacklistReason": customer.BlacklistReason,
		"IDPhoto":         customer.IDPhoto,
		"createdAt":       customer.CreatedAt,
		"updatedAt":       customer.UpdatedAt,
	}
}

func (parser CustomerParser) CreateActivity(action string) interface{} {
	customer := parser.Object

	return map[string]interface{}{
		"id":              customer.ID,
		"uuid":            customer.UUID,
		"name":            customer.Name,
		"idNumber":        customer.IDNumber,
		"simNumber":       customer.SIMNumber,
		"phone":           customer.Phone,
		"address":         customer.Address,
		"status":          constant.CustomerStatus{}.IDAndName(customer.StatusId),
		"blacklistReason": customer.BlacklistReason,
		"IDPhoto":         customer.IDPhoto,
	}
}

func (parser CustomerParser) UpdateActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser CustomerParser) DeleteActivity(action string) interface{} {
	return parser.CreateActivity(action)
}

func (parser CustomerParser) GeneralActivity(action string) interface{} {
	if action == constant.ACTION_STATUS {
		customer := parser.Object

		return map[string]interface{}{
			"status":          constant.CustomerStatus{}.IDAndName(customer.StatusId),
			"blacklistReason": customer.BlacklistReason,
		}
	}

	return parser.CreateActivity(action)
}
