package executor

import (
	"fmt"

	"github.com/tachRoutine/smartrun/pkg/types"
	"github.com/tacheraSasi/tripwire/utils"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

// Execute runs the given command and returns its output or an error
func (e *Executor) Execute(llmRes types.LLMResponse) error {
	for _, command := range llmRes.Commands {
		output, err := runCommand(command)
		if err != nil {
			return err
		}
		fmt.Println("Command output:", output)
	}
	return nil
}

func runCommand(command types.Command) (string, error) {
	if command.Dangerous{
		if utils.AskForConfirmation("Do you want to continue?") {
			fmt.Println("Continuing...")
		} else {
			return "Cancelled", fmt.Errorf("Cancelled")
		}
	}
	return "", nil
}