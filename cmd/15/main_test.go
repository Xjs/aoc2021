package main

import (
	"reflect"
	"testing"

	"github.com/Xjs/aoc2021/integer/grid"
)

func Test_pointPriorityQueue_addWithPriority(t *testing.T) {
	type fields struct {
		queue    []grid.Point
		priority map[grid.Point]int
	}
	type args struct {
		p        grid.Point
		priority int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   pointPriorityQueue
	}{
		{
			"test",
			fields{
				queue:    []grid.Point{grid.P(0, 0)},
				priority: map[grid.Point]int{grid.P(0, 0): 0},
			},
			args{
				p:        grid.P(0, 1),
				priority: 1,
			},
			pointPriorityQueue{
				queue:    []grid.Point{grid.P(0, 0), grid.P(0, 1)},
				priority: map[grid.Point]int{grid.P(0, 0): 0, grid.P(0, 1): 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &pointPriorityQueue{
				queue:    tt.fields.queue,
				priority: tt.fields.priority,
			}
			q.addWithPriority(tt.args.p, tt.args.priority)
			if !reflect.DeepEqual(*q, tt.want) {
				t.Errorf("got %v, want %v", q, tt.want)
			}
		})
	}
}
