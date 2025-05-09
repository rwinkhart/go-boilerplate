package back

import (
	"fmt"
	"os"
)

// RemoveTrailingEmptyStrings removes empty strings from the end of a slice.
func RemoveTrailingEmptyStrings(slice []string) []string {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] != "" {
			return slice[:i+1]
		}
	}
	return nil
}

// PrintError prints an error message in the standard libmutton format and exits with the specified exit code.
// Requires: message (the error message to print),
// exitCode (the exit code to use),
// forceHardExit (if true, exit immediately; if false, allow soft exit for interactive clients).
func PrintError(message string, exitCode int, forceHardExit bool) {
	fmt.Println(AnsiError + message + AnsiReset)
	if forceHardExit {
		os.Exit(exitCode)
	} else {
		Exit(exitCode)
	}
}
