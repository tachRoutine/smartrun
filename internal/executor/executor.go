package executor

import (
	"fmt"
	"os/exec"
	"runtime"
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
	fmt.Println(strings.Repeat("-", 50))

	for i, command := range llmRes.Commands {
		fmt.Printf("\n[%d/%d] Executing: %s\n", i+1, len(llmRes.Commands), command.Command)
		fmt.Printf("Description: %s\n", command.Description)
		fmt.Printf("Platform: %s | Dangerous: %v\n", command.Platform, command.Dangerous)

		// Check platform compatibility
		if !e.isPlatformCompatible(command.Platform) {
			fmt.Printf("Skipping command - not compatible with current platform (%s)\n", runtime.GOOS)
			continue
		}

		output, err := e.runCommand(command)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("Output:\n%s\n", output)
		fmt.Println(strings.Repeat("-", 30))
	}
	return nil
}

// runCommand executes a single command, handling dangerous commands with confirmation
func (e *Executor) runCommand(command types.Command) (string, error) {
	if command.Dangerous {
		fmt.Printf("WARNING: This command is marked as dangerous!\n")
		fmt.Printf("Command: %s\n", command.Command)

		if !utils.AskForConfirmation("Do you want to continue?") {
			return "Command cancelled by user", nil
		}
	}

	// Split command properly by spaces,
	args := e.parseCommand(command.Command)
	if len(args) == 0 {
		return "", fmt.Errorf("empty command")
	}

	// Create the command
	cmd := exec.Command(args[0], args[1:]...)

	// Execute and get output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("command failed: %v", err)
	}

	return string(output), nil
}

// parseCommand splits a command string into command and arguments
// Handles basic quoted strings
func (e *Executor) parseCommand(cmdStr string) []string {
	// TODO: I will enhance this for quoted strings
	return strings.Fields(strings.TrimSpace(cmdStr))
}

// isPlatformCompatible checks if the command is compatible with current OS
func (e *Executor) isPlatformCompatible(platform string) bool {
	if platform == "all" {
		return true
	}

	currentOS := runtime.GOOS
	switch platform {
	case "linux":
		return currentOS == "linux"
	case "macos":
		return currentOS == "darwin"
	case "windows":
		return currentOS == "windows"
	default:
		return false
	}
}
