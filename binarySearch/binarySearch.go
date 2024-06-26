package main

import (
	"fmt"
	"time"
)

func binarySearchRecursiveOption(searchField []int, target int) (index int, searchCount int) {
	mid := len(searchField) / 2

	switch {
	case len(searchField) == 0:
		index = -1
	case target < searchField[mid]:
		index, searchCount = binarySearchRecursiveOption(searchField[:mid], target)
	case target > searchField[mid]:
		index, searchCount = binarySearchRecursiveOption(searchField[mid+1:], target)
		if index >= 0 {
			index += mid + 1
		}
	default: // index == mid
		index = mid
	}
	searchCount++
	return
}

func main() {
	var field []int = []int{1, 4, 8, 12, 15, 18, 21, 25, 28, 31, 32}

	for _, val := range field {
		fmt.Printf("Searching for %d, in list %v\n", val, field)
		index, searchCount := binarySearchRecursiveOption(field, val)
		fmt.Printf("Numbers has been finded in postion %d, after %d step", index, searchCount)
		if searchCount == 1 {
			fmt.Print("\n\n")
		} else {
			fmt.Print("s\n\n")
		}
	}

	testField := make([]int, 10000000)
	for i := range testField {
		testField[i] = i
	}

	start := time.Now()
	for i := 0; i < 100; i++ {
		binarySearchRecursiveOption(testField, 7777777)
	}
	end := time.Since(start)
	nanoseconds := end.Nanoseconds() / 100
	fmt.Printf("Time: %d ns\n", nanoseconds)
}
