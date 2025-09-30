package util

import "math/rand"

// Creates a slice of size N
func CreateSlice(n int) []int {
	const intRange int = 100
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = rand.Intn(intRange)
	}
	return result
}
