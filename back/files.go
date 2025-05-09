package back

import (
	"os"
	"strings"
)

// TargetIsFile checks if the targetLocation is a file, directory, or is inaccessible.
// Requires: failCondition (0 = fail on inaccessible, 1 = fail on inaccessible&file, 2 = fail on inaccessible&directory).
// Returns: isFile, isAccessible.
func TargetIsFile(targetLocation string, errorOnFail bool, failCondition uint8) (bool, bool) {
	targetInfo, err := os.Stat(targetLocation)
	if err != nil {
		if errorOnFail {
			PrintError("Failed to access \""+targetLocation+"\" - Ensure it exists and has the correct permissions", ErrorTargetNotFound, true)
		}
		return false, false
	}
	if targetInfo.IsDir() {
		if errorOnFail && failCondition == 2 {
			PrintError("\""+targetLocation+"\" is a directory", ErrorTargetWrongType, true)
		}
		return false, true
	} else {
		if errorOnFail && failCondition == 1 {
			PrintError("\""+targetLocation+"\" is a file", ErrorTargetWrongType, true)
		}
		return true, true
	}
}

// CreateTempFile creates a temporary file and returns a pointer to it.
func CreateTempFile() *os.File {
	tempFile, err := os.CreateTemp("", "*.markdown")
	if err != nil {
		PrintError("Failed to create temporary file: "+err.Error(), ErrorWrite, true)
	}
	return tempFile
}

// ExpandPathWithHome, given a path (as a string) containing "~", returns the path with "~" expanded to the user's home directory.
func ExpandPathWithHome(path string) string {
	return strings.Replace(path, "~", Home, 1)
}
