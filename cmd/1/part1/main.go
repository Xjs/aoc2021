package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// countIncreases reads data from r line to line until the end, parses
// each line as integer (ignoring lines that do not contain integers)
// and counts how often a number is strictly larger than the previous
// one, returning the number of such occurences.
func countIncreases(r io.Reader) int {
	var (
		increases int
		// Set to MaxInt to never count the first measurement as increase
		previous = math.MaxInt
	)

	s := bufio.NewScanner(r)
	for s.Scan() {
		current, err := strconv.Atoi(s.Text())
		if err != nil {
			continue
		}
		if current > previous {
			increases++
		}

		previous = current
	}

	return increases
}

func main() {
	fmt.Println(countIncreases(os.Stdin))
}
