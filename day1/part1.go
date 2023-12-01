package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1() {
	var numbers []int
	var total int

	f, err := os.Open("day1/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		number := extractNumbersFromStr(line)
		numbers = append(numbers, number)
	}

	for _, number := range numbers {
		total += number
	}

	fmt.Println(total)
}

func extractNumbersFromStr(line string) int {
	var numbers string

	for _, char := range line {
		if num, err := strconv.Atoi(string(char)); err == nil {
			numbers += strconv.Itoa(num)
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
