package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/Xjs/aoc2021/integer"
	"github.com/Xjs/aoc2021/integer/grid"
	"github.com/Xjs/aoc2021/part"
)

func main() {
	var g integer.Grid

	if gg, err := integer.ReadGrid(os.Stdin); err != nil {
		log.Fatal(err)
	} else {
		g = *gg
	}

	if !part.One() {
		const factor = 5

		width := g.Width()
		height := g.Height()

		largegrid := integer.NewGrid(factor*width, factor*height)

		for i := uint(0); i < factor; i++ {
			for j := uint(0); j < factor; j++ {

				for x := uint(0); x < width; x++ {
					for y := uint(0); y < height; y++ {
						p := grid.P(x+(i*width), y+(j*height))
						v := g.MustAt(grid.P(x, y)) + int(i) + int(j)
						if v > 9 {
							v -= 9
						}
						largegrid.MustSet(p, v)
					}
				}

			}
		}
		g = largegrid
	}

	ps := dijkstra(g, grid.P(0, 0), grid.P(g.Width()-1, g.Height()-1))
	fmt.Println(ps)

	var risk int
	for _, p := range ps {
		risk += g.MustAt(p)
	}

	fmt.Println("Risk:", risk)
}

func dijkstra(g integer.Grid, start, end grid.Point) []grid.Point {
	// tentative := integer.NewGrid(g.Width(), g.Height())
	// unvisited := make(map[grid.Point]struct{})
	tentativeQueue := newPointPriorityQueue()
	heap.Push(tentativeQueue, start)
	tentativeQueue.update(start, 0)

	for x := uint(0); x < g.Width(); x++ {
		for y := uint(0); y < g.Height(); y++ {
			p := grid.P(x, y)
			// unvisited[p] = struct{}{}
			if p == start {
				continue
			}
			heap.Push(tentativeQueue, p)
			tentativeQueue.update(p, math.MaxInt)
		}
	}

	heap.Init(tentativeQueue)

	predecessors := make(map[grid.Point]grid.Point)

	iter := 0
	for {
		if tentativeQueue.Len() == 0 {
			break
		}

		iter++

		current, curDist := tentativeQueue.pop()

		if current == end {
			break
		}

		if curDist == math.MaxInt {
			break
		}

		for _, p := range g.Environment4(current) {
			dist := curDist + g.MustAt(p)
			if tentativeQueue.decrease(p, dist) {
				predecessors[p] = current
			}
		}
	}

	var result []grid.Point

	current := end
	for current != start {
		var ok bool
		result = append([]grid.Point{current}, result...)
		current, ok = predecessors[current]
		if !ok {
			panic("no predecessor found")
		}
	}

	return result
}
