package web

import (
	"fmt"
	"log"
	"os"

	"github.com/rwinkhart/go-boilerplate/back"
)

// PrintLog colorizes a log and determines the best method for printing it.
func PrintLog(message, ansiColor string) {
	if ansiColor != "" {
		message = ansiColor + message + back.AnsiReset
	}
	if os.Getenv("INVOCATION_ID") != "" {
		fmt.Println(message)
	} else {
		log.Println(message)
	}
}
