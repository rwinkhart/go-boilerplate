package back

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/exec"

	"github.com/rwinkhart/go-boilerplate/security"
)

// WriteToStdin is a utility function that writes a byte slice to a command's stdin.
func WriteToStdinAndZeroizeInput(cmd *exec.Cmd, input []byte) error {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return errors.New("unable to access stdin for system command: " + err.Error())
	}
	go func() {
		defer func(stdin io.WriteCloser) {
			_ = stdin.Close() // error ignored; if stdin could be accessed, it can probably be closed
		}(stdin)
		_, _ = stdin.Write(input)
		security.ZeroizeBytes(input)
	}()
	return nil
}

// ReadFromStdin is a utility function that reads a byte slice from stdin.
func ReadFromStdin() []byte {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Bytes()
	}
	return nil
}
