package integer

import "math"

// Sum returns the sum of the list of integers
func Sum(ns []int) int {
	var sum int
	for _, n := range ns {
		sum += n
	}
	return sum
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(xs []int) int {
	max := math.MinInt
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}
