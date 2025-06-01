package back

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/exec"
)

// WriteToStdin is a utility function that writes a string to a command's stdin.
func WriteToStdin(cmd *exec.Cmd, input string) error {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return errors.New("unable to access stdin for system command: " + err.Error())
	}
	go func() {
		defer func(stdin io.WriteCloser) {
			_ = stdin.Close() // error ignored; if stdin could be accessed, it can probably be closed
		}(stdin)
		_, _ = io.WriteString(stdin, input)
	}()
	return nil
}

// ReadFromStdin is a utility function that reads a string from stdin.
func ReadFromStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
