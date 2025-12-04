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
	return partOneNaive(input)
}

func partOneNaive(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sum := 0
	for _, line := range lines {
		digits := strings.Split(line, "")
		maxNum := 0
		maxIndex := 0
		maxString := ""
		for i, digit := range digits {
			if i == len(digits)-1 {
				break
			}
			intDigit, _ := strconv.Atoi(digit)
			if maxNum < intDigit {
				maxNum = intDigit
				maxIndex = i
				maxString = digit
			}
		}
		nextMaxNum := 0
		nextMaxString := ""
		//now loop through starting from maxIndex + 1
		for i := maxIndex + 1; i < len(digits); i++ {
			curr := digits[i]
			intDigit, _ := strconv.Atoi(curr)
			if nextMaxNum < intDigit {
				nextMaxNum = intDigit
				nextMaxString = curr
			}
		}

		sumStr := maxString + nextMaxString
		sumInt, _ := strconv.Atoi(sumStr)
		sum += sumInt

	}
	return sum
}

func partTwo(input string) int64 {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var sum int64
	const k = 12
	for _, line := range lines {
		n := len(line)
		digits := strings.Split(line, "")
		var lineVal int64
		start := 0
		remainder := k
		for remainder > 0 {
			end := n - remainder // inclusive
			digit, idx := findMaxDigit(start, end, digits)
			lineVal = lineVal*10 + int64(digit)

			start = idx + 1
			remainder--
		}
		sum += lineVal
	}
	return sum
}

func findMaxDigit(start, end int, digits []string) (int, int) {
	maxDigit := 0
	maxIndex := 0
	//loop from zero index to the last possible index in which we can still have a twelve v
	for i := start; i <= end; i++ {
		digit := digits[i]
		intDigit, _ := strconv.Atoi(digit)
		if maxDigit < intDigit {
			maxDigit = intDigit
			maxIndex = i
		}
	}
	return maxDigit, maxIndex
}
