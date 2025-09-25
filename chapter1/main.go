package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func main() {
	myTestSlice, err := buildTestSlice(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("My slice is: %v\n", myTestSlice)

	result, err := findMaxBasic(myTestSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[Basic] The max value is %v\n", result)

	result, err = findMaxBasicAlt(myTestSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[BasicAlt] The max value is %v\n", result)

	topTwo, err := twoLargestValues(myTestSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[TopTwo] The max values are %v\n", topTwo)
}

// Returns a slice of random positive integers.
func buildTestSlice(length int) ([]int, error) {
	if length < 1 {
		return nil, errors.New("slice length must be an integer greater than 0")
	}
	result := make([]int, length)
	const randRange = 100
	for i := range length {
		result[i] = rand.Intn(randRange) + 1 // [1, randRange]
	}
	return result, nil
}

// Basic algorithm to find the highest value in a slice of integers.
// O(n) comparisons
func findMaxBasic(slc []int) (int, error) {
	if len(slc) == 0 {
		return 0, errors.New("slice is empty")
	}
	result := slc[0]
	for i := range len(slc) {
		if slc[i] > result {
			result = slc[i]
		}
	}
	return result, nil
}

// Alternative Basic algorithm that breaks the loop if a larger value is found.
// Best case: O(n) comparisons
// Worst case: O(n^2) comparisons
func findMaxBasicAlt(slc []int) (int, error) {
	if len(slc) == 0 {
		return 0, errors.New("slice is empty")
	}
	for i := range len(slc) {
		v := slc[i]
		largest := true
		for j := range len(slc) {
			if v < slc[j] {
				largest = false
				break
			}
		}
		if largest {
			return v, nil
		}
	}
	return 0, errors.New("unreachable")
}

// Finds the two largest values in a slice
func twoLargestValues(slc []int) ([2]int, error) {
	if len(slc) < 2 {
		return [2]int{0, 0}, errors.New("slice must have >2 elements")
	}
	var mx int
	var second int
	if slc[0] > slc[1] {
		mx, second = slc[0], slc[1]
	} else {
		mx, second = slc[1], slc[0]
	}
	for i := 2; i < len(slc); i++ {
		v := slc[i]
		if v > mx {
			second = mx
			mx = v
		} else if v > second {
			second = v
		}
	}
	return [2]int{mx, second}, nil
}
