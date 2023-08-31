package perl

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func Exec(command string) error {
	cmd := exec.Command("perl")
	cmd.Stdin = bytes.NewBufferString(command)

	// The `Output` method executes the command and
	// collects the output, returning its value
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, _ := cmd.Output()
	if stderr.Len() > 0 {
		return errors.New(stderr.String())
	}
	// otherwise, print the output from running the command
	fmt.Println("Output: ", string(out))
	return nil
}
