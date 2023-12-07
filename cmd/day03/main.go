package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numRex = regexp.MustCompile(`(\d+)`)

func charIsNum(c rune) bool {
	return c >= '0' && c <= '9'
}

func isBetween(x, y, idx int) bool {
	return idx >= x && idx <= y
}

func rangeInRange(idx, x, y, idxStart, idxEnd int) bool {
	return isBetween(x, y, idxStart) || isBetween(x, y, idxEnd) || isBetween(idxStart, idxEnd, idx)
}

func checkRangePartOne(line string, y int) int {
	var sum int

	results := numRex.FindAllStringIndex(line, -1)
	for _, rng := range results {
		if !rangeInRange(y, y-1, y+1, rng[0], rng[1]-1) {
			continue
		}

		i, err := strconv.Atoi(line[rng[0]:rng[1]])
		if err != nil {
			fmt.Printf("Invalid number (%d|%d): %s\n%s\n", rng[0], rng[1]-1, line[rng[0]:rng[1]], line)
			continue
		}
		sum += i
	}

	return sum
}

func checkRangePartTwo(line string, y int) (int, int) {
	var sum, count int

	results := numRex.FindAllStringIndex(line, -1)
	for _, rng := range results {
		if !rangeInRange(y, y-1, y+1, rng[0], rng[1]-1) {
			continue
		}

		i, err := strconv.Atoi(line[rng[0]:rng[1]])
		if err != nil {
			fmt.Printf("Invalid number (%d|%d): %s\n%s\n", rng[0], rng[1]-1, line[rng[0]:rng[1]], line)
			continue
		}

		if sum == 0 {
			sum = i
		} else {
			sum *= i
		}
		count++
	}

	return sum, count
}

func partOneCalc(lines []string, x, y int) int {
	var sum int

	start := x - 1
	if start < 0 {
		start = 0
	}

	end := x + 1
	if end > len(lines) {
		end = len(lines)
	}

	for l := start; l <= end; l++ {
		sum += checkRangePartOne(lines[l], y)
	}

	return sum
}

func partTwoCalc(lines []string, x, y int) int {
	var sum, count int

	start := x - 1
	if start < 0 {
		start = 0
	}

	end := x + 1
	if end > len(lines) {
		end = len(lines)
	}

	for l := start; l <= end; l++ {
		i, checkCount := checkRangePartTwo(lines[l], y)
		if i > 0 {
			if sum == 0 {
				sum = i
			} else {
				sum *= i
			}
		}
		count += checkCount
	}

	if count < 2 {
		return 0
	}

	return sum
}

func partOne(lines []string) int {
	var sum int

	for x, line := range lines {
		for y, c := range line {
			if c == '.' || charIsNum((c)) {
				continue
			}

			i := partOneCalc(lines, x, y)
			sum += i
		}
	}

	return sum
}

func partTwo(lines []string) int {
	var sum int

	for x, line := range lines {
		for y, c := range line {
			if c != '*' {
				continue
			}

			i := partTwoCalc(lines, x, y)
			sum += i
		}
	}

	return sum
}

func main() {
	b, err := os.ReadFile("input/day03.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	ansOne := partOne(lines)
	fmt.Printf("PartOne: %d\n", ansOne)

	ansTwo := partTwo(lines)
	fmt.Printf("PartTwo: %d\n", ansTwo)
}
