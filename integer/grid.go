package integer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/Xjs/aoc2021/integer/grid"
	"github.com/Xjs/aoc2021/parse"
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

// Width returns the grid's width.
func (g Grid) Width() uint {
	return g.width
}

// Height returns the grid's height.
func (g Grid) Height() uint {
	return g.height
}

// NewGrid creates a new zero-filled grid with the given dimensions.
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

// ReadGrid reads digit lists from r until EOF is encountered,
// and creates a grid from them.
func ReadGrid(r io.Reader) (*Grid, error) {
	var values [][]int
	s := bufio.NewScanner(r)
	for s.Scan() {
		ds, err := parse.DigitList(s.Text())
		if err != nil {
			return nil, err
		}
		values = append(values, ds)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	g, err := GridFrom(values)
	return &g, err
}

// ErrOutOfBounds is returned by At and Set if an out-of-bounds coordinate is accessed.
var ErrOutOfBounds = errors.New("out of bounds access to grid")

// At returns the value at the given point. It returns ErrOutOfBounds if
// an out-of-bounds point is attempted to be read.
func (g Grid) At(p grid.Point) (int, error) {
	if p.Y >= g.height || p.X >= g.width {
		return 0, ErrOutOfBounds
	}
	return g.values[p.Y][p.X], nil
}

// MustAt is At, but panics instead of returning an error.
func (g Grid) MustAt(p grid.Point) int {
	v, err := g.At(p)
	if err != nil {
		panic(err)
	}
	return v
}

// Environment4 returns a slice of points that represent the 4-environment
// of p, i. e. the points to the left, right, top and bottom. Any points would be
// out of bounds are not returned.
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

// Environment8 returns a slice of points that represent the 8-environment
// of p, i. e. the points to the left, right, top and bottom, and all diagonals.
//  Any points would be out of bounds are not returned.
func (g Grid) Environment8(p grid.Point) []grid.Point {
	result := make([]grid.Point, 0, 8)
	result = append(result, g.Environment4(p)...)

	x, y := p.X, p.Y
	if x > 0 && y > 0 {
		result = append(result, grid.P(x-1, y-1))
	}
	if x < g.width-1 && y < g.height-1 {
		result = append(result, grid.P(x+1, y+1))
	}
	if x > 0 && y < g.height-1 {
		result = append(result, grid.P(x-1, y+1))
	}
	if x < g.width-1 && y > 0 {
		result = append(result, grid.P(x+1, y-1))
	}
	return result
}

// Set sets the given grid point to the given value. It returns ErrOutOfBounds if
// an out-of-bounds point is attempted to be set.
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

// MustSet is Set, but panics instead of returning an error.
func (g *Grid) MustSet(p grid.Point, v int) {
	if err := g.Set(p, v); err != nil {
		panic(err)
	}
}

// String creates a multi-line string from the grid.
func (g Grid) String() string {
	var b strings.Builder
	max := 0
	for x := uint(0); x < g.width; x++ {
		for y := uint(0); y < g.height; y++ {
			if v := g.MustAt(grid.P(x, y)); v > max {
				max = v
			}
		}
	}

	l := len(fmt.Sprint(max))
	sep := ""
	fill := ' '
	if l > 1 {
		sep = " "
	}

	for x := uint(0); x < g.width; x++ {
		for y := uint(0); y < g.height; y++ {
			v := g.MustAt(grid.P(x, y))
			rep := fmt.Sprint(v)
			for i := 0; i < l-len(rep); i++ {
				b.WriteRune(fill)
			}
			b.WriteString(rep)
			b.WriteString(sep)
		}
		b.WriteRune('\n')
	}

	return b.String()
}
