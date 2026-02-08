package back

// EraseBytesSecurely overwrites all
// bytes in a slice with zeros.
func EraseBytesSecurely(input []byte) {
	for i := range input {
		input[i] = 0
	}
}
