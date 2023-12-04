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
		answer = part2(data)
	}

	fmt.Printf("Day 4 Part %d Answer: %d\n", part, answer)
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
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

func part2(data []string) int {
	cardWinsNoOfCards := make(map[int]int)
	cardWinningNumbers := make(map[int]map[string]bool)
	cardQueue := make([]int, 0)

	for i, card := range data {
		currentCardWins := 0
		winningNumbers := make(map[string]bool)

		parts := strings.Split(card, ":")
		winningNums := strings.Fields(strings.Split(parts[1], " | ")[0])
		myNumbers := strings.Fields(strings.Split(parts[1], " | ")[1])

		for _, number := range winningNums {
			winningNumbers[number] = true
		}

		for _, number := range myNumbers {
			if winningNumbers[number] {
				currentCardWins++
			}
		}

		cardWinsNoOfCards[i+1] = currentCardWins
		cardWinningNumbers[i+1] = winningNumbers
		cardQueue = append(cardQueue, i+1)
	}

	cardsProcessed := 0
	queueLength := len(cardQueue)

	for cardsProcessed < queueLength {
		start := cardQueue[cardsProcessed]
		end := start + cardWinsNoOfCards[start]

		cardQueue = append(cardQueue, generateRange(start+1, end)...)
		cardsProcessed++
		queueLength += cardWinsNoOfCards[start]
	}

	return cardsProcessed
}

func generateRange(start, end int) []int {
	numRange := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		numRange = append(numRange, i)
	}

	return numRange
}
