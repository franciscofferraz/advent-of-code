package main

import (
	"flag"

	"github.com/franciscofferraz/advent-of-code/day1"
)

func main() {
	flag.Int("day", 1, "Day")
	flag.Parse()

	day := flag.Lookup("day").Value.(flag.Getter).Get().(int)

	processDay(day)
}

func processDay(day int) {
	switch day {
	case 1:
		day1.Run()
	}
}
