package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Run() {
	main()
}

func main() {
	var answer1, answer2 int

	f, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		answer1 += part1(scanner.Text())
		answer2 += part2(scanner.Text())
	}

	fmt.Println("Part1 Answer:", answer1)
	fmt.Println("Part2 Answer:", answer2)
}

func part1(str string) int {
	return 0
}

func part2(str string) int {
	return 0
}
