package main

import (
	"container/heap"

	"github.com/Xjs/aoc2021/integer/grid"
)

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

// Len implements sort.Interface
func (pq *pointPriorityQueue) Len() int { return len(pq.queue) }

// Less implements sort.Interface
func (pq *pointPriorityQueue) Less(i, j int) bool {
	return pq.priority[pq.queue[i]] < pq.priority[pq.queue[j]]
}

// Swap implements sort.Interface
func (pq *pointPriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
	pq.indexes[pq.queue[i]] = i
	pq.indexes[pq.queue[j]] = j
}

// Push implements heap.Interface
func (pq *pointPriorityQueue) Push(x interface{}) {
	n := len(pq.queue)
	item := x.(grid.Point)
	pq.indexes[item] = n
	pq.queue = append(pq.queue, item)
}

// Pop implements heap.Interface
func (pq *pointPriorityQueue) Pop() interface{} {
	old := pq.queue
	n := len(old)
	item := old[n-1]
	delete(pq.indexes, item)
	pq.queue = old[0 : n-1]
	return item
}

// pop returns the point with lowest priority and its priority.
func (pq *pointPriorityQueue) pop() (grid.Point, int) {
	p := heap.Pop(pq).(grid.Point)
	prio := pq.priority[p]
	delete(pq.priority, p)
	return p, prio
}

// decrease modifies the priority and value of a point in the queue only if the priority is lower than the current priority.
// It returns true if a modification was made.
func (pq *pointPriorityQueue) decrease(item grid.Point, priority int) bool {
	oldPrio := pq.priority[item]
	if oldPrio < priority {
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
