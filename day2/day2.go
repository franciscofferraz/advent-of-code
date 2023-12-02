package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	red   = 12
	green = 13
	blue  = 14
)

func Run() {
	main()
}

func main() {
	startTime := time.Now()
	var answer int

	f, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		answer += checkGame(scanner.Text())
	}

	fmt.Println("Answer:", answer)

	used := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", used)
}

func checkGame(str string) int {
	gameStr := strings.Split(str, ": ")

	gameId, err := strconv.Atoi(gameStr[0][5:])
	if err != nil {
		gameId = 0
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
					if number > red {
						return 0
					}
				case "green":
					number, _ := strconv.Atoi(set_element[idx-1])
					if number > green {
						return 0
					}
				case "blue":
					number, _ := strconv.Atoi(set_element[idx-1])
					if number > blue {
						return 0
					}
				}
			}
		}
	}

	return gameId
}
