package core

import (
	"encoding/json"
	"errors"
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type FormInterface interface {
	Validate()
}

type APIFormInterface interface {
	APIParse(r *http.Request)
}

type RabbitMQFormInterface interface {
	RabbitMQParse(message xtrememodel.RabbitMQMessage) error
}

type RabbitMQProcessedFormInterface interface {
	RabbitMQProcessedParse(message xtrememodel.RabbitMQMessage) (*xtremerabbitmq.RabbitMQDeliveryResponse, error)
}

type BaseForm struct{}

func (BaseForm) APIParse(r *http.Request, form interface{}) interface{} {
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		xtremeres.ErrXtremeBadRequest(err.Error())
	}

	return form
}

func (BaseForm) AsyncWorkflowParse(payload interface{}, form interface{}) error {
	if payload == nil {
		return errors.New("Your message is nil")
	}

	payloadMap, ok := payload.(map[string]interface{})
	if !ok {
		return errors.New("Your message is not a map")
	}

	err := mapstructure.Decode(payloadMap, &form)
	if err != nil {
		return errors.New("Your message parameter is invalid: " + err.Error())
	}

	return nil
}

func (BaseForm) RabbitMQParse(message xtrememodel.RabbitMQMessage, form interface{}) error {
	data, ok := message.Payload["data"].(map[string]interface{})
	if !ok {
		return errors.New("Your message is not map[string]interface{}")
	}

	err := mapstructure.Decode(data, &form)
	if err != nil {
		return errors.New("Your message parameter is invalid")
	}

	return nil
}

func (BaseForm) RabbitMQProcessedParse(message xtrememodel.RabbitMQMessage, form interface{}) (*xtremerabbitmq.RabbitMQDeliveryResponse, error) {
	data, ok := message.Payload["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("Your message is not map[string]interface{}")
	}

	response := &xtremerabbitmq.RabbitMQDeliveryResponse{}
	err := mapstructure.Decode(data, &response)
	if err != nil {
		return nil, errors.New("Your message parameter is invalid")
	}

	if response.Status.ID == xtremerabbitmq.RABBITMQ_MESSAGE_DELIVERY_STATUS_FINISH_ID {
		err = mapstructure.Decode(response.Result, &form)
		if err != nil {
			return nil, errors.New("Your response after processed is invalid")
		}
	}

	return response, nil
}
