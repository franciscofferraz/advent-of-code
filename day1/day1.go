package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	minLen = 3
	maxLen = 5
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

var reversedDigits = []digit{
	{"eno", 1},
	{"owt", 2},
	{"eerht", 3},
	{"ruof", 4},
	{"evif", 5},
	{"xis", 6},
	{"neves", 7},
	{"thgie", 8},
	{"enin", 9},
}

func Run(part int, file *os.File) {
	var answer int
	startTime := time.Now()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		answer += calibrate(scanner.Text())
	}

	used := time.Since(startTime)
	if part == 1 {
		fmt.Print("No part 1 for this day, sorry!\n")
	}
	fmt.Printf("Day 1 Part 2 Answer: %d\n", answer)
	fmt.Printf("Execution time: %s\n", used)
}

func calibrate(str string) int {
	var numbers []int

	numbers = append(numbers, getFirstNumber(str))
	numbers = append(numbers, getLastNumber(str))

	if len(numbers) >= 1 {
		concatenatedNum := numbers[0]*10 + numbers[len(numbers)-1]
		return concatenatedNum
	} else {
		return 0
	}
}

func getFirstNumber(str string) int {
	var buf strings.Builder

	for i := 0; i < len(str); i++ {
		if unicode.IsNumber(rune(str[i])) {
			num, _ := strconv.Atoi(string(str[i]))
			return num
		} else {
			buf.WriteString(string(str[i]))

			if len(buf.String()) >= minLen {
				for _, d := range digits {
					if strings.HasSuffix(buf.String(), d.str) {
						return d.num
					}
				}
			}
		}
	}

	return 0
}

func getLastNumber(str string) int {
	var buf strings.Builder

	for i := len(str) - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(str[i])) {
			num, _ := strconv.Atoi(string(str[i]))
			return num
		} else {
			buf.WriteString(string(str[i]))
			if len(buf.String()) >= minLen {
				for _, d := range reversedDigits {
					if strings.HasSuffix(buf.String(), d.str) {
						return d.num
					}
				}
			}
		}
	}

	return 0
}
