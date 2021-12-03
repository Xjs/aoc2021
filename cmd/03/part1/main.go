package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/Xjs/aoc2021/position"
)

func main() {
	c := make(chan string)
	var countsByPosition map[int]map[rune]int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		countsByPosition = position.Histogram(c)
		wg.Done()
	}()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		c <- s.Text()
	}
	close(c)
	if err := s.Err(); err != nil {
		log.Fatalf("Error scanning: %v", err)
	}
	wg.Wait()

	var length int
	for pos := range countsByPosition {
		if pos > length {
			length = pos
		}
	}
	length++

	gammaRaw := make([]rune, length)
	epsilonRaw := make([]rune, length)
	for i := 0; i < length; i++ {
		gammaRaw[i] = position.Top(countsByPosition[i])
		epsilonRaw[i] = position.Bottom(countsByPosition[i])
	}

	gamma, err := strconv.ParseInt(string(gammaRaw), 2, 64)
	if err != nil {
		log.Fatalf("Error parsing gamma (%q): %v", gammaRaw, err)
	}
	epsilon, err := strconv.ParseInt(string(epsilonRaw), 2, 64)
	if err != nil {
		log.Fatalf("Error parsing epsilon (%q): %v", gammaRaw, err)
	}

	fmt.Println("Power consumption (gamma * epsilon) = ", gamma*epsilon)
}
