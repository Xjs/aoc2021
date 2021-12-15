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

type pointPriorityQueue struct {
	queue    []grid.Point
	priority map[grid.Point]int
	indexes  map[grid.Point]int
}

func newPointPriorityQueue() *pointPriorityQueue {
	return &pointPriorityQueue{
		priority: make(map[grid.Point]int),
		indexes:  make(map[grid.Point]int),
	}
}

func (pq *pointPriorityQueue) Len() int { return len(pq.queue) }

func (pq *pointPriorityQueue) Less(i, j int) bool {
	return pq.priority[pq.queue[i]] > pq.priority[pq.queue[j]]
}

func (pq *pointPriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
	pq.indexes[pq.queue[i]] = i
	pq.indexes[pq.queue[j]] = j
}

func (pq *pointPriorityQueue) Push(x interface{}) {
	n := len(pq.queue)
	item := x.(grid.Point)
	pq.indexes[item] = n
	pq.queue = append(pq.queue, item)
}

func (pq *pointPriorityQueue) Pop() interface{} {
	old := pq.queue
	n := len(old)
	item := old[n-1]
	delete(pq.indexes, item)
	pq.queue = old[0 : n-1]
	return item
}

func (pq *pointPriorityQueue) pop() (grid.Point, int) {
	p := heap.Pop(pq).(grid.Point)
	prio := pq.priority[p]
	delete(pq.priority, p)
	return p, -prio
}

// decrease modifies the priority and value of an Item in the queue only if it is a decrease.
func (pq *pointPriorityQueue) decrease(item grid.Point, priority int) bool {
	oldPrio := pq.priority[item]
	if oldPrio > priority {
		return false
	}
	pq.update(item, priority)
	return true
}

// update modifies the priority and value of an Item in the queue.
func (pq *pointPriorityQueue) update(item grid.Point, priority int) {
	pq.priority[item] = priority
	heap.Fix(pq, pq.indexes[item])
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
			tentativeQueue.update(p, -math.MaxInt)
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
			if tentativeQueue.decrease(p, -dist) {
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
