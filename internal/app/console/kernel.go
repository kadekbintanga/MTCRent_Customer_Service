package console

import (
	xtremeconsole "github.com/globalxtreme/go-core/v2/console"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"service/internal/app/console/command"
)

func RegisterCommand(cobraCmd *cobra.Command) {
	xtremeconsole.Commands(cobraCmd, []xtremeconsole.BaseCommand{
		&command.TestCommand{},
	})
}

func RegisterSchedule(sch *gocron.Scheduler) {
	//addSchedule(sch.Every(1).Minute(), &Command.TestCommand{})
}

func addSchedule(schedule *gocron.Scheduler, command xtremeconsole.BaseCommand) {
	schedule.Do(command.Handle)
}
