package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func Run() {
	main()
}

func main() {
	var answer1, answer2 int

	f, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		answer1 += part1(scanner.Text())
		answer2 += part2(scanner.Text())
	}

	fmt.Println("Part1 Answer:", answer1)
	fmt.Println("Part2 Answer:", answer2)
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
