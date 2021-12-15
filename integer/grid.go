package integer

import (
	"errors"
	"fmt"

	"github.com/Xjs/aoc2021/integer/grid"
)

// ulen is a convenience function that returns the length of an int
// slice as uint.
func ulen(s []int) uint {
	return uint(len(s))
}

// A Grid represents a two-dimensional rectangular grid of integers.
type Grid struct {
	width, height uint
	values        [][]int
}

func (g Grid) Width() uint {
	return g.width
}

func (g Grid) Height() uint {
	return g.height
}

// NewGrid creates a new zero-filled grid
func NewGrid(w, h uint) Grid {
	g := Grid{width: w, height: h, values: make([][]int, h)}
	for i := uint(0); i < h; i++ {
		g.values[i] = make([]int, w)
	}
	return g
}

// GridFrom creates a new Grid from the given values, using the entries of
// the outer slice as rows. It will return an error
// if the rows are not of the same length.
func GridFrom(values [][]int) (Grid, error) {
	g := Grid{height: uint(len(values)), values: values}
	for i, row := range values {
		if i == 0 {
			g.width = ulen(row)
		}
		if ulen(row) != g.width {
			return g, fmt.Errorf("length of row %d is unequal to previous: %d", len(row), g.width)
		}
	}

	return g, nil
}

var ErrOutOfBounds = errors.New("out of bounds access to grid")

func (g Grid) At(p grid.Point) (int, error) {
	if p.Y >= g.height || p.X >= g.width {
		return 0, ErrOutOfBounds
	}
	return g.values[p.Y][p.X], nil
}

func (g Grid) MustAt(p grid.Point) int {
	v, err := g.At(p)
	if err != nil {
		panic(err)
	}
	return v
}

func (g Grid) Environment4(p grid.Point) []grid.Point {
	x, y := p.X, p.Y
	result := make([]grid.Point, 0, 4)
	if x > 0 {
		result = append(result, grid.P(x-1, y))
	}
	if x < g.width-1 {
		result = append(result, grid.P(x+1, y))
	}
	if y > 0 {
		result = append(result, grid.P(x, y-1))
	}
	if y < g.height-1 {
		result = append(result, grid.P(x, y+1))
	}
	return result
}

func (g *Grid) Set(p grid.Point, v int) error {
	if g == nil {
		return errors.New("grid is nil")
	}

	if p.Y >= g.height || p.X >= g.width {
		return ErrOutOfBounds
	}

	g.values[p.Y][p.X] = v
	return nil
}

func (g *Grid) MustSet(p grid.Point, v int) {
	if err := g.Set(p, v); err != nil {
		panic(err)
	}
}
