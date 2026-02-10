package security

// ZeroizeBytes overwrites all
// bytes in a slice with zeros.
func ZeroizeBytes(input []byte) {
	for i := range input {
		input[i] = 0
	}
}
