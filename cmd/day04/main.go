package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func toIntArr(input []string) []int {
	arr := make([]int, len(input))

	for i := range input {
		x, _ := strconv.Atoi(input[i])
		arr[i] = x
	}
	return arr
}

func isInList(winning []int, num int) bool {
	for _, i := range winning {
		if i == num {
			return true
		}
	}

	return false
}

var cardRe = regexp.MustCompile(`(\d+)`)

func partOne(lines []string) int {
	var sum int

	for _, line := range lines {
		input := strings.Split(line, ": ")[1]
		cards := strings.Split(input, " | ")
		cards[0] = strings.TrimSpace(cards[0])
		cards[1] = strings.TrimSpace(cards[1])

		winning := toIntArr(cardRe.FindAllString(cards[0], -1))
		myCards := toIntArr(cardRe.FindAllString(cards[1], -1))

		var multiplier int
		for _, i := range myCards {
			if isInList(winning, i) {
				if multiplier == 0 {
					multiplier = 1
				} else {
					multiplier *= 2
				}
			}
		}

		sum += multiplier
	}

	return sum
}

type Game struct {
	ID      int
	Winning []int
	Cards   []int
}

func parseLine(line string) ([]int, []int) {
	input := strings.Split(line, ": ")[1]
	cards := strings.Split(input, " | ")
	cards[0] = strings.TrimSpace(cards[0])
	cards[1] = strings.TrimSpace(cards[1])

	winning := toIntArr(cardRe.FindAllString(cards[0], -1))
	myCards := toIntArr(cardRe.FindAllString(cards[1], -1))

	return winning, myCards
}

func partTwo(lines []string) int {
	var sum int

	m := make(map[int]Game)
	test := make(map[int]int)
	for x, line := range lines {
		winning, myCards := parseLine(line)
		m[x] = Game{ID: x, Winning: winning, Cards: myCards}
		test[x] = 1
	}

	for id := 0; id < len(test); id++ {
		game := m[id]

		var found int
		for _, card := range game.Cards {
			if isInList(game.Winning, card) {
				found++
			}
		}

		for i := id + 1; i <= found+id; i++ {
			test[i] += test[id]
		}
	}

	for id := 0; id < len(test); id++ {
		sum += test[id]
	}

	return sum
}

func main() {
	b, err := os.ReadFile("input/day04.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	ansOne := partOne(lines)
	fmt.Printf("PartOne: %d\n", ansOne)

	ansTwo := partTwo(lines)
	fmt.Printf("PartTwo: %d\n", ansTwo)
}
