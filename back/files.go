package back

import (
	"errors"
	"os"
	"strings"
)

// TargetIsFile checks if the targetLocation is a file, directory, or is inaccessible.
// Requires: failCondition (0 = fail on inaccessible, 1 = fail on inaccessible&file, 2 = fail on inaccessible&directory).
// Returns: isFile, isAccessible.
func TargetIsFile(targetLocation string, errorOnFail bool, failCondition uint8) (bool, bool, error) {
	targetInfo, err := os.Stat(targetLocation)
	if err != nil {
		if errorOnFail {
			return false, false, errors.New("unable to access \"" + targetLocation + "\": " + err.Error())
		}
		return false, false, nil
	}
	if targetInfo.IsDir() {
		if errorOnFail && failCondition == 2 {
			return false, true, errors.New("\"" + targetLocation + "\" is a directory")
		}
		return false, true, nil
	} else {
		if errorOnFail && failCondition == 1 {
			return true, true, errors.New("\"" + targetLocation + "\" is a file")
		}
		return true, true, nil
	}
}

// CreateTempFile creates a temporary file and returns a pointer to it.
func CreateTempFile() (*os.File, error) {
	tempFile, err := os.CreateTemp("", "*.markdown")
	if err != nil {
		return nil, errors.New("unable to create temporary file: " + err.Error())
	}
	return tempFile, nil
}

// ExpandPathWithHome returns the given path with "~" expanded to the user's home directory.
func ExpandPathWithHome(path string) string {
	return strings.Replace(path, "~", Home, 1)
}
