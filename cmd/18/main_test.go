package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    Number
		wantErr bool
	}{
		{"[1,2]", "[1,2]", Number{LeftBracket, 1, Comma, 2, RightBracket}, false},
		{"[[1,2],3]", "[[1,2],3]", Number{LeftBracket, LeftBracket, 1, Comma, 2, RightBracket, Comma, 3, RightBracket}, false},
		{"[9,[8,7]]", "[9,[8,7]]", Number{LeftBracket, 9, Comma, LeftBracket, 8, Comma, 7, RightBracket, RightBracket}, false},
		{"[[1,9],[8,5]]", "[[1,9],[8,5]]", Number{LeftBracket, LeftBracket, 1, Comma, 9, RightBracket, Comma, LeftBracket, 8, Comma, 5, RightBracket, RightBracket}, false},
		{"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]", "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]", Number{LeftBracket, LeftBracket, LeftBracket, LeftBracket, 1, Comma, 2, RightBracket, Comma, LeftBracket, 3, Comma, 4, RightBracket, RightBracket, Comma, LeftBracket, LeftBracket, 5, Comma, 6, RightBracket, Comma, LeftBracket, 7, Comma, 8, RightBracket, RightBracket, RightBracket, Comma, 9, RightBracket}, false},
		{"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]", "[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]", Number{LeftBracket, LeftBracket, LeftBracket, 9, Comma, LeftBracket, 3, Comma, 8, RightBracket, RightBracket, Comma, LeftBracket, LeftBracket, 0, Comma, 9, RightBracket, Comma, 6, RightBracket, RightBracket, Comma, LeftBracket, LeftBracket, LeftBracket, 3, Comma, 7, RightBracket, Comma, LeftBracket, 4, Comma, 9, RightBracket, RightBracket, Comma, 3, RightBracket, RightBracket}, false},
		{"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]", "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]", Number{LeftBracket, LeftBracket, LeftBracket, LeftBracket, 1, Comma, 3, RightBracket, Comma, LeftBracket, 5, Comma, 3, RightBracket, RightBracket, Comma, LeftBracket, LeftBracket, 1, Comma, 3, RightBracket, Comma, LeftBracket, 8, Comma, 7, RightBracket, RightBracket, RightBracket, Comma, LeftBracket, LeftBracket, LeftBracket, 4, Comma, 9, RightBracket, Comma, LeftBracket, 6, Comma, 9, RightBracket, RightBracket, Comma, LeftBracket, LeftBracket, 8, Comma, 2, RightBracket, Comma, LeftBracket, 7, Comma, 3, RightBracket, RightBracket, RightBracket, RightBracket}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_Explode(t *testing.T) {
	tests := []struct {
		name string
		n    Number
		want Number
	}{
		{"ex1", MustParse("[[[[[9,8],1],2],3],4]"), MustParse("[[[[0,9],2],3],4]")},
		{"ex2", MustParse("[7,[6,[5,[4,[3,2]]]]]"), MustParse("[7,[6,[5,[7,0]]]]")},
		{"ex3", MustParse("[[6,[5,[4,[3,2]]]],1]"), MustParse("[[6,[5,[7,0]]],3]")},
		{"ex4", MustParse("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"), MustParse("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")},
		{"ex5", MustParse("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"), MustParse("[[3,[2,[8,0]]],[9,[5,[7,0]]]]")},
		{"ex-process-1", MustParse("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"), MustParse("[[[[0,7],4],[7,[[8,4],9]]],[1,1]]")},
		{"ex-process-2", MustParse("[[[[0,7],4],[7,[[8,4],9]]],[1,1]]"), MustParse("[[[[0,7],4],[15,[0,13]]],[1,1]]")},
		{"ex-process-3", MustParse("[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"), MustParse("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")},
		{"ex-failing", MustParse("[[[[4,0],[5,4]],[[7,7],[6,0]]],[[[6,6],[5,0]],[[6,6],[8,[[5,6],8]]]]]"), MustParse("[0,0]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Explode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Number.Explode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_Split(t *testing.T) {
	tests := []struct {
		name string
		n    Number
		want Number
	}{
		{"ex-process-1", MustParse("[[[[0,7],4],[15,[0,13]]],[1,1]]"), MustParse("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]")},
		{"ex-process-2", MustParse("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"), MustParse("[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Split(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Number.Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_Reduce(t *testing.T) {
	tests := []struct {
		name string
		n    Number
		want Number
	}{
		{"example", MustParse("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"), MustParse("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Reduce(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Number.Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_PlainAdd(t *testing.T) {
	tests := []struct {
		name string
		n    Number
		n2   Number
		want Number
	}{
		{"ex1", MustParse("[1,2]"), MustParse("[[3,4],5]"), MustParse("[[1,2],[[3,4],5]]")},
		{"ex2", MustParse("[[[[4,3],4],4],[7,[[8,4],9]]]"), MustParse("[1,1]"), MustParse("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.PlainAdd(tt.n2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Number.PlainAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_Add(t *testing.T) {
	tests := []struct {
		name string
		n    Number
		n2   Number
		want Number
	}{
		{"ex1", MustParse("[1,2]"), MustParse("[[3,4],5]"), MustParse("[[1,2],[[3,4],5]]")},
		{"ex1b", MustParse("[[[[4,3],4],4],[7,[[8,4],9]]]"), MustParse("[1,1]"), MustParse("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")},
		{"ex2", MustParse("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]"), MustParse("[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"), MustParse("[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]")},
		{"ex3", MustParse("[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"), MustParse("[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]"), MustParse("[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]")},
		{"ex4", MustParse("[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]"), MustParse("[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]"), MustParse("[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]")},
		{"ex5", MustParse("[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]"), MustParse("[7,[5,[[3,8],[1,4]]]]"), MustParse("[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]")},
		{"ex6", MustParse("[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]"), MustParse("[[2,[2,2]],[8,[8,1]]]"), MustParse("[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]")},
		{"ex7", MustParse("[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]"), MustParse("[2,9]"), MustParse("[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]")},
		{"ex8", MustParse("[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]"), MustParse("[1,[[[9,3],9],[[9,0],[0,7]]]]"), MustParse("[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]")},
		{"ex9", MustParse("[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]"), MustParse("[[[5,[7,4]],7],1]"), MustParse("[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]")},
		{"ex10", MustParse("[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]"), MustParse("[[[[4,2],2],6],[8,7]]"), MustParse("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Add(tt.n2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Number.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddList(t *testing.T) {
	type args []Number
	tests := []struct {
		name string
		args args
		want Number
	}{
		{"ex1", args{
			MustParse("[1,1]"),
			MustParse("[2,2]"),
			MustParse("[3,3]"),
			MustParse("[4,4]"),
		}, MustParse("[[[[1,1],[2,2]],[3,3]],[4,4]]")},
		{"ex2", args{
			MustParse("[1,1]"),
			MustParse("[2,2]"),
			MustParse("[3,3]"),
			MustParse("[4,4]"),
			MustParse("[5,5]"),
		}, MustParse("[[[[3,0],[5,3]],[4,4]],[5,5]]")},
		{"ex3", args{
			MustParse("[1,1]"),
			MustParse("[2,2]"),
			MustParse("[3,3]"),
			MustParse("[4,4]"),
			MustParse("[5,5]"),
			MustParse("[6,6]"),
		}, MustParse("[[[[5,0],[7,4]],[5,5]],[6,6]]")},
		{"large-example", args{
			MustParse("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]"),
			MustParse("[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"),
			MustParse("[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]"),
			MustParse("[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]"),
			MustParse("[7,[5,[[3,8],[1,4]]]]"),
			MustParse("[[2,[2,2]],[8,[8,1]]]"),
			MustParse("[2,9]"),
			MustParse("[1,[[[9,3],9],[[9,0],[0,7]]]]"),
			MustParse("[[[5,[7,4]],7],1]"),
			MustParse("[[[[4,2],2],6],[8,7]]"),
		}, MustParse("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")},
		{"another-example", args{
			MustParse("[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]"),
			MustParse("[[[5,[2,8]],4],[5,[[9,9],0]]]"),
			MustParse("[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]"),
			MustParse("[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]"),
			MustParse("[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]"),
			MustParse("[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]"),
			MustParse("[[[[5,4],[7,7]],8],[[8,3],8]]"),
			MustParse("[[9,3],[[9,9],[6,[4,9]]]]"),
			MustParse("[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]"),
			MustParse("[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]"),
		}, MustParse("[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddList(tt.args[0], tt.args[1:]...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		arg  string
		want int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
		{"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", 4140},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := Magnitude(MustParse(tt.arg)); got != tt.want {
				t.Errorf("Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagnitudeSimple(t *testing.T) {
	tests := []struct {
		name string
		args Number
		want Number
	}{
		{"simple", MustParse("[[1,2],[[3,4],5]]"), MustParse("[7,[17,5]]")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MagnitudeSimple(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MagnitudeSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}
