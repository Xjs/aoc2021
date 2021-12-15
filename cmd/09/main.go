package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Xjs/aoc2021/integer"
	"github.com/Xjs/aoc2021/integer/grid"
	"github.com/Xjs/aoc2021/parse"
	"github.com/Xjs/aoc2021/part"
)

func fill(source, target integer.Grid, p grid.Point, id int) {
	if target.MustAt(p) != 0 {
		return
	}
	if source.MustAt(p) == 9 {
		(&target).MustSet(p, 9)
		return
	}
	(&target).MustSet(p, id)
	for _, p := range target.Environment4(p) {
		fill(source, target, p, id)
	}
}

// fillAOC goes through the target, looks for the first zero, then fills all areas that are connected
// to this point and don't have a 9 in the source with the number id.
// Pixels with a 9 will become a 9 in the target.
// It returns false if no zero was found.
func fillAOC(source, target integer.Grid, id int) bool {
	if source.Width() != target.Width() || source.Height() != target.Height() {
		panic("sizes don't match")
	}

	for x := uint(0); x < source.Width(); x++ {
		for y := uint(0); y < source.Height(); y++ {
			p := grid.P(x, y)
			if target.MustAt(p) == 0 {
				fill(source, target, p, id)
				return true
			}
		}
	}

	return false
}

// fillAllAOC returns a grid where all connected areas are filled with the same number != 9,
// and all 9s are kept as-is.
func fillAllAOC(source integer.Grid) integer.Grid {
	id := 10
	target := integer.NewGrid(source.Width(), source.Height())
	for fillAOC(source, target, id) {
		id++
	}
	return target
}

func main() {
	var depths [][]int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ds, err := parse.DigitList(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, ds)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	g, err := integer.GridFrom(depths)
	if err != nil {
		log.Fatal(err)
	}

	minima := make([]grid.Point, 0)

	if part.One() {
		for y := uint(0); y < g.Height(); y++ {
			for x := uint(0); x < g.Width(); x++ {
				p := grid.P(x, y)
				v := g.MustAt(p)
				smallest := true

				for _, p2 := range g.Environment4(p) {
					v2 := g.MustAt(p2)
					if v2 < v {
						smallest = false
						break
					}
				}
				if !smallest {
					continue
				}

				minima = append(minima, p)
			}
		}

		var risk int
		for _, minimum := range minima {
			v := g.MustAt(minimum)
			fmt.Println(minimum.X, minimum.Y, "->", v)
			risk += (v + 1)
		}

		fmt.Println("Risk:", risk)
		return
	}

	areas := fillAllAOC(g)
	hist := make(map[int]int)
	for y := uint(0); y < g.Height(); y++ {
		for x := uint(0); x < g.Width(); x++ {
			hist[areas.MustAt(grid.P(x, y))]++
		}
	}

	// q'n'd multiply highest 3 values
	delete(hist, 9)

	var highestValues []int
	for i := 0; i < 3; i++ {
		max := 0
		var maxK int
		for k, v := range hist {
			if v > max {
				max = v
				maxK = k
			}
		}
		delete(hist, maxK)
		highestValues = append(highestValues, max)
	}

	fmt.Println(highestValues)
}
