package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func Run(part int, file *os.File) {
	var answer int
	startTime := time.Now()

	scanner := bufio.NewScanner(file)
	switch part {
	case 1:
		for scanner.Scan() {
			answer += part1(scanner.Text())
		}

	case 2:
		for scanner.Scan() {
			answer += part2(scanner.Text())
		}
	}

	fmt.Printf("Day 2 Part %d Answer: %d\n", part, answer)
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
}

func part1(str string) int {
	gameStr := strings.Split(str, ": ")

	gameId, err := strconv.Atoi(strings.TrimPrefix(gameStr[0], "Game "))
	if err != nil {
		return 0
	}

	gameSets := strings.Split(gameStr[1], "; ")
	for _, set := range gameSets {
		set_elements := strings.Split(set, ", ")
		for _, current := range set_elements {
			set_element := strings.Split(current, " ")
			for idx, value := range set_element {
				switch value {
				case "red":
					number, _ := strconv.Atoi(set_element[idx-1])
					if number > maxRed {
						return 0
					}
				case "green":
					number, _ := strconv.Atoi(set_element[idx-1])
					if number > maxGreen {
						return 0
					}
				case "blue":
					number, _ := strconv.Atoi(set_element[idx-1])
					if number > maxBlue {
						return 0
					}
				}
			}
		}
	}

	return gameId
}

func part2(str string) int {
	var red, green, blue int
	var power int

	gameStr := strings.Split(str, ": ")[1]

	gameSets := strings.Split(gameStr, "; ")
	for _, set := range gameSets {
		set_elements := strings.Split(set, ", ")
		for _, current := range set_elements {
			set_element := strings.Split(current, " ")
			for idx := 0; idx < len(set_element); idx++ {
				if idx > 0 {
					switch set_element[idx] {
					case "red":
						number, _ := strconv.Atoi(set_element[idx-1])
						if number > red {
							red = number
						}
					case "green":
						number, _ := strconv.Atoi(set_element[idx-1])
						if number > green {
							green = number
						}
					case "blue":
						number, _ := strconv.Atoi(set_element[idx-1])
						if number > blue {
							blue = number
						}
					}
				}
			}
		}
	}
	power = red * green * blue

	return power
}
