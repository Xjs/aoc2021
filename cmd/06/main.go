package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Xjs/aoc2021/parse"
)

type anglerfish struct {
	counter int
	next    *anglerfish
}

func fishFromList(ns []int) *anglerfish {
	head := &anglerfish{}
	current := head
	for _, i := range ns {
		current.counter = i
		current.next = &anglerfish{}
		current = current.next
	}
	return head
}

func count(f *anglerfish) int {
	var sum int
	for current := f; current != nil; current = current.next {
		sum++
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
	transformations := 256
	l := 0
	for i := 0; i < transformations; i++ {
		transform(fish, live)
		l = count(fish)
		fmt.Println(i, "days:", l)
	}

	fmt.Println("After", transformations, "transformations:", l)
}

func transform(state *anglerfish, transformer func(*anglerfish)) {
	for current := state; current != nil; current = current.next {
		transformer(current)
	}
}

func live(f *anglerfish) {
	if f == nil {
		panic("nil fish doesn't live")
	}

	const cooldown = 6
	const birthPenalty = 2

	if f.counter == 0 {
		f.counter = cooldown
		oldNext := f.next
		f.next = &anglerfish{counter: cooldown + birthPenalty, next: oldNext}
	} else {
		f.counter--
	}
}
