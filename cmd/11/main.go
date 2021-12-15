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
	g, err := integer.ReadGrid(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	maxSteps := 100
	if !part.One() {
		maxSteps = math.MaxInt
	}
	var flashes int

	for i := 0; i < maxSteps; i++ {
		g.Foreach(func(p grid.Point) {
			g.MustSet(p, g.MustAt(p)+1)
		})

		didFlash := make(map[grid.Point]struct{})
		g.Foreach(func(p grid.Point) {
			if g.MustAt(p) > 9 {
				flash(g, p, didFlash)
			}
		})

		var flashesThisStep int
		g.Foreach(func(p grid.Point) {
			if g.MustAt(p) > 9 {
				g.MustSet(p, 0)
				flashesThisStep++
			}
		})

		flashes += flashesThisStep
		if flashesThisStep == int(g.Width())*int(g.Height()) {
			fmt.Println("Synchronized at:", i+1)
			break
		}
	}

	fmt.Println("Flashes:", flashes)
}

func flash(g *integer.Grid, p grid.Point, didFlash map[grid.Point]struct{}) {
	if g.MustAt(p) < 10 {
		panic("invalid flash")
	}

	if _, ok := didFlash[p]; ok {
		return
	}
	didFlash[p] = struct{}{}

	env := g.Environment8(p)
	for _, p2 := range env {
		g.MustSet(p2, g.MustAt(p2)+1)
	}
	for _, p2 := range env {
		v := g.MustAt(p2)
		if v > 9 {
			flash(g, p2, didFlash)
		}
	}
}
