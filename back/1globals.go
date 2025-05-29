package back

import "os"

var Home, _ = os.UserHomeDir()

const (
	AnsiBold  = "\033[1m"
	AnsiError = "\033[38;5;9m"
	AnsiReset = "\033[0m"

	ErrorRead            = 101
	ErrorWrite           = 102
	ErrorTargetNotFound  = 105
	ErrorTargetWrongType = 107
	ErrorOther           = 111
)
