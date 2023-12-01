package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1() {
	var total int

	f, err := os.Open("day1/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := make(map[int]int)

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		value := extractNumbersFromStr(line)

		m[i] = value
	}

	for _, number := range m {
		total += number
	}

	println(total)
}

func extractNumbersFromStr(line string) int {
	var numbers string

	for _, char := range line {
		if _, err := strconv.Atoi(string(char)); err == nil {
			numbers += string(char)
		}
	}

	if numbers == "" {
		return 0
	}

	firstNum, _ := strconv.Atoi(string(numbers[0]))
	lastNum, _ := strconv.Atoi(string(numbers[len(numbers)-1]))

	concatenatedNum := fmt.Sprintf("%d%d", firstNum, lastNum)
	intNum, _ := strconv.Atoi(concatenatedNum)

	return intNum
}
