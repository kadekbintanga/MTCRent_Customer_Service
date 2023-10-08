package main

import (
	"Service/App/Job/Resizing"
	"Service/App/Service/Constant/Queue"
	"flag"
	"github.com/globalxtreme/gobaseconf/queue"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	queueNames := flag.String("queue", Queue.QUEUE_HIGH, "Queue name")
	flag.Parse()

	configurations := [...]queue.JobConf{
		{
			Context:     Resizing.Resizing{},
			JobFunc:     (*Resizing.Resizing).Image,
			Concurrency: 10,
			QueueName:   Queue.QUEUE_HIGH,
			JobName:     Queue.JOB_REZISE_IMAGE,
			Priority:    1,
		},
	}

	worker := queue.Queue{Names: *queueNames}
	worker.Work(configurations[:])
}
