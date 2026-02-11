//go:build !windows

package front

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Input prompts the user for input and
// returns the input as a string.
func Input(prompt string) string {
	fmt.Print(prompt + " ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimRight(userInput, "\n\r ") // remove trailing newlines, carriage returns, and spaces
}

// InputBinary prompts the user with a yes/no
// question and returns the response as a boolean.
func InputBinary(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + " (y/N) ")
	char, _, _ := reader.ReadRune()
	if char == 'y' || char == 'Y' {
		return true
	}
	return false
}

// InputInt prompts the user for input and returns the input as an integer.
// A negative min/max value will disable the respective limit.
func InputInt(prompt string, min, max int) int {
	for {
		fmt.Print(prompt + " ")
		var userInput int
		_, err := fmt.Scanln(&userInput)
		if err == nil && (userInput >= min || min < 0) && (userInput <= max || max < 0) {
			return userInput
		}
		fmt.Println()
	}
}
