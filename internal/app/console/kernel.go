package console

import (
	"github.com/globalxtreme/gobaseconf/console"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"service/internal/testing/command"
)

func RegisterCommand(cobraCmd *cobra.Command) {
	console.Commands(cobraCmd, []console.BaseInterface{
		&command.TestCommand{},
	})
}

func RegisterSchedule(sch *gocron.Scheduler) {
	//addSchedule(sch.Every(1).Minute(), &Command.TestCommand{})
}

func addSchedule(schedule *gocron.Scheduler, command console.BaseInterface) {
	schedule.Do(command.Handle)
}
