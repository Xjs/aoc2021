package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// readAllInts reads from r line to line and returns a slice containing all lines that are valid integers.
func readAllInts(r io.Reader) []int {
	var result []int

	s := bufio.NewScanner(r)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			continue
		}
		result = append(result, i)
	}

	return result
}

// countIncreases sums up the numbers in ns in rolling windows, comparing the sums
// and returning the number of times a sum is larger than the sum of the previous window.
func countIncreases(ns []int, windowSize int) int {
	var increases int
	// Set to MaxInt to never count the first measurement as increase
	previous := math.MaxInt
	for i := 0; i+windowSize <= len(ns); i++ {
		var current int
		for j := 0; j < windowSize; j++ {
			current += ns[i+j]
		}
		if current > previous {
			increases++
		}
		previous = current
	}
	return increases
}

func main() {
	is := readAllInts(os.Stdin)
	fmt.Printf("Window size 1: %d\n", countIncreases(is, 1))
	fmt.Printf("Window size 3: %d\n", countIncreases(is, 3))
}
