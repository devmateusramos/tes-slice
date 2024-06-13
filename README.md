FindMax Function in Go
This repository contains two implementations of a function to find the maximum element in an integer slice ([]int), along with corresponding tests.

Internal Package Solution
The primary implementation resides in the internal package. The FindMax function located in internal/internal.go iterates through the slice to determine the maximum value. Tests for this function can be found in internal/service_test.go, covering scenarios such as normal cases, an empty slice, and a slice with a single element.

Alternative Solution using slices Package
For newer versions of Go, an alternative solution is showcased in main.go. This solution utilizes the slices.Max function from the standard library.
