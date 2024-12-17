package main

import (
	"fmt"
	"regexp"
)

type Button struct {
	x, y int
}

type Machine struct {
	a, b Button
	x, y int
}

func Day13() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(13)
	machines := getMachines(lines)
	calcTokens(machines)
}

func calcTokens(machines []Machine) {
	tokencount := 0
	for _, m := range machines {
		if m.x%GCD(m.a.x, m.b.x) != 0 || m.y%GCD(m.a.y, m.b.y) != 0 {
			continue
		}
		mintoken := 400
		solved := false

		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				if a*m.a.x+b*m.b.x == m.x && a*m.a.y+b*m.b.y == m.y && 3*a+b < mintoken {
					solved = true
					mintoken = 3*a + b
				}
			}
		}
		if solved {
			tokencount += mintoken
		}
	}
	fmt.Println(tokencount)
}

func getMachines(lines []string) []Machine {
	machines := make([]Machine, 0)
	num := regexp.MustCompile(`\d+`)
	for i := 0; i < len(lines); i += 4 {
		a := num.FindAllString(lines[i], -1)
		b := num.FindAllString(lines[i+1], -1)
		prize := num.FindAllString(lines[i+2], -1)
		machines = append(machines, Machine{
			a: Button{x: strToInt(a[0]), y: strToInt(a[1])},
			b: Button{x: strToInt(b[0]), y: strToInt(b[1])},
			x: strToInt(prize[0]),
			y: strToInt(prize[1]),
		})
	}
	return machines
}
