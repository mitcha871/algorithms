package main

import (
	"algorithms/util"
	"errors"
	"fmt"
	"log"
	"sort"
)

func main() {
	mySlice := util.CreateSlice(100)

	sort.Ints(mySlice)
	const searchTarget int = 35
	location, err := binarySearch(mySlice, searchTarget)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v was found in location %v\n", searchTarget, location)
}

// Binary search for a sorted array.
// Duplicates elements in the slice are not considered.
func binarySearch(slc []int, trgt int) (int, error) {
	lo := 0
	hi := len(slc)
	if hi <= 2 {
		return 0, errors.New("slice must have >2 elements")
	}
	rds := 0
	for lo < hi {
		rds++
		fmt.Printf("slice: %v\n", slc[lo:hi])
		n := (lo + hi) / 2
		if slc[n] > trgt {
			hi = n - 1
		} else if slc[n] < trgt {
			lo = n + 1
		} else if slc[n] == trgt {
			fmt.Printf("Found target in %v rounds\n", rds)
			return n, nil
		} else {
			return 0, errors.New("unreachable")
		}
	}
	return 0, errors.New("element not found in slice")
}
