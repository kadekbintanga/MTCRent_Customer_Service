package Console

import (
	"Service/App/Console/Command"
	"github.com/globalxtreme/gobaseconf/console"
	"github.com/globalxtreme/gobaseconf/rabbitmq/command"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
)

func RegisterCommand(cobraCmd *cobra.Command) {
	addCommands(cobraCmd, []console.BaseInterface{
		&console.DeleteLogFileCommand{},
		&command.RabbitMQConsumeCommand{},

		&Command.TestCommand{},
	})
}

func RegisterSchedule(sch *gocron.Scheduler) {
	addSchedule(sch.Every(1).Minute(), &Command.TestCommand{})
}

func addCommands(cmd *cobra.Command, newCommands []console.BaseInterface) {
	for _, newCommand := range newCommands {
		newCommand.Command(cmd)
	}
}

func addSchedule(schedule *gocron.Scheduler, command console.BaseInterface) {
	schedule.Do(command.Handle)
}
