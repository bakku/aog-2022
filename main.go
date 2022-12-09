package main

import (
	"log"
	"os"

	"bakku.dev/aog2022/puzzles"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day5.txt")
	if err != nil {
		log.Fatal("input file could not be read")
	}

	puzzles.Day5(string(bytes))
}
