package other

import (
	"fmt"
	"os"

	"github.com/rwinkhart/go-boilerplate/back"
)

// PrintError prints an error message in the MUTN format and exits with the specified exit code.
// Requires: message (the error message to print),
// exitCode (the exit code to use),
// forceHardExit (if true, exit immediately; if false, allow soft exit for interactive clients).
func PrintError(message string, exitCode int, forceHardExit bool) {
	fmt.Println(back.AnsiError + message + back.AnsiReset)
	if forceHardExit {
		os.Exit(exitCode)
	} else {
		back.Exit(exitCode)
	}
}
