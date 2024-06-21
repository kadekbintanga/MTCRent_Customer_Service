package consumer

import (
	"errors"
	"fmt"
)

type TestingConsumer struct{}

func (cons TestingConsumer) Consume(message any) error {
	data, ok := message.(map[string]interface{})
	if !ok {
		return errors.New("Your message is not map[string]interface{}")
	}

	fmt.Println(data["name"])
	return nil
}
