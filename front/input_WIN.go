//go:build windows

package front

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	procGetStdHandle     = kernel32.NewProc("GetStdHandle")
	procPeekConsoleInput = kernel32.NewProc("PeekConsoleInputW")
	procReadConsoleInput = kernel32.NewProc("ReadConsoleInputW")
)

const (
	stdInputHandle = ^uintptr(10) - 1 // STD_INPUT_HANDLE = -10
)

type inputRecord struct {
	EventType uint16
	_         [2]byte // padding
	Event     [16]byte
}

// pollForInput checks if input is available without blocking
func pollForInput() bool {
	handle, _, _ := procGetStdHandle.Call(stdInputHandle)
	var numEvents uint32
	var record inputRecord

	ret, _, _ := procPeekConsoleInput.Call(
		handle,
		uintptr(unsafe.Pointer(&record)),
		1,
		uintptr(unsafe.Pointer(&numEvents)),
	)

	return ret != 0 && numEvents > 0
}

// Input prompts the user for input and
// returns the input as a string.
// Uses polling to be more responsive
// to interrupts on Windows.
func Input(prompt string) string {
	fmt.Print(prompt + " ")

	// Poll until input is available
	for !pollForInput() {
		time.Sleep(50 * time.Millisecond)
	}

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimRight(userInput, "\n\r ") // remove trailing newlines, carriage returns, and spaces
}

// InputBinary prompts the user with a yes/no
// question and returns the response as a boolean.
// Uses polling to be more responsive to interrupts on Windows.
func InputBinary(prompt string) bool {
	fmt.Print(prompt + " (y/N) ")

	// Poll until input is available
	for !pollForInput() {
		time.Sleep(50 * time.Millisecond)
	}

	reader := bufio.NewReader(os.Stdin)
	char, _, _ := reader.ReadRune()
	if char == 'y' || char == 'Y' {
		return true
	}
	return false
}

// InputInt prompts the user for input and returns the input as an integer.
// A negative min/max value will disable the respective limit.
// Uses polling to be more responsive to interrupts on Windows.
func InputInt(prompt string, min, max int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt + " ")

		// Poll until input is available
		for !pollForInput() {
			time.Sleep(50 * time.Millisecond)
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println()
			continue
		}

		line = strings.TrimSpace(line)
		userInput, err := strconv.Atoi(line)

		if err == nil && (userInput >= min || min < 0) && (userInput <= max || max < 0) {
			return userInput
		}
		fmt.Println()
	}
}
