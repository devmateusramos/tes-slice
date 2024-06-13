package main

import (
	"fmt"
	"slices"
)

func main() {
	myList := []int{1, 5, 3}

	// In newer versions of Go you can do this, but I left the solution
	// iterating over the slice in the internal package along with the test file.
	max := slices.Max(myList)
	fmt.Println("Max in Slice =", max)
}
