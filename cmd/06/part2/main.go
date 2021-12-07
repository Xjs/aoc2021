package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Xjs/aoc2021/parse"
)

func fishFromList(ns []int) map[int]int {
	result := make(map[int]int)
	for _, i := range ns {
		result[i]++
	}
	return result
}

func count(f map[int]int) int {
	var sum int
	for _, v := range f {
		sum += v
	}
	return sum
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	state, err := parse.IntList(string(input))
	if err != nil {
		log.Fatal(err)
	}
	fish := fishFromList(state)
	fmt.Println(fish)
	transformations := 256
	l := 0
	for i := 0; i < transformations; i++ {
		fish = transform(fish)
		l = count(fish)
		fmt.Println(i+1, "days:", l)
	}

	fmt.Println("After", transformations, "transformations:", l)
}

func transform(state map[int]int) map[int]int {
	const cooldown = 6
	const birthPenalty = 2

	newState := make(map[int]int)

	for counter, fish := range state {
		if counter == 0 {
			newState[cooldown] += fish
			newState[cooldown+birthPenalty] = fish
		} else {
			newState[counter-1] += fish
		}
	}

	return newState
}
