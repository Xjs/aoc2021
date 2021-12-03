package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type position struct {
	horizontalPosition, depth, aim int
}

// forward X does two things:
// It increases your horizontal position by X units.
// It increases your depth by your aim multiplied by X.
func (p *position) forward(units int) {
	p.horizontalPosition += units
	p.depth += p.aim * units
}

// down X increases the aim by X units.
func (p *position) down(units int) {
	p.aim += units
}

// up X decreases the aim by X units.
func (p *position) up(units int) {
	p.aim -= units
}

func (p *position) parse(command []string) error {
	if len(command) != 2 {
		return errors.New("command must contain 2 words")
	}

	argument, err := strconv.Atoi(command[1])
	if err != nil {
		return fmt.Errorf("error parsing argument (2nd word) as integer: %w", err)
	}

	switch command[0] {
	case "forward":
		p.forward(argument)
	case "up":
		p.up(argument)
	case "down":
		p.down(argument)
	default:
		return fmt.Errorf("unknown command %q", command[0])
	}

	return nil
}

func main() {
	cr := csv.NewReader(os.Stdin)
	cr.Comma = ' '
	commands, err := cr.ReadAll()
	if err != nil {
		log.Fatalf("Error reading commands: %v", err)
	}

	p := new(position)

	for _, command := range commands {
		if err := p.parse(command); err != nil {
			log.Fatalf("Error parsing command %v: %v", command, err)
		}
	}

	fmt.Println("horizontalPosition * depth", p.horizontalPosition*p.depth)
}
