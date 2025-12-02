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
	lines := strings.Split(strings.TrimSpace(input), "\n")
	position := 50
	zeroCount := 0
	for _, line := range lines {
		direction := line[0]
		magnitude, _ := strconv.Atoi(line[1:])
		if direction == 'R' {
			position += magnitude
			position %= 100
		} else if direction == 'L' {
			position -= magnitude
			position %= 100
			if position < 0 {
				position += 100
			}
		}
		if position == 0 {
			zeroCount++
		}
	}
	return zeroCount
}

func partTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	position := 50
	zeroCount := 0
	for _, line := range lines {
		direction := line[0]
		magnitude, _ := strconv.Atoi(line[1:])
		step := 1
		if direction == 'L' {
			step = -1
		}
		for i := 0; i < magnitude; i++ {
			position += step
			position %= 100
			if position < 0 {
				position += 100
			}

			if position == 0 {
				zeroCount++
			}
		}
	}
	return zeroCount
}
