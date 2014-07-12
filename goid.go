package goid

import (
	"math"
)

// Generates a pseudo-random (look alike) ID
// by using another base to represent it.
//
// The ID is generated within the given digits and
// the base is deducted from the size of the array.
//
// The seed is the number of already generated IDs.
//
// The minDigitCount is how many digits minimum should
// have the generated ID.
//
// This method doesn't return an error but can return an empty string
// if the given array of digits is empty.
func GenerateNext(digits []byte, seed int, minDigitCount int) string {
	if len(digits) == 0 {
		return "";
	}

	result := make([]byte, 0)

	// Special case for the 0 
	if seed == 0 {
		result = append(result, '0');
		for i := 0; i < minDigitCount-1; i++ {
			result = append(result, '0');
		}
	} else {
		// Guess which base we want to target
		base := len(digits)

		// offset to have a minimum amount of digits
		if minDigitCount > 1 {
			seed += int(math.Pow(float64(base), float64(minDigitCount-1)))
		}

		for i := seed; i%base > 0 || i/base > 0; i /= base {
			result = append(result, digits[i%base])
		}
	}


	return string(result)
}
