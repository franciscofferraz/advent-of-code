package main

import (
	"flag"

	"github.com/franciscofferraz/advent-of-code/day1"
)

func main() {
	flag.Int("day", 1, "Day")
	flag.Int("part", 1, "Part")
	flag.Parse()

	day := flag.Lookup("day").Value.(flag.Getter).Get().(int)
	part := flag.Lookup("part").Value.(flag.Getter).Get().(int)

	processDay(day, part)
}

func processDay(day int, part int) {
	switch day {
	case 1:
		day1.ProcessPart(part)
	}
}
