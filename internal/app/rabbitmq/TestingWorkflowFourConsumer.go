package rabbitmq

import (
	"encoding/json"
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"
	"service/internal/pkg/core"
	form2 "service/internal/pkg/form"
	"sync"
)

type TestingWorkflowFourConsumer struct {
	xtremerabbitmq.AsyncWorkflowConsumerBase

	mutex sync.Mutex
	form  *form2.TestingForm
}

func (c *TestingWorkflowFourConsumer) Consume(payload interface{}) (interface{}, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return core.RabbitMQErrorHandler(func() (interface{}, error) {
		c.form = &form2.TestingForm{}
		err := c.form.AsyncWorkflowParse(payload)
		if err != nil {
			return nil, err
		}

		xtremepkg.LogInfo("Step 4")
		xtremepkg.LogInfo(c.GetReferenceId())
		xtremepkg.LogInfo(c.GetReferenceType())
		formJson, _ := json.Marshal(c.form)
		xtremepkg.LogInfo(string(formJson))

		return c.Response(payload), nil
	})
}

func (c *TestingWorkflowFourConsumer) Response(payload interface{}, data ...interface{}) interface{} {
	if c.form == nil {
		c.form = &form2.TestingForm{}
		err := c.form.AsyncWorkflowParse(payload)
		if err != nil {
			return nil
		}
	}

	return map[string]interface{}{
		"name": c.form.Name,
		"subs": c.form.Subs,
	}
}
