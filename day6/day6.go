package day6

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var r = regexp.MustCompile(`\d+`)

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

	fmt.Printf("Day 6 Part %d Answer: %d\n", part, answer)
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
}

func part1(lines []string) int {
	var result = 1

	durations := r.FindAllString(lines[0], -1)
	distances := r.FindAllString(lines[1], -1)

	var durationInts []int
	var distanceInts []int

	for _, durStr := range durations {
		duration, _ := strconv.Atoi(durStr)
		durationInts = append(durationInts, duration)
	}

	for _, distStr := range distances {
		distance, _ := strconv.Atoi(distStr)
		distanceInts = append(distanceInts, distance)
	}

	races := make([]int, len(durationInts))

	for idx, duration := range durationInts {
		for i := 0; i <= duration; i++ {
			speed := i * 1
			race := speed * (duration - i)

			if race > distanceInts[idx] {
				races[idx]++
			}
		}
	}

	for _, race := range races {
		result = result * race
	}

	return result
}

func part2(data []string) int {
	return 0
}
