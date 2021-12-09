package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Xjs/aoc2021/parse"
	"github.com/Xjs/aoc2021/part"
)

type grid struct {
	width, height int
	values        [][]int
}

type point struct {
	x, y int
}

func newGrid(w, h int) grid {
	g := grid{width: w, height: h, values: make([][]int, h)}
	for i := 0; i < h; i++ {
		g.values[i] = make([]int, w)
	}
	return g
}

func fill(source, target grid, p point, id int) {
	if target.at(p) != 0 {
		return
	}
	if source.at(p) == 9 {
		target.values[p.y][p.x] = 9
		return
	}
	target.values[p.y][p.x] = id
	for _, p := range target.environment(p) {
		fill(source, target, p, id)
	}
}

// fillAOC goes through the target, looks for the first zero, then fills all areas that are connected
// to this point and don't have a 9 in the source with the number id.
// Pixels with a 9 will become a 9 in the target.
// It returns false if no zero was found.
func fillAOC(source, target grid, id int) bool {
	if source.width != target.width || source.height != target.height {
		panic("sizes don't match")
	}

	for x := 0; x < source.width; x++ {
		for y := 0; y < source.height; y++ {
			if target.values[y][x] == 0 {
				fill(source, target, point{x, y}, id)
				return true
			}
		}
	}

	return false
}

// fillAllAOC returns a grid where all connected areas are filled with the same number != 9,
// and all 9s are kept as-is.
func fillAllAOC(source grid) grid {
	id := 10
	target := newGrid(source.width, source.height)
	for fillAOC(source, target, id) {
		id++
	}
	return target
}

func (g grid) at(p point) int {
	if p.y >= g.height || p.x >= g.width {
		panic(p)
	}
	return g.values[p.y][p.x]
}

func (g grid) environment(p point) []point {
	x, y := p.x, p.y
	result := make([]point, 0, 4)
	if x > 0 {
		result = append(result, point{x - 1, y})
	}
	if x < g.width-1 {
		result = append(result, point{x + 1, y})
	}
	if y > 0 {
		result = append(result, point{x, y - 1})
	}
	if y < g.height-1 {
		result = append(result, point{x, y + 1})
	}
	return result
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

	g := grid{values: depths}
	if len(depths) > 0 {
		g.width = len(depths[0])
	}
	g.height = len(depths)

	minima := make([]point, 0)

	if part.One() {
		for y := 0; y < g.height; y++ {
			for x := 0; x < g.width; x++ {
				v := g.values[y][x]
				smallest := true

				for _, p2 := range g.environment(point{x, y}) {
					v2 := g.at(p2)
					if v2 < v {
						smallest = false
						break
					}
				}
				if !smallest {
					continue
				}

				minima = append(minima, point{x, y})
			}
		}

		for _, minimum := range minima {
			fmt.Println(minimum.x, minimum.y, "->", g.values[minimum.y][minimum.x])
		}
		return
	}

	areas := fillAllAOC(g)
	hist := make(map[int]int)
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			hist[areas.values[y][x]]++
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
