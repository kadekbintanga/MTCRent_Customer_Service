package queue

import (
	"github.com/globalxtreme/gobaseconf/queue"
	Telegram "service/internal/app/queue/job"
	"service/internal/pkg/constant"
)

func Register() []queue.JobConf {
	return []queue.JobConf{
		{
			Context:     Telegram.TelegramMessageJob{},
			JobFunc:     (*Telegram.TelegramMessageJob).Consume,
			Concurrency: 1,
			QueueName:   constant.QUEUE_HIGH,
			JobName:     constant.JOB_TELEGRAM_MESSAGE,
			Priority:    10,
		},
	}
}
