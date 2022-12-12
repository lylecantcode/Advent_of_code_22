package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func Day5main() {
	lines := ReadInput("day5.txt")
	var cargoArray [][]string
	inputStart := 0
	for i, v := range lines {
		if v[:3] == " 1 " {
			inputStart = i + 2
			break
		}
		row := make([][]string, 1)
		for j := 0; j <= len(v)-2; j++ {
			if v[j:j+3] == "   " {
				row[0] = append(row[0], "")
				j += 3
			} else if v[j:j+1] == "[" {
				row[0] = append(row[0], v[j+1:j+2])
				j += 2
			}
		}
		// reversed order so more levels can be stacked on by appending to end, rather than having to add to beginning
		cargoArray = append(row, cargoArray...)
	}

	maxHeight := len(cargoArray)
	splitPattern := regexp.MustCompile(`[^\d]+`)
	for _, v := range lines[inputStart : len(lines)-1] {
		var values []int
		valuesStr := splitPattern.Split(v, 4)

		for i := 1; i < len(valuesStr); i++ {
			val, err := strconv.Atoi(valuesStr[i])
			if err != nil {
				log.Println(err)
			}
			values = append(values, val)
			// count, from (1 based index), to
		}
		fmt.Println(values)
		from := findLast(cargoArray, values[1]-1)
		to := findLast(cargoArray, values[2]-1)
		fmt.Println(from, to)
		newHeight := to + values[0]
		if newHeight >= maxHeight {
			for i := maxHeight; i < newHeight; i++ {
				newRow := make([]string, len(cargoArray[0]))
				cargoArray = append(cargoArray, newRow)
			}
			maxHeight = newHeight
		}
	}

	fmt.Println("day5 \n", "	part 1:", "", "\n	part 2:", cargoArray, len(cargoArray))
}

func findLast(crates [][]string, row int) int {
	for i := len(crates) - 1; i >= 0; i-- {
		if crates[i][row] != "" {
			return i
		}
	}
	return 0
}
