package position

import (
	"reflect"
	"testing"
)

var aocExample = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestHistogramOfList(t *testing.T) {
	type args struct {
		l []string
	}
	tests := []struct {
		name string
		args args
		want map[int]map[rune]int
	}{
		{
			"aoc_example",
			args{aocExample},
			map[int]map[rune]int{
				0: {'1': 7, '0': 5},
				1: {'1': 5, '0': 7},
				2: {'1': 8, '0': 4},
				3: {'1': 7, '0': 5},
				4: {'1': 5, '0': 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HistogramOfList(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HistogramOfList() = %v, want %v", got, tt.want)
			}
		})
	}
}
