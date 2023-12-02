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
	var buf strings.Builder

	goto inner

inner:
	for i := 0; i < len(str); i++ {
		if unicode.IsNumber(rune(str[i])) {
			num, _ := strconv.Atoi(string(str[i]))
			numbers = append(numbers, num)
			goto outer
		} else {
			buf.WriteString(string(str[i]))

			if len(buf.String()) >= minLen {
				for _, d := range digits {
					if strings.HasSuffix(buf.String(), d.str) {
						numbers = append(numbers, d.num)
						buf.Reset()
						goto outer
					}
				}
			}
		}
	}
	buf.Reset()

outer:
	for i := len(str) - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(str[i])) {
			num, _ := strconv.Atoi(string(str[i]))
			numbers = append(numbers, num)
			goto end
		} else {
			buf.WriteString(string(str[i]))
			if len(buf.String()) >= minLen {
				for _, d := range reversedDigits {
					if strings.HasSuffix(buf.String(), d.str) {
						numbers = append(numbers, d.num)
						buf.Reset()
						goto end
					}
				}
			}
		}
	}

end:
	if len(numbers) >= 1 {
		concatenatedNum := numbers[0]*10 + numbers[len(numbers)-1]
		println(concatenatedNum)
		return concatenatedNum
	} else {
		return 0
	}
}
