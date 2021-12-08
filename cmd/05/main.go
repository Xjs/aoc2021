package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Xjs/aoc2021/parse"
	"github.com/Xjs/aoc2021/part"
)

type point struct {
	x, y int
}

type line struct {
	p1, p2 point
}

func newPoint(s string) (point, error) {
	ps, err := parse.IntList(s)
	if err != nil {
		return point{}, err
	}
	if len(ps) != 2 {
		return point{}, fmt.Errorf("invalid point %q, must have syntax x,y", s)
	}

	return point{ps[0], ps[1]}, nil
}

func newLine(s string) (line, error) {
	fields := strings.Split(s, "->")
	if len(fields) != 2 {
		return line{}, fmt.Errorf("invalid line %q, must have syntax x1,x2 -> y1,y2", s)
	}

	p1, err := newPoint(fields[0])
	if err != nil {
		return line{}, err
	}
	p2, err := newPoint(fields[1])
	if err != nil {
		return line{}, err
	}

	return line{p1, p2}, nil
}

func inc(x1, x2 int) int {
	if x1 > x2 {
		return -1
	} else if x1 < x2 {
		return 1
	}
	return 0
}

func main() {
	var lines []line

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		l, err := newLine(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, l)
	}

	var xMax, yMax int
	xMax = 10
	yMax = 10

	for _, line := range lines {
		for _, p := range []point{line.p1, line.p2} {
			if p.x > xMax {
				xMax = p.x
			}
			if p.y > yMax {
				yMax = p.y
			}
		}
	}
	xMax++
	yMax++

	grid := make([][]int, yMax)
	for i := 0; i < yMax; i++ {
		grid[i] = make([]int, xMax)
	}

	for _, line := range lines {
		if part.One() {
			if line.p1.x != line.p2.x && line.p1.y != line.p2.y {
				continue
			}
		}

		incX := inc(line.p1.x, line.p2.x)
		incY := inc(line.p1.y, line.p2.y)

		for x, y := line.p1.x, line.p1.y; ; x, y = x+incX, y+incY {
			grid[y][x]++

			if x == line.p2.x && y == line.p2.y {
				break
			}
		}
	}

	var count int
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if grid[y][x] > 1 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func printgrid(grid [][]int, xMax, yMax int) {
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
}
