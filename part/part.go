package part

import "os"

// One returns true if the command line args are not the word "part2".
func One() bool {
	part1 := true
	if len(os.Args) > 1 && os.Args[1] == "part2" {
		part1 = false
	}
	return part1
}
