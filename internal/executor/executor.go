package executor

import (
	"fmt"
	"strings"

	"github.com/tachRoutine/smartrun/pkg/types"
	"github.com/tacheraSasi/tripwire/utils"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

// Execute runs the given command and returns its output or an error
func (e *Executor) Execute(llmRes types.LLMResponse) error {
	fmt.Println("Instructions:", llmRes.Instructions)
	for _, command := range llmRes.Commands {
		fmt.Println("Executing command:", command.Command, "\nDescription:", command.Description, "\nPlatform:", command.Platform, "\nDangerous:", command.Dangerous)
		output, err := runCommand(command)
		if err != nil {
			return err
		}
		fmt.Println("Command output:", output)
	}
	return nil
}

// runCommand executes a single command, handling dangerous commands with confirmation
func runCommand(command types.Command) (string, error) {
	if command.Dangerous {
		if utils.AskForConfirmation("Do you want to continue?") {
			commands := strings.Split(command.Command, "")
			output, err := Run(commands, "Error executing command:"+command.Command)
			if err != nil {
				return output, err
			}
		} else {
			return "Cancelled", fmt.Errorf("Cancelled")
		}
	}
	commands := strings.Split(command.Command, "")
	output, err := Run(commands, "Error executing command:"+command.Command)
	if err != nil {
		return "", err
	}
	return output, nil
}
