package front

import (
	"fmt"
	"os"
	"slices"

	"golang.org/x/term"
)

// InputMenuGen prompts the user with a menu
// and returns the user's choice as an integer.
func InputMenuGen(prompt string, options []string) int {
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	return InputInt("\n"+prompt, 1, len(options))
}

// InputMenuGenWithRuneInputs prompts the user with a menu
// and returns the user's choice as a rune.
// Ensure runeInputs contains only unique runes and matches the
// length of options.
func InputMenuGenWithRuneInputs(prompt string, options []string, runeInputs []rune) rune {
	for i, option := range options {
		fmt.Printf("%c. %s\n", runeInputs[i], option)
	}
	for {
		userInput := InputRune("\n" + prompt)
		if slices.Contains(runeInputs, userInput) {
			return userInput
		}
	}
}

// InputSecret prompts the user for input and returns the
// input as a byte array, hiding the input from the terminal.
func InputSecret(prompt string) []byte {
	fmt.Print(prompt + " ")
	byteInput, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return byteInput
}
