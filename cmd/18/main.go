package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Token int

const (
	Zero Token = -iota
	LeftBracket
	RightBracket
	Comma
)

type Number []Token

func Parse(s string) (Number, error) {
	var result Number
	var digits []rune

	dumpDigits := func() {
		if len(digits) > 0 {
			n, _ := strconv.Atoi(string(digits))
			result = append(result, Token(n))
			digits = nil
		}
	}

	for _, r := range s {
		switch r {
		case '[':
			// Actually would be invalid at this point but still, only tokenising here
			dumpDigits()
			result = append(result, LeftBracket)
		case ']':
			dumpDigits()
			result = append(result, RightBracket)
		case ',':
			dumpDigits()
			result = append(result, Comma)
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			digits = append(digits, r)
		default:
			return nil, fmt.Errorf("invalid rune %c", r)
		}
	}

	return result, result.WellFormed()
}

func (n Number) WellFormed() error {
	type marker struct {
		first, comma, last bool
	}

	var stack int
	var seen []marker
	for _, t := range n {
		if stack == 0 && t >= 0 {
			return fmt.Errorf("numbers must be a pair")
		}

		switch t {
		case LeftBracket:
			stack++
			seen = append(seen, marker{})
		case Comma:
			m := seen[stack-1]
			if !m.first {
				return fmt.Errorf("missing first at depth %d", stack)
			}
			seen[stack-1].comma = true
		case RightBracket:
			if stack == 0 {
				return fmt.Errorf("mismatched brackets")
			}
			m := seen[stack-1]
			if !m.first {
				return fmt.Errorf("missing first at depth %d", stack)
			}
			if !m.comma {
				return fmt.Errorf("missing comma at depth %d", stack)
			}
			if !m.last {
				return fmt.Errorf("missing last at depth %d", stack)
			}
			stack--
			if stack > 0 {
				if seen[stack-1].first && seen[stack-1].comma {
					seen[stack-1].last = true
				} else if seen[stack-1] == (marker{}) {
					seen[stack-1].first = true
				}
			}
		default:
			if t < 0 {
				return fmt.Errorf("invalid token %d", t)
			}

			if stack == 0 {
				return fmt.Errorf("numbers must be a pair")
			}

			if stack > 0 {
				if seen[stack-1].first && seen[stack-1].comma {
					seen[stack-1].last = true
				} else if seen[stack-1] == (marker{}) {
					seen[stack-1].first = true
				}
			}
		}
	}

	if stack > 0 {
		return fmt.Errorf("too few closing brackets")
	}
	return nil
}

func MustParse(s string) Number {
	n, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return n
}

func (n Number) Explode() Number {
	var stack int

	var result Number
	var addRight Token
	var haveAddRight bool

	var foundExplodingPair bool
	var addedLeft bool
	var skipNextRightBracket bool
	var done bool

	for i, t := range n {
		if done {
			result = append(result, t)
			continue
		}

		if foundExplodingPair {
			for j := i - 2; j >= 0; j-- {
				if result[j] >= 0 {
					result[j] += t
					t = 0
				}
			}
			addedLeft = true
			foundExplodingPair = false
			continue
		}

		if addedLeft {
			if t == Comma {
				continue
			}

			addRight = t
			haveAddRight = true
			addedLeft = false
			continue
		}

		switch t {
		case LeftBracket:
			stack++
		case RightBracket:
			stack--
			if skipNextRightBracket {
				skipNextRightBracket = false
				continue
			}
		}

		if stack == 5 && !haveAddRight {
			foundExplodingPair = true
			skipNextRightBracket = true
			result = append(result, 0)
			continue
		}

		if t >= 0 && haveAddRight {
			result = append(result, t+addRight)
			addRight = 0
			haveAddRight = false
			done = true
			continue
		}
		result = append(result, t)
	}
	return result
}

func (n Number) Split() Number {
	var result Number
	var done bool
	for _, t := range n {
		if t >= 10 && !done {
			left := t / 2
			right := t - left
			result = append(result, LeftBracket, left, Comma, right, RightBracket)
			done = true
			continue
		}
		result = append(result, t)
	}
	return result
}

func (n Number) Reduce() Number {
	var iter int
	for {
		previous := n
		n = n.Explode()
		if err := n.WellFormed(); err != nil {
			panic(fmt.Errorf("%d: after expldode: %v (was %s)", iter, err, previous))
		}
		if !n.Equals(previous) {
			continue
		}
		n = n.Split()
		if err := n.WellFormed(); err != nil {
			panic(fmt.Errorf("%d: after split: %v", iter, err))
		}
		if n.Equals(previous) {
			return n
		}
		iter++
	}
}

func (n Number) PlainAdd(n2 Number) Number {
	result := make(Number, 0, len(n)+len(n2)+3)
	result = append(result, LeftBracket)
	result = append(result, n...)
	result = append(result, Comma)
	result = append(result, n2...)
	result = append(result, RightBracket)
	return result
}

func (n Number) Add(n2 Number) Number {
	return n.PlainAdd(n2).Reduce()
}

func AddList(n Number, rest ...Number) Number {
	for _, n2 := range rest {
		n = n.Add(n2)
	}
	return n
}

func Sum(n ...Number) Number {
	if len(n) == 0 {
		return nil
	}
	num, rest := n[0], n[1:]

	for _, n2 := range rest {
		num = num.Add(n2)
	}
	return num
}

func MagnitudeSimple(n Number) Number {
	var result Number
	var skip int
	for i, t := range n {
		if skip > 0 {
			skip--
			continue
		}
		result = append(result, t)
		if t == Comma {
			left, right := n[i-1], n[i+1]
			if left < 0 || right < 0 {
				continue
			}
			result = append(result[:len(result)-3], left*3+right*2)
			skip = 2
			continue
		}
	}
	return result
}

func Magnitude(n Number) int {
	for {
		n = MagnitudeSimple(n)
		if len(n) == 1 {
			return int(n[0])
		}
	}
}

func (n Number) String() string {
	var b strings.Builder
	for _, t := range n {
		switch t {
		case LeftBracket:
			b.WriteRune('[')
		case RightBracket:
			b.WriteRune(']')
		case Comma:
			b.WriteRune(',')
		default:
			b.WriteString(strconv.Itoa(int(t)))
		}
	}
	return b.String()
}

func (n Number) Equals(n2 Number) bool {
	if len(n) != len(n2) {
		return false
	}
	for i, t := range n {
		if t != n2[i] {
			return false
		}
	}
	return true
}

func main() {
	var list []Number
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := s.Text()
		n, err := Parse(text)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(Magnitude(Sum(list...)))

	var maxMag int
	for _, n := range list {
		for _, n2 := range list {
			mag := Magnitude(n.Add(n2))
			if mag > maxMag {
				maxMag = mag
			}
		}
	}

	fmt.Println(maxMag)
}
