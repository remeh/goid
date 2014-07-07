package goid

// Generates a pseudo-random (look alike) ID
// by using another base to represent it.
//
// The ID is generated within the given values and
// the base is deducted from the size of the array.
//
// The seed is the number of already generated IDs.
func GenerateNext(values []byte, seed int) string {
	if seed == 0 {
		return "0"
	}

	if len(values) == 0 {
		return ""
	}

	// Guess which base we want to target
	base := len(values)

	result := make([]byte, 0)

	for i := seed; i%base > 0 || i/base > 0; i /= base {
		result = append(result, values[i%base])
	}

	return string(result)
}
