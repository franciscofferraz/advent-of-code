package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	}

	usedTime := time.Since(startTime)
	fmt.Printf("Day 4 Part %d Answer: %d\n", part, answer)
	fmt.Printf("Execution time: %s\n", usedTime)
}

func part1(lines []string) int {
	var points int

	for _, line := range lines {
		loopPoints := 0
		splittedLine := strings.Split(line, "|")

		winningNumbers := strings.Fields(strings.TrimSpace(splittedLine[0]))
		myNumbers := strings.Fields(strings.TrimSpace(splittedLine[1]))

		for _, winningNum := range winningNumbers {
			for _, myNum := range myNumbers {
				if winningNum == myNum {
					if loopPoints == 0 {
						loopPoints++
					} else {
						loopPoints = loopPoints * 2
					}
				}
			}
		}
		points += loopPoints
	}

	return points
}
