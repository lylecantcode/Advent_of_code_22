package main

import (
	"os"
	"strings"
)

func ReadInput(input string) []string {
	inputFile, err := os.ReadFile(input) //"Coding Advent day 10 input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(inputFile), "\n")
}
