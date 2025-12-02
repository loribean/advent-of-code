package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return partTwo(input)
	}
	// solve part 1 here
	return partOne(input)
}

func partOne(input string) int {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	invalidIDSum := 0
	for _, pair := range ranges {
		list := strings.Split(pair, "-")
		start, _ := strconv.Atoi(list[0])
		end, _ := strconv.Atoi(list[1])
		// iterate between the ranges and find the invalid ids
		// an id is invalid IF it is made up of the same digits repeated twice
		// so, split the number as a string and see if it equates to each other

		for i := start; i <= end; i++ {
			stringI := strconv.Itoa(i)
			//split a string into half
			length := len(stringI)
			if length%2 == 1 {
				continue
			}
			mid := length / 2
			string1 := stringI[0:mid]
			string2 := stringI[mid:]

			if string1 == string2 {
				invalidIDSum += i
			}
		}
	}
	return invalidIDSum
}

func partTwo(input string) int {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	invalidIDSum := 0
	for _, pair := range ranges {
		list := strings.Split(pair, "-")
		start, _ := strconv.Atoi(list[0])
		end, _ := strconv.Atoi(list[1])
		// iterate between the ranges and find the invalid ids
		// an id is invalid IF it is made up of the same digits repeated at least more than once (aka twice)
		for i := start; i <= end; i++ {
			if isRepeating(i) {
				invalidIDSum += i
			}
		}
	}
	return invalidIDSum
}

func isRepeating(id int) bool {
	s := strconv.Itoa(id)
	totalLen := len(s)

	for patternLength := 1; patternLength <= totalLen/2; patternLength++ {
		// if we can't cleanly divide total length by pattern length, then we can skip this pattern length
		if totalLen%patternLength != 0 {
			continue
		}

		potentialPattern := s[:patternLength]
		patternOccurence := totalLen / patternLength
		//reconstructed := ""
		//for i := 0; i < patternOccurence; i++ {
		//	reconstructed = reconstructed + potentialPattern
		//}
		reconstructed := strings.Repeat(potentialPattern, patternOccurence)
		if reconstructed == s {
			return true
		}

	}
	return false
}
