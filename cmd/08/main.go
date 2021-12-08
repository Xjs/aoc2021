package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Xjs/aoc2021/part"
)

type segment int

const (
	SegA segment = 1 << iota
	SegB
	SegC
	SegD
	SegE
	SegF
	SegG
)

func toBits(s string) segment {
	var result segment
	for _, r := range s {
		switch r {
		case 'a':
			result |= SegA
		case 'b':
			result |= SegB
		case 'c':
			result |= SegC
		case 'd':
			result |= SegD
		case 'e':
			result |= SegE
		case 'f':
			result |= SegF
		case 'g':
			result |= SegG
		default:
			// XXX: could handle errors if expecting broken input
			return 0
		}
	}
	return result

}

type line struct {
	patterns []string
	output   []string
}

func decode(l line) int {
	cipher := make(map[segment]int)
	reverse := make(map[int]segment)
	for _, pattern := range l.patterns {
		seg := toBits(pattern)
		switch len(pattern) {
		case 2:
			cipher[seg] = 1
			reverse[1] = seg
		case 4:
			cipher[seg] = 4
			reverse[4] = seg
		case 3:
			cipher[seg] = 7
			reverse[7] = seg
		case 7:
			cipher[seg] = 8
			reverse[8] = seg
		}
	}

	for _, pattern := range l.patterns {
		seg := toBits(pattern)
		if len(pattern) != 6 {
			continue
		}
		if seg&reverse[4] == reverse[4] {
			cipher[seg] = 9
			reverse[9] = seg
		} else if seg&reverse[7] == reverse[7] {
			cipher[seg] = 0
			reverse[0] = seg
		} else {
			cipher[seg] = 6
			reverse[6] = seg
		}
	}

	for _, pattern := range l.patterns {
		seg := toBits(pattern)
		if len(pattern) != 5 {
			continue
		}
		if seg&reverse[1] == reverse[1] {
			cipher[seg] = 3
			reverse[3] = seg
		}
	}

	bPiece := reverse[9] ^ reverse[3]

	for _, pattern := range l.patterns {
		seg := toBits(pattern)
		if len(pattern) != 5 {
			continue
		}
		if cipher[seg] == 3 {
			continue
		}
		if seg&bPiece == bPiece {
			cipher[seg] = 5
			reverse[5] = seg
		} else {
			cipher[seg] = 2
			reverse[2] = seg
		}
	}

	result := make([]int, 4)
	for i, pattern := range l.output {
		result[i] = cipher[toBits(pattern)]
	}
	return result[0]*1000 + result[1]*100 + result[2]*10 + result[3]
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var lines []line
	for s.Scan() {
		text := s.Text()
		parts := strings.Split(text, " | ")
		if len(parts) != 2 {
			log.Fatalf("Invalid line: %q", text)
		}
		lines = append(lines, line{patterns: strings.Fields(parts[0]), output: strings.Fields(parts[1])})
	}

	//   (6)     (2)     (5)     (5)     (4)
	//   0:      1:      2:      3:      4:
	//  aaaa    ....    aaaa    aaaa    ....
	// b    c  .    c  .    c  .    c  b    c
	// b    c  .    c  .    c  .    c  b    c
	//  ....    ....    dddd    dddd    dddd
	// e    f  .    f  e    .  .    f  .    f
	// e    f  .    f  e    .  .    f  .    f
	//  gggg    ....    gggg    gggg    ....
	//
	//   (5)     (6)     (3)     (7)     (6)
	//   5:      6:      7:      8:      9:
	//  aaaa    aaaa    aaaa    aaaa    aaaa
	// b    .  b    .  .    c  b    c  b    c
	// b    .  b    .  .    c  b    c  b    c
	//  dddd    dddd    ....    dddd    dddd
	// .    f  e    f  .    f  e    f  .    f
	// .    f  e    f  .    f  e    f  .    f
	//  gggg    gggg    ....    gggg    gggg

	if part.One() {
		var counts [10]int
		for _, line := range lines {
			for _, output := range line.output {
				switch len(output) {
				case 2:
					counts[1]++
				case 4:
					counts[4]++
				case 3:
					counts[7]++
				case 7:
					counts[8]++
				}
			}
		}

		for i, count := range counts {
			fmt.Println(i, "counted", count, "times")
		}

		return
	}

	var sum int
	for _, l := range lines {
		sum += decode(l)
	}
	fmt.Println(sum)
}
