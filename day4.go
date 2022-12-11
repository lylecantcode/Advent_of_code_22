package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day4main() {
	lines := ReadInput("day4.txt")
	splitPattern := regexp.MustCompile(`[^\d]`)
	// split on anything that isn't a number
	part1count := 0
	part2count := 0
	for _, v := range lines {
		var values []int
		valuesStr := splitPattern.Split(v, 4)
		for i := 0; i < len(valuesStr); i++ {
			val, err := strconv.Atoi(valuesStr[i])
			if err != nil {
				panic(err)
			}
			values = append(values, val)
		}
		if len(values) == 4 {
			part1count += Day4Part1(values)
			part2count += Day4Part2(values)
		}
	}
	fmt.Println("day4 \n", "	part 1:", part1count, "\n	part 2:", part2count)
}

func Day4Part1(values []int) int {
	if (values[0] <= values[2] && values[1] >= values[3]) ||
		(values[0] >= values[2] && values[1] <= values[3]) {
		return 1
	}
	return 0
}

func Day4Part2(values []int) int {
	if (values[0] >= values[2] && values[0] <= values[3]) ||
		(values[1] >= values[2] && values[1] <= values[3]) ||
		(values[2] >= values[0] && values[2] <= values[1]) ||
		(values[3] >= values[0] && values[3] <= values[1]) {
		return 1
	}
	return 0
}
