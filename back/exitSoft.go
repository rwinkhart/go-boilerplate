//go:build interactive

package back

// Exit (soft) is meant to be used in interactive implementations (GUIs/TUIs) to keep the program running after an operation.
func Exit(code int) {
	return
}
