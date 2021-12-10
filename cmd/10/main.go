package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/Xjs/aoc2021/part"
)

type stack struct {
	content []byte
}

func (s *stack) push(x byte) {
	s.content = append(s.content, x)
}

func (s *stack) pop() byte {
	var x byte
	x, s.content = s.content[len(s.content)-1], s.content[:len(s.content)-1]
	return x
}

var scorePart1 = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var scorePart2 = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func invert(b byte) byte {
	switch b {
	case '[':
		return ']'
	case '<':
		return '>'
	case '{':
		return '}'
	case '(':
		return ')'
	case ']':
		return '['
	case '>':
		return '<'
	case '}':
		return '{'
	case ')':
		return '('
	default:
		return b
	}
}

func check(line []byte) (int, stack) {
	var s stack
	for _, b := range line {
		switch b {
		case '(', '[', '{', '<':
			s.push(b)
		case ')', ']', '}', '>':
			open := s.pop()
			if open != invert(b) {
				return scorePart1[b], s
			}
		}
	}
	return 0, s
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(input, []byte("\n"))
	var score int
	var finals [][]byte
	for _, line := range lines {
		points, s := check(line)
		score += points
		if points > 0 {
			continue
		}
		var finalise []byte
		for len(s.content) > 0 {
			finalise = append(finalise, invert(s.pop()))
		}
		finals = append(finals, finalise)
	}

	if part.One() {
		fmt.Println(score)
		return
	}

	var finalScores []int
	for _, finalise := range finals {
		var score int
		for _, b := range finalise {
			score *= 5
			score += scorePart2[b]
		}
		finalScores = append(finalScores, score)
	}
	sort.IntSlice(finalScores).Sort()
	fmt.Println(finalScores[len(finalScores)/2])
}
