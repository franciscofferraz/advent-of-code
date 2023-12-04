package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/franciscofferraz/advent-of-code/day1"
	"github.com/franciscofferraz/advent-of-code/day2"
	"github.com/franciscofferraz/advent-of-code/day3"
	"github.com/franciscofferraz/advent-of-code/day4"
)

func main() {
	flag.Int("day", 1, "Day")
	flag.Int("part", 1, "Part")
	flag.Parse()

	// day is int
	day := flag.Lookup("day").Value.(flag.Getter).Get().(int)
	if day < 1 || day > 25 {
		day = 1
	}

	part := flag.Lookup("part").Value.(flag.Getter).Get().(int)
	if part < 1 || part > 2 {
		part = 1
	}

	processDay(day, part)
}

func processDay(day int, part int) {
	filePath := fmt.Sprintf("day%d/input.txt", day)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	switch day {
	case 1:
		defer file.Close()
		day1.Run(part, file)
	case 2:
		day2.Run(part, file)
	case 3:
		day3.Run(part, file)
	case 4:
		day4.Run(part, file)
	}
}
