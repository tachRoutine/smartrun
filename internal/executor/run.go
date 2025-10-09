package executor

import (
	"fmt"
	"os/exec"
)

func Run(cmdArgs []string, errMsg string) (string, error) {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s %s\n", errMsg, err)
		return "", err
	}
	if len(output) > 0 {
		fmt.Printf("Output:\n%s\n", output)
	}
	return string(output), nil
}
