package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Run() {
	day1()
}

func day1() {
	var total int

	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		total += extractNumbersFromStr(reader.Text())
	}

	fmt.Println(total)
}

func extractNumbersFromStr(str string) int {
	var numbers []int
	var sum string

	for _, char := range str {
		if ok := unicode.IsNumber(char); ok {
			num, _ := strconv.Atoi(string(char))
			numbers = append(numbers, num)
		} else {
			sum += string(char)
			for value, num := range digitMap {
				if strings.Contains(sum, value) {
					numbers = append(numbers, num)
					if len(sum)-1 > 0 {
						sum = string(sum[len(sum)-1])
					}
				}
			}
		}
	}

	if len(numbers) >= 1 {
		concatenatedNum := fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1])
		intNum, _ := strconv.Atoi(concatenatedNum)
		return intNum
	} else {
		return 0
	}
}
