package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Xjs/aoc2021/part"
)

const Size = 5

type Board struct {
	Numbers [Size][Size]int
	Marks   [Size][Size]bool
	won     bool
}

func (b *Board) bingoRow(row int) bool {
	for i := 0; i < Size; i++ {
		if !b.Marks[row][i] {
			return false
		}
	}
	return true
}

func (b *Board) bingoColumn(column int) bool {
	for i := 0; i < Size; i++ {
		if !b.Marks[i][column] {
			return false
		}
	}
	return true
}

// Mark marks the given number on the board and returns true if this resulted in a Bingo.
// If the board is already won, it will not be marked further.
func (b *Board) Mark(n int) bool {
	if b.won {
		return true
	}
	type mark struct{ row, column int }
	marks := make([]mark, 0, 1)
	for row := 0; row < Size; row++ {
		for column := 0; column < Size; column++ {
			if b.Numbers[row][column] == n && !b.Marks[row][column] {
				marks = append(marks, mark{row, column})
			}
		}
	}

	for _, mark := range marks {
		b.Marks[mark.row][mark.column] = true
	}

	for _, mark := range marks {
		if b.bingoRow(mark.row) || b.bingoColumn(mark.column) {
			b.won = true
			return true
		}
	}

	return false
}

func ReadBoard(s *bufio.Scanner) (Board, error) {
	var b Board
	var row int
	var eof bool = true
	for s.Scan() {
		text := s.Text()
		eof = false
		if text == "" {
			break
		}
		for col, word := range strings.Fields(text) {
			if col >= Size {
				return b, errors.New("input too large for bingo board")
			}
			word = strings.TrimSpace(word)
			n, err := strconv.Atoi(word)
			if err != nil {
				return b, err
			}
			b.Numbers[row][col] = n
		}
		row++
	}
	if eof {
		return b, io.EOF
	}
	return b, s.Err()
}

func (b *Board) SumUnmarked() int {
	var sum int
	for row := 0; row < Size; row++ {
		for column := 0; column < Size; column++ {
			if b.Marks[row][column] {
				continue
			}
			sum += b.Numbers[row][column]
		}
	}
	return sum
}

func main() {
	var inputs []int
	s := bufio.NewScanner(os.Stdin)
	if !s.Scan() {
		log.Fatal("no more inputs")
	}
	for _, word := range strings.Split(s.Text(), ",") {
		n, err := strconv.Atoi(word)
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	if !s.Scan() {
		log.Fatal("No more inputs")
	} else {
		s.Text()
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	var boards []*Board
	for {
		b, err := ReadBoard(s)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		boards = append(boards, &b)
	}

	winsAt := make([]int, len(boards))

	for i, input := range inputs {
		for j, board := range boards {
			if board.Mark(input) {
				if part.One() {
					log.Printf("Board %d wins after input %d: %d", j, i, input)
					fmt.Println(board.String())
					log.Println(board.SumUnmarked() * input)
					os.Exit(0)
				}
				if winsAt[j] == 0 {
					winsAt[j] = i
				}
			}
		}
	}

	var lastBoard, lastTurn int
	for board, turn := range winsAt {
		if turn > lastTurn {
			lastTurn = turn
			lastBoard = board
		}
	}

	log.Printf("Board %d wins after input %d: %d", lastBoard, lastTurn, inputs[lastTurn])
	fmt.Println(boards[lastBoard].String())
	log.Println(boards[lastBoard].SumUnmarked() * inputs[lastTurn])
}

func (b *Board) String() string {
	var s strings.Builder
	for row, r := range b.Numbers {
		for column, n := range r {
			if n < 10 {
				s.WriteString(" ")
			}
			s.WriteString(strconv.Itoa(n))
			if b.Marks[row][column] {
				s.WriteString("X")
			} else {
				s.WriteString(" ")
			}
			s.WriteString(" ")
		}
		s.WriteString("\n")
	}
	return s.String()
}
