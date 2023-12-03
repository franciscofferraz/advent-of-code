package day3

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	main()
}

const symbols string = "+-*/@&$#=%"

var r = regexp.MustCompile(`\d+`)

func main() {
	var data []string

	f, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	println("Part 1:", part1(data))
}

func part1(data []string) int {
	var sum int

	for lineIdx, line := range data {
		numsIdx := r.FindAllStringIndex(line, -1)
		for _, numIdx := range numsIdx {
			num, _ := strconv.Atoi(line[numIdx[0]:numIdx[1]])

			var hasTop, hasBottom, hasLeft, hasRight bool
			hasRight = checkRight(numIdx, line)
			hasLeft = checkLeft(numIdx, line)
			if lineIdx != 0 {
				hasTop = isValid(numIdx, data[lineIdx-1])
			}
			if lineIdx != len(data)-1 {
				hasBottom = isValid(numIdx, data[lineIdx+1])
			}

			if hasLeft || hasRight || hasTop || hasBottom {
				sum += num
			}
		}
	}

	return sum
}

func checkRight(numIdx []int, line string) bool {
	if numIdx[1] != len(line) {
		if strings.Contains(symbols, string(line[numIdx[1]])) {
			return true
		}
	}

	return false
}

func checkLeft(numIdx []int, line string) bool {
	if numIdx[0] != 0 {
		if strings.Contains(symbols, string(line[numIdx[0]-1])) {
			return true
		}
	}

	return false
}

func isValid(numIdx []int, line string) bool {
	if strings.Contains(symbols, string(line[numIdx[0]])) || strings.Contains(symbols, string(line[numIdx[1]-1])) {
		return true
	}

	idxRange := make([]int, 2)

	if numIdx[0] == 0 {
		idxRange[0] = 0
	} else {
		idxRange[0] = numIdx[0] - 1
	}

	if numIdx[1] == len(line) {
		idxRange[1] = len(line)
	} else {
		idxRange[1] = numIdx[1] + 1
	}

	for _, char := range line[idxRange[0]:idxRange[1]] {
		if strings.Contains(symbols, string(char)) {
			return true
		}
	}

	return false
}
