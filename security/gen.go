package security

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

// BytesGen generates a random byte slice of a specified length and complexity.
// Requires: complexity (minimum percentage of special characters to be returned in the generated output; set to 0 for a "simple" result),
// complexCharsetLevel (1 = safe for filenames, 2 = safe for most password entries, 3 = safe only for well-made password entries)
func BytesGen(length int, complexity float64, complexCharsetLevel uint8) []byte {
	var actualSpecialChars int // track the number of special characters in the generated output
	var minSpecialChars int    // track the minimum number of special characters to accept
	var extendedCharset string // additions to character set used for complex outputs

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // default character set used for all strings
	const extendedCharsetFiles = "!#$%&+,-.;=@_~^()[]{}`'"                      // additional special characters for complex strings (safe in file names)
	const extendedCharsetMostPassword = "*:><?|"                                // additional special characters for complex strings (NOT safe in file names)
	const extendedCharsetSpecialPassword = "\"/\\"                              // additional special characters for complex strings (NOT safe in file names)

	if complexity > 0 {
		minSpecialChars = int(math.Round(float64(length) * complexity)) // determine minimum number of special characters to accept
		switch complexCharsetLevel {
		case 1:
			extendedCharset = extendedCharsetFiles
		case 2:
			extendedCharset = extendedCharsetMostPassword + extendedCharsetFiles[:len(extendedCharsetFiles)-9]
		case 3:
			extendedCharset = extendedCharsetFiles + extendedCharsetMostPassword + extendedCharsetSpecialPassword
		}
		charset += extendedCharset
	}

	// loop until a byte slice of the desired complexity is generated
	for {
		// generate a random output
		result := make([]byte, length)
		for i := range result {
			val, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
			result[i] = charset[val.Int64()]
		}

		// return early if the desired output is not complex
		if complexity <= 0 {
			return result
		}

		// count the number of special characters in the generated output
		for i := range result {
			if bytes.Contains([]byte(extendedCharset), []byte{result[i]}) {
				actualSpecialChars++
			}
		}

		// return the generated output if it contains enough special characters
		if actualSpecialChars >= minSpecialChars {
			return result
		}

		// reset special character counter
		fmt.Println("Regenerating output until desired complexity is achieved...")
		actualSpecialChars = 0
	}
}
