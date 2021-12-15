package main

import (
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

type pointPriorityQueue struct {
	queue    []grid.Point
	priority map[grid.Point]int
}

func (q *pointPriorityQueue) addWithPriority(p grid.Point, priority int) {
	if q.priority == nil {
		q.priority = make(map[grid.Point]int)
	}

	q.priority[p] = priority
	var idx int
	if priority == math.MaxInt {
		idx = len(q.queue)
	} else {
		idx = q.search(p)
	}

	// Insert at idx
	q.queue = append(q.queue, p)
	if idx < len(q.queue)-1 {
		copy(q.queue[idx+1:], q.queue[idx:])
		q.queue[idx] = p
	}
}

// search does a binary search for the index at which to insert a point p with a pre-filled priority, or for the index of the concrete point p.
func (q *pointPriorityQueue) search(p grid.Point) int {
	if len(q.queue) == 0 {
		return 0
	}

	min := 0
	max := len(q.queue)
	var idx int
	prio := q.priority[p]
	for {
		idx = (max + min) / 2
		if max == min {
			break
		}
		if q.queue[idx] == p {
			break
		}

		if idxPrio := q.priority[q.queue[idx]]; idxPrio < prio {
			min = idx + 1
		} else if idxPrio > prio {
			max = idx
		} else {
			for i := min; i < max; i++ {
				if q.queue[i] == p {
					return i
				}
			}
		}
	}
	return idx
}

func (q *pointPriorityQueue) decreasePriority(p grid.Point, altPriority int) bool {
	if q.priority[p] < altPriority {
		// ignore, we will only decrease
		return false
	}

	idx := q.search(p)

	// precondition: q.queue[idx] == p
	for i := idx - 1; i > 0 && q.priority[q.queue[i]] > altPriority; i-- {
		q.queue[i+1], q.queue[i] = q.queue[i], q.queue[i+1]
	}
	q.priority[p] = altPriority
	return true
}

func (q *pointPriorityQueue) extractMin() (grid.Point, int) {
	var p grid.Point
	p, q.queue = q.queue[0], q.queue[1:]
	prio := q.priority[p]
	delete(q.priority, p)
	return p, prio
}

func (q *pointPriorityQueue) length() int {
	return len(q.queue)
}

func dijkstra(g integer.Grid, start, end grid.Point) []grid.Point {
	// tentative := integer.NewGrid(g.Width(), g.Height())
	// unvisited := make(map[grid.Point]struct{})
	tentativeQueue := new(pointPriorityQueue)
	tentativeQueue.addWithPriority(start, 0)

	for x := uint(0); x < g.Width(); x++ {
		for y := uint(0); y < g.Height(); y++ {
			p := grid.P(x, y)
			// unvisited[p] = struct{}{}
			if p == start {
				continue
			}
			tentativeQueue.addWithPriority(p, math.MaxInt)
		}
	}

	predecessors := make(map[grid.Point]grid.Point)

	iter := 0
	for {
		if tentativeQueue.length() == 0 {
			break
		}

		iter++

		current, curDist := tentativeQueue.extractMin()
		if current == end {
			break
		}

		if curDist == math.MaxInt {
			break
		}

		for _, p := range g.Environment4(current) {
			dist := curDist + g.MustAt(p)
			if tentativeQueue.decreasePriority(p, dist) {
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
