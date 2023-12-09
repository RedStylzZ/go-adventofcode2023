package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Map struct {
	DestStart int
	SrcStart  int
	Len       int
}

func toIntArr(input []string) []int {
	arr := make([]int, len(input))

	for i := range input {
		x, _ := strconv.Atoi(input[i])
		arr[i] = x
	}
	return arr
}

func initMaps(lines []string) []Map {
	m := make([]Map, len(lines))

	for i := range lines {
		info := toIntArr(strings.Split(lines[i], " "))
		m[i] = Map{
			DestStart: info[0],
			SrcStart:  info[1],
			Len:       info[2] - 1,
		}
	}

	return m
}

func isNumBetween(start, end, num int) bool {
	return num >= start && num <= end
}

func getInfo(maps []Map, d int) int {
	for _, m := range maps {
		if !isNumBetween(m.SrcStart, m.SrcStart+m.Len, d) {
			continue
		}

		offset := d - m.SrcStart
		return m.DestStart + offset
	}

	return d
}

func partOne(content string) int {
	var sum int

	data := strings.Split(content, "\n\n")
	seeds := toIntArr(strings.Split(data[0], " ")[1:])

	instr := make(map[int][]Map)
	for i := 1; i < len(data); i++ {
		info := strings.Split(data[i], "\n")
		instr[i-1] = initMaps(info[1:])
	}

	sum = -1
	for _, seed := range seeds {
		soil := getInfo(instr[0], seed)
		fert := getInfo(instr[1], soil)
		watr := getInfo(instr[2], fert)
		ligh := getInfo(instr[3], watr)
		temp := getInfo(instr[4], ligh)
		humi := getInfo(instr[5], temp)
		loca := getInfo(instr[6], humi)

		if loca < sum || sum == -1 {
			sum = loca
		}
	}

	return sum
}

func partTwo(content string) int {
	var sum int

	data := strings.Split(content, "\n\n")
	initSeeds := toIntArr(strings.Split(data[0], " ")[1:])

	instr := make([][]Map, len(data))
	for i := 1; i < len(data); i++ {
		info := strings.Split(data[i], "\n")
		instr[i-1] = initMaps(info[1:])
	}

	jobs := make(chan []int, len(initSeeds)/2)
	for i := 0; i < len(initSeeds); i += 2 {
		seed := initSeeds[i : i+2]
		jobs <- seed
	}
	close(jobs)

	workers := 10

	if workers > len(jobs) {
		workers = len(jobs)
	}

	result := make(chan int, workers)

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			lowest := -1
			for seed := range jobs {
				start := seed[0]
				end := start + seed[1]
				fmt.Printf("Start: %d | End: %d\n", start, end)
				for x := start; x < end; x++ {
					soil := getInfo(instr[0], x)
					fert := getInfo(instr[1], soil)
					watr := getInfo(instr[2], fert)
					ligh := getInfo(instr[3], watr)
					temp := getInfo(instr[4], ligh)
					humi := getInfo(instr[5], temp)
					loca := getInfo(instr[6], humi)

					if loca < lowest || lowest == -1 {
						lowest = loca
					}
				}
				fmt.Printf("Finished: %d\n", start)
			}
			if lowest > 0 {
				result <- lowest
			}

		}(&wg)
	}
	wg.Wait()
	close(result)

	sum = -1
	for i := range result {
		if i < sum || sum == -1 {
			sum = i
		}
	}

	return sum
}

func main() {
	b, err := os.ReadFile("input/day05.txt")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	ansOne := partOne(string(b))
	fmt.Printf("PartOne: %d (%s)\n", ansOne, time.Since(start))

	start = time.Now()
	ansTwo := partTwo(string(b))
	fmt.Printf("PartTwo: %d (%s)\n", ansTwo, time.Since(start))
}
