package integer

// Sum returns the sum of the list of integers
func Sum(ns []int) int {
	var sum int
	for _, n := range ns {
		sum += n
	}
	return sum
}
