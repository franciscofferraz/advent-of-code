package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type digit struct {
	str string
	num int
}

var digits = []digit{
	{"one", 1},
	{"two", 2},
	{"three", 3},
	{"four", 4},
	{"five", 5},
	{"six", 6},
	{"seven", 7},
	{"eight", 8},
	{"nine", 9},
}

func Run() {
	main()
}

func main() {
	startTime := time.Now()
	var answer int

	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		answer += calibrate(scanner.Text())
	}

	fmt.Println(answer)

	used := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", used)
}

func calibrate(str string) int {
	var numbers []int
	var tempBuffer string

	for _, char := range str {
		if ok := unicode.IsNumber(char); ok {
			num, _ := strconv.Atoi(string(char))
			numbers = append(numbers, num)
		} else {
			tempBuffer += string(char)
			for _, d := range digits {
				if strings.HasSuffix(tempBuffer, d.str) {
					numbers = append(numbers, d.num)
					if len(tempBuffer)-1 > 0 {
						tempBuffer = string(tempBuffer[len(tempBuffer)-1])
					} else {
						tempBuffer = ""
					}
				}
			}
		}
	}

	if len(numbers) >= 1 {
		concatenatedNum := numbers[0]*10 + numbers[len(numbers)-1]
		return concatenatedNum
	} else {
		return 0
	}
}
