package back

import (
	"errors"
	"os"
	"strings"
)

// TargetIsFile checks if the targetLocation is a file, directory, or is inaccessible.
// Requires: targetLocation,
// failOnDir (whether to return an error if targetLocation is a directory or a file).
// Returns: isAccessible.
func TargetIsFile(targetLocation string, failOnDir bool) (bool, error) {
	targetInfo, err := os.Stat(targetLocation)
	if err != nil {
		return false, errors.New("unable to access \"" + targetLocation + "\": " + err.Error())
	}
	if targetInfo.IsDir() {
		if failOnDir {
			return true, errors.New("\"" + targetLocation + "\" is a directory")
		}
		return true, nil
	} else {
		if !failOnDir {
			return true, errors.New("\"" + targetLocation + "\" is a file")
		}
		return true, nil
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
