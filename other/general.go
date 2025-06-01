package other

import (
	"fmt"
	"os"

	"github.com/rwinkhart/go-boilerplate/back"
)

// PrintError prints an error message in the MUTN format and exits with the specified exit code.
// Requires: message (the error message to print),
// exitCode (the exit code to use)
func PrintError(message string, exitCode int) {
	fmt.Println(back.AnsiError + message + back.AnsiReset)
	os.Exit(exitCode)
}
