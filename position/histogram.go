package position

// Histogram takes strings and creates a histogram of read runes by position in the string.
func Histogram(c <-chan string) (result map[int]map[rune]int) {
	result = make(map[int]map[rune]int)
	for input := range c {
		for i, r := range input {
			if result[i] == nil {
				result[i] = make(map[rune]int)
			}
			result[i][r]++
		}
	}
	return result
}

// HistogramOfList is Histogram applied to a list of strings as a convenience function.
func HistogramOfList(l []string) map[int]map[rune]int {
	c := make(chan string)
	go func() {
		for _, s := range l {
			c <- s
		}
		close(c)
	}()
	return Histogram(c)
}
