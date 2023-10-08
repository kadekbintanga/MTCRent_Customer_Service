package Console

import (
	"github.com/globalxtreme/gobaseconf/console"
	"github.com/go-co-op/gocron"
)

func Schedules(sch *gocron.Scheduler) {
	//addSchedule(sch.Every(1).Day().At("00:05"), Command.TestCommand{})
}

func addSchedule(schedule *gocron.Scheduler, command console.BaseInterface) {
	schedule.Do(command.Handle)
}
