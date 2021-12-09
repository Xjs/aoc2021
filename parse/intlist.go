package parse

import (
	"strconv"
	"strings"
)

// IntList takes a string with comma-separated integers (may be surrounded by whitespace) and returns the integers as a slice.
func IntList(s string) ([]int, error) {
	fields := strings.Split(s, ",")
	ns := make([]int, 0, len(fields))

	for _, field := range fields {
		n, err := strconv.Atoi(strings.TrimSpace(field))
		if err != nil {
			return ns, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

// DigitList takes a string with digits integers and returns the digits as a slice.
func DigitList(s string) ([]int, error) {
	ns := make([]int, 0, len(s))

	for _, r := range s {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			return ns, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}
