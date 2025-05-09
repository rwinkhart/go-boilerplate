package back

import (
	"bufio"
	"io"
	"os"
	"os/exec"
)

// WriteToStdin is a utility function that writes a string to a command's stdin.
func WriteToStdin(cmd *exec.Cmd, input string) {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		PrintError("Failed to access stdin for system command: "+err.Error(), ErrorOther, true)
	}

	go func() {
		defer func(stdin io.WriteCloser) {
			_ = stdin.Close() // error ignored; if stdin could be accessed, it can probably be closed
		}(stdin)
		_, _ = io.WriteString(stdin, input)
	}()
}

// ReadFromStdin is a utility function that reads a string from stdin.
func ReadFromStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
