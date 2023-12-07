package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Red   int
	Blue  int
	Green int
}

var (
	redLimit   = 12
	greenLimit = 13
	blueLimit  = 14
	gameIDRe   = regexp.MustCompile(`Game (\d{1,3})`)
	setRe      = regexp.MustCompile(`\s?(\d{1,3})\s?(\w{3,5})`)
)

func parseColors(game *Game, subsets []string) {
	for _, subset := range subsets {
		sets := strings.Split(subset, ", ")

		for _, set := range sets {
			groups := setRe.FindStringSubmatch(set)

			num, err := strconv.Atoi(groups[1])
			if err != nil {
				log.Printf("ERROR: %s : %s", groups[1], err)
			}

			switch groups[2] {
			case "red":
				if game.Red < num {
					game.Red = num
				}
			case "blue":
				if game.Blue < num {
					game.Blue = num
				}
			case "green":
				if game.Green < num {
					game.Green = num
				}
			}
		}
	}
}

func isPossible(game *Game) bool {
	return game.Red <= redLimit && game.Green <= greenLimit && game.Blue <= blueLimit
}

func partOne(lines []string) int {
	var idSum int

	for _, line := range lines {
		groups := gameIDRe.FindStringSubmatch(line)

		id, err := strconv.Atoi(groups[1])
		if err != nil {
			log.Printf("ERROR string to int %c : %s", line[5], err)
		}

		game := new(Game)

		subsets := strings.Split(strings.Split(line, ":")[1], ";")
		parseColors(game, subsets)

		if isPossible(game) {
			idSum += id
		}
	}

	return idSum
}

func partTwo(lines []string) int {
	var idSum int

	for _, line := range lines {
		game := new(Game)

		subsets := strings.Split(strings.Split(line, ":")[1], ";")
		parseColors(game, subsets)

		idSum += game.Red * game.Green * game.Blue
	}

	return idSum
}

func main() {
	b, err := os.ReadFile("input/day02.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	ansOne := partOne(lines)
	fmt.Printf("PartOne: %d\n", ansOne)

	ansTwo := partTwo(lines)
	fmt.Printf("PartTwo: %d\n", ansTwo)
}
