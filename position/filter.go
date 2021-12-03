package position

import "math"

// Top returns the most common rune in the histogram.
// For equally common runes, returns the rune with highest significance.
func Top(hist map[rune]int) rune {
	var result rune
	var foundCount int
	for r, count := range hist {
		if count > foundCount {
			result = r
			foundCount = count
		} else if count == foundCount && r > result {
			result = r
		}
	}
	return result
}

// Bottom returns the least common rune in the histogram.
// For equally common runes, returns the rune with lowest significance.
func Bottom(hist map[rune]int) rune {
	var result rune
	foundCount := math.MaxInt
	for r, count := range hist {
		if count < foundCount {
			result = r
			foundCount = count
		} else if count == foundCount && r < result {
			result = r
		}
	}
	return result
}

// Filter deletes all strings from the input that do not have
// rune r in position position.
func Filter(input []string, position int, r rune) []string {
	var output []string
	for _, s := range input {
		if len(s) <= position {
			continue
		}
		if []rune(s)[position] != r {
			continue
		}
		output = append(output, s)
	}
	return output
}
