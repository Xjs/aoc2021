package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/Xjs/aoc2021/integer"
	"github.com/Xjs/aoc2021/parse"
	"github.com/Xjs/aoc2021/part"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	positions, err := parse.IntList(string(input))
	if err != nil {
		log.Fatal(err)
	}

	fuel := fuel1
	if !part.One() {
		fuel = fuel2
	}

	max := integer.Max(positions)
	minSum := math.MaxInt
	minPos := -1

	for i := 0; i <= max; i++ {
		var sum int
		for _, pos := range positions {
			sum += fuel(integer.Abs(pos - i))
		}
		if sum < minSum {
			minSum = sum
			minPos = i
		}
	}

	fmt.Printf("Optimal position: %d, fuel use: %d\n", minPos, minSum)
}

func fuel1(distance int) int {
	return distance
}

func fuel2(distance int) int {
	return distance * (distance + 1) / 2
}
