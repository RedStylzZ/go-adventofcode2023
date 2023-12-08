package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func charIsNum(c rune) bool {
	return c >= '0' && c <= '9'
}

func isBetween(x, y, idx int) bool {
	return idx >= x && idx <= y
}

func toIntArr(input []string) []int {
	arr := make([]int, len(input))

	for i := range input {
		x, _ := strconv.Atoi(input[i])
		arr[i] = x
	}
	return arr
}

func partOne(lines []string) int {
	var sum int

	for _, line := range lines {
		input := strings.Split(line, ": ")[1]
		cards := strings.Split(input, " | ")
		cards[0] = strings.TrimSpace(cards[0])
		cards[1] = strings.TrimSpace(cards[1])
		winning := toIntArr(strings.Split(cards[0], " "))
		myCards := toIntArr(strings.Split(cards[1], " "))

		fmt.Printf("%+v | %+v\n", winning, myCards)
	}

	return sum
}

func partTwo(lines []string) int {
	var sum int

	return sum
}

func main() {
	b, err := os.ReadFile("input/day04_test")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	ansOne := partOne(lines)
	fmt.Printf("PartOne: %d\n", ansOne)

	ansTwo := partTwo(lines)
	fmt.Printf("PartTwo: %d\n", ansTwo)
}
