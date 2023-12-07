package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var digMap = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func checkNumbers(line string) rune {
	for k, v := range digMap {
		digLen := len(k)

		if len(line) >= digLen && line[0:digLen] == k {
			return v
		}
	}

	return '0'
}

func partTwo(lines []string) int {
	var sum int

	for _, line := range lines {
		buf := make([]rune, 0, len(line))

		for i, r := range line {
			if isNumber(r) {
				buf = append(buf, r)
				continue
			}

			num := checkNumbers(line[i:])
			if num > '0' {
				buf = append(buf, num)
			}
		}

		buf = append([]rune{buf[0]}, buf[len(buf)-1])

		i, err := strconv.Atoi(string(buf))
		if err != nil {
			log.Fatal(err)
		}

		sum += i
	}

	return sum
}

func partOne(lines []string) int {
	var sum int

	for _, line := range lines {
		buf := make([]rune, 0, len(line))

		for _, r := range line {
			if isNumber(r) {
				buf = append(buf, r)
			}
		}

		buf = append([]rune{buf[0]}, buf[len(buf)-1])

		i, err := strconv.Atoi(string(buf))
		if err != nil {
			log.Fatal(err)
		}

		sum += i
	}

	return sum
}

func main() {
	b, err := os.ReadFile("input/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	ansOne := partOne(lines)
	fmt.Printf("PartOne: %d\n", ansOne)

	ansTwo := partTwo(lines)
	fmt.Printf("PartTwo: %d\n", ansTwo)
}
