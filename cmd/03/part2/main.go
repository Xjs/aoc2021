package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Xjs/aoc2021/position"
)

func main() {
	var fullReport []string

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := s.Text()
		fullReport = append(fullReport, t)
	}
	if err := s.Err(); err != nil {
		log.Fatalf("Error scanning: %v", err)
	}

	var oxygenGeneratorRatingRaw, co2ScrubberRatingRaw []rune
	oxygenInputs := make([]string, len(fullReport))
	copy(oxygenInputs, fullReport)
	co2Inputs := make([]string, len(fullReport))
	copy(co2Inputs, fullReport)

	filter := func(s []string, r func([]string, int) rune) []string {
		var pos int
		for len(s) > 1 {
			s = position.Filter(s, pos, r(s, pos))
			pos++
		}
		return s
	}

	oxygenGeneratorRatingRaw = []rune(filter(oxygenInputs, func(inputs []string, pos int) rune {
		return position.Top(position.HistogramOfList(inputs)[pos])
	})[0])

	co2ScrubberRatingRaw = []rune(filter(co2Inputs, func(inputs []string, pos int) rune {
		return position.Bottom(position.HistogramOfList(inputs)[pos])
	})[0])

	oxygenGeneratorRating, err := strconv.ParseInt(string(oxygenGeneratorRatingRaw), 2, 64)
	if err != nil {
		log.Fatalf("Error parsing oxygen generator rating (%v): %v", oxygenGeneratorRatingRaw, err)
	}

	co2ScrubberRating, err := strconv.ParseInt(string(co2ScrubberRatingRaw), 2, 64)
	if err != nil {
		log.Fatalf("Error parsing CO2 scrubber rating (%v): %v", co2ScrubberRatingRaw, err)
	}

	fmt.Println("Life support rating (oxygen generator rating * CO2 scrubber rating) = ", oxygenGeneratorRating*co2ScrubberRating)
}
