package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func partOne(lines []string) int {
	var sum int

	return sum
}

func partTwo(lines []string) int {
	var sum int

	return sum
}

func main() {
	b, err := os.ReadFile("input/dayXX.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	ansOne := partOne(lines)
	fmt.Printf("PartOne: %d\n", ansOne)

	ansTwo := partTwo(lines)
	fmt.Printf("PartTwo: %d\n", ansTwo)
}
