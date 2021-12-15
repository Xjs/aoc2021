package integer

import (
	"reflect"
	"testing"

	"github.com/Xjs/aoc2021/integer/grid"
)

func TestGrid_Environment4(t *testing.T) {
	type fields struct {
		width  uint
		height uint
		values [][]int
	}
	type args struct {
		p grid.Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []grid.Point
	}{
		{"regular", fields{42, 42, nil}, args{grid.P(5, 5)}, []grid.Point{grid.P(4, 5), grid.P(6, 5), grid.P(5, 4), grid.P(5, 6)}},
		{"lower bound", fields{42, 42, nil}, args{grid.P(0, 0)}, []grid.Point{grid.P(1, 0), grid.P(0, 1)}},
		{"upper bound", fields{42, 42, nil}, args{grid.P(41, 41)}, []grid.Point{grid.P(40, 41), grid.P(41, 40)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Grid{
				width:  tt.fields.width,
				height: tt.fields.height,
				values: tt.fields.values,
			}
			if got := g.Environment4(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.Environment4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_Environment8(t *testing.T) {
	type fields struct {
		width  uint
		height uint
		values [][]int
	}
	type args struct {
		p grid.Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []grid.Point
	}{
		{"regular", fields{42, 42, nil}, args{grid.P(5, 5)}, []grid.Point{
			grid.P(4, 5), grid.P(6, 5), grid.P(5, 4), grid.P(5, 6),
			grid.P(4, 4), grid.P(6, 6), grid.P(4, 6), grid.P(6, 4)},
		},
		{"regular", fields{42, 42, nil}, args{grid.P(10, 5)}, []grid.Point{
			grid.P(9, 5), grid.P(11, 5), grid.P(10, 4), grid.P(10, 6),
			grid.P(9, 4), grid.P(11, 6), grid.P(9, 6), grid.P(11, 4)},
		},
		{"lower bound", fields{42, 42, nil}, args{grid.P(0, 0)}, []grid.Point{grid.P(1, 0), grid.P(0, 1), grid.P(1, 1)}},
		{"upper bound", fields{42, 42, nil}, args{grid.P(41, 41)}, []grid.Point{grid.P(40, 41), grid.P(41, 40), grid.P(40, 40)}},
		{"edge top", fields{42, 42, nil}, args{grid.P(5, 0)}, []grid.Point{
			grid.P(4, 0), grid.P(6, 0), grid.P(5, 1),
			grid.P(6, 1), grid.P(4, 1)},
		},
		{"edge bottom", fields{42, 42, nil}, args{grid.P(5, 41)}, []grid.Point{
			grid.P(4, 41), grid.P(6, 41), grid.P(5, 40),
			grid.P(4, 40), grid.P(6, 40)},
		},
		{"edge left", fields{42, 42, nil}, args{grid.P(0, 5)}, []grid.Point{
			grid.P(1, 5), grid.P(0, 4), grid.P(0, 6),
			grid.P(1, 6), grid.P(1, 4)},
		},
		{"edge right", fields{42, 42, nil}, args{grid.P(41, 5)}, []grid.Point{
			grid.P(40, 5), grid.P(41, 4), grid.P(41, 6),
			grid.P(40, 4), grid.P(40, 6)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Grid{
				width:  tt.fields.width,
				height: tt.fields.height,
				values: tt.fields.values,
			}
			if got := g.Environment8(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.Environment8() = %v, want %v", got, tt.want)
			}
		})
	}
}
