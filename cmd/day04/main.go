package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
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

func isInList(winning []int, num int) bool {
	for _, i := range winning {
		if i == num {
			return true
		}
	}

	return false
}

func printNums(arr []string) string {

	buf := bytes.Buffer{}
	buf.WriteString("[")
	for i, s := range arr {
		last := ", "
		if i == len(arr)-1 {
			last = ""
		}
		buf.WriteString(fmt.Sprintf("'%s'%s", s, last))
	}
	buf.WriteString("]")
	return buf.String()
}

func printNumsInt(arr []int) string {

	buf := bytes.Buffer{}
	buf.WriteString("[")
	for i, s := range arr {
		last := ", "
		if i == len(arr)-1 {
			last = ""
		}
		buf.WriteString(fmt.Sprintf("'%d'%s", s, last))
	}
	buf.WriteString("]")
	return buf.String()
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
	for x, line := range lines {
		winning, myCards := parseLine(line)
		m[x] = Game{ID: x, Winning: winning, Cards: myCards}
	}

	arr := make([][]Game, len(m))
	for k, g := range m {
		fmt.Printf("Game %d: %+v | %+v\n", k, printNumsInt(g.Winning), printNumsInt(g.Cards))
		arr[k] = []Game{g}
	}

	return sum
}

func main() {
	b, err := os.ReadFile("input/day04_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	ansOne := partOne(lines)
	fmt.Printf("PartOne: %d\n", ansOne)

	ansTwo := partTwo(lines)
	fmt.Printf("PartTwo: %d\n", ansTwo)
}
