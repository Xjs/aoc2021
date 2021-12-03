package position

import (
	"reflect"
	"testing"
)

func TestTop(t *testing.T) {
	type args struct {
		hist map[rune]int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"top0", args{HistogramOfList(aocExample)[0]}, '1'},
		{"top1", args{HistogramOfList(aocExample)[1]}, '0'},
		{"top2", args{HistogramOfList(aocExample)[2]}, '1'},
		{"top3", args{HistogramOfList(aocExample)[3]}, '1'},
		{"top4", args{HistogramOfList(aocExample)[4]}, '0'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Top(tt.args.hist); got != tt.want {
				t.Errorf("Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBottom(t *testing.T) {
	type args struct {
		hist map[rune]int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"bottom0", args{HistogramOfList(aocExample)[0]}, '0'},
		{"bottom1", args{HistogramOfList(aocExample)[1]}, '1'},
		{"bottom2", args{HistogramOfList(aocExample)[2]}, '0'},
		{"bottom3", args{HistogramOfList(aocExample)[3]}, '0'},
		{"bottom4", args{HistogramOfList(aocExample)[4]}, '1'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bottom(tt.args.hist); got != tt.want {
				t.Errorf("Bottom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		input    []string
		position int
		r        rune
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"aoc_example_filter_top", args{aocExample, 0, '1'}, []string{
			"11110",
			"10110",
			"10111",
			"10101",
			"11100",
			"10000",
			"11001",
		}},
		{"aoc_example_filter_bottom", args{aocExample, 0, '0'}, []string{
			"00100",
			"01111",
			"00111",
			"00010",
			"01010",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.input, tt.args.position, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
