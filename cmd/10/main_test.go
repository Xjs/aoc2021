package main

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		line []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]byte("{([(<{}[<>[]}>{[]{[(<()>")}, 1197},
		{"2", args{[]byte("[[<[([]))<([[{}[[()]]]")}, 3},
		{"3", args{[]byte("[{[{({}]{}}([{[{{{}}([]")}, 57},
		{"4", args{[]byte("[<(<(<(<{}))><([]([]()")}, 3},
		{"5", args{[]byte("<{([([[(<>()){}]>(<<{{")}, 25137},
		{"6", args{[]byte("[({(<(())[]>[[{[]{<()<>>")}, 0},
		{"7", args{[]byte("[(()[<>])]({[<{<<[]>>(")}, 0},
		{"8", args{[]byte("(((({<>}<{<{<>}{[]{[]{}")}, 0},
		{"9", args{[]byte("{<[[]]>}<{[{[{[]{()[[[]")}, 0},
		{"10", args{[]byte("<{([{{}}[<[[[<>{}]]]>[]]")}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := check(tt.args.line); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack(t *testing.T) {
	var s stack
	s.push('a')
	s.push('b')
	s.push('c')
	if v := s.pop(); v != 'c' {
		t.Errorf("pop() = %v, want 'c'", v)
	}
	if v := s.pop(); v != 'b' {
		t.Errorf("pop() = %v, want 'b'", v)
	}
	s.push('d')
	if v := s.pop(); v != 'd' {
		t.Errorf("pop() = %v, want 'd'", v)
	}
	if v := s.pop(); v != 'a' {
		t.Errorf("pop() = %v, want 'a'", v)
	}
}
