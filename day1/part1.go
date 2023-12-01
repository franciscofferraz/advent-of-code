package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
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
	var numbers []int

	for _, char := range line {
		if ok := unicode.IsNumber(char); ok {
			num, _ := strconv.Atoi(string(char))
			numbers = append(numbers, num)
		}
	}

	if len(numbers) == 0 {
		return 0
	}

	concatenatedNum := fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1])
	intNum, _ := strconv.Atoi(concatenatedNum)

	return intNum
}
