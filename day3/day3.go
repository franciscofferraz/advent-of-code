package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func Run(part int, file *os.File) {
	var answer int
	var data []string
	startTime := time.Now()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	switch part {
	case 1:
		answer = part1(data)
	case 2:
		answer = part2(data)
	}

	usedTime := time.Since(startTime)
	fmt.Printf("Day 3 Part %d Answer: %d\n", part, answer)
	fmt.Printf("Execution time: %s\n", usedTime)
}

var symbolsMap = map[byte]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
	'@': true,
	'&': true,
	'$': true,
	'#': true,
	'=': true,
	'%': true,
}

var r = regexp.MustCompile(`\d+`)

type GearPart struct {
	startIdx int
	endIdx   int
	value    int
}

func part1(data []string) int {
	var sum int

	for lineIdx, line := range data {
		numsIdx := r.FindAllStringIndex(line, -1)
		for _, numIdx := range numsIdx {
			var hasTop, hasBottom, hasLeft, hasRight bool
			num, _ := strconv.Atoi(line[numIdx[0]:numIdx[1]])

			hasRight = checkRight(numIdx, line)
			hasLeft = checkLeft(numIdx, line)

			if lineIdx != 0 {
				hasTop = isAdjacentToSymbol(numIdx, data[lineIdx-1])
			}

			if lineIdx != len(data)-1 {
				hasBottom = isAdjacentToSymbol(numIdx, data[lineIdx+1])
			}

			if hasTop || hasBottom || hasLeft || hasRight {
				sum += num
			}
		}
	}

	return sum
}

func part2(data []string) int {
	var sum int

	reGear := regexp.MustCompile(`\*`)
	reDigits := regexp.MustCompile(`\d+`)

	for lineIdx, line := range data {
		gears := reGear.FindAllStringIndex(line, -1)

		for _, gear := range gears {
			gearIdx := gear[0]

			var gearParts []GearPart
			addGearParts := func(parts [][]int, lineOffset int) {
				for _, part := range parts {
					gearPart := GearPart{
						startIdx: part[0],
						endIdx:   part[1],
					}
					value, _ := strconv.Atoi(data[lineIdx+lineOffset][part[0]:part[1]])
					gearPart.value = value

					if part[1] >= gearIdx && part[0]-1 <= gearIdx {
						gearParts = append(gearParts, gearPart)
					}
				}
			}

			addGearParts(reDigits.FindAllStringIndex(line, -1), 0)

			if lineIdx != 0 {
				addGearParts(reDigits.FindAllStringIndex(data[lineIdx-1], -1), -1)
			}

			if lineIdx != len(data)-1 {
				addGearParts(reDigits.FindAllStringIndex(data[lineIdx+1], -1), 1)
			}

			if len(gearParts) != 2 {
				continue
			}

			sum += gearParts[0].value * gearParts[1].value
		}
	}

	return sum
}

func checkRight(numIdx []int, line string) bool {
	return numIdx[1] != len(line) && symbolsMap[line[numIdx[1]]]
}

func checkLeft(numIdx []int, line string) bool {
	return numIdx[0] != 0 && symbolsMap[line[numIdx[0]-1]]
}

func isAdjacentToSymbol(numIdx []int, line string) bool {
	if symbolsMap[line[numIdx[0]]] || symbolsMap[line[numIdx[1]-1]] {
		return true
	}

	idxRange := [2]int{}

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
		if symbolsMap[byte(char)] {
			return true
		}
	}

	return false
}
