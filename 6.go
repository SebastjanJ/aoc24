package main

import (
	"fmt"
)

type Point struct {
	i, j int
}

type Guard struct {
	cor   Point
	dir   rune
	steps int
}

func (g Guard) String() string {
	return fmt.Sprintf("Guard at (%d,%d) facing %c, steps %d", g.cor.i, g.cor.j, g.dir, g.steps)
}

func (g *Guard) move(lines []string, obstacle *Point) bool {
	switch g.dir {
	case '^':
		nextPos := Point{g.cor.i - 1, g.cor.j}
		if lines[nextPos.i][nextPos.j] == '#' || (obstacle != nil && nextPos == *obstacle) {
			g.dir = '>'
			return false
		}
		g.cor.i--
	case 'v':
		nextPos := Point{g.cor.i + 1, g.cor.j}
		if lines[nextPos.i][nextPos.j] == '#' || (obstacle != nil && nextPos == *obstacle) {
			g.dir = '<'
			return false
		}
		g.cor.i++
	case '>':
		nextPos := Point{g.cor.i, g.cor.j + 1}
		if lines[nextPos.i][nextPos.j] == '#' || (obstacle != nil && nextPos == *obstacle) {
			g.dir = 'v'
			return false
		}
		g.cor.j++
	case '<':
		nextPos := Point{g.cor.i, g.cor.j - 1}
		if lines[nextPos.i][nextPos.j] == '#' || (obstacle != nil && nextPos == *obstacle) {
			g.dir = '^'
			return false
		}
		g.cor.j--
	}
	return true
}

func Day6() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(6)

	guard := findGuard(lines)
	visits := partone6(lines, guard)
	parttwo6(lines, guard, visits)

}

func partone6(lines []string, g Guard) map[Point]bool {
	visits := make(map[Point]bool)
	for g.cor.i > 0 && g.cor.i+1 < len(lines) && g.cor.j > 0 && g.cor.j+1 < len(lines[0]) {
		visits[Point{g.cor.i, g.cor.j}] = true
		if !g.move(lines, nil) {
			continue
		}

		if !visits[Point{g.cor.i, g.cor.j}] {
			g.steps++
		}

	}
	visits[Point{g.cor.i, g.cor.j}] = true
	fmt.Println(g.steps)
	return visits
}

func parttwo6(lines []string, g Guard, visits map[Point]bool) {
	count := 0
	start := g
visitLoop:
	for point := range visits {
		if point == start.cor {
			continue
		}
		currvisits := make(map[Guard]bool)
		g = start
		for g.cor.i > 0 && g.cor.i+1 < len(lines) && g.cor.j > 0 && g.cor.j+1 < len(lines[0]) {
			if currvisits[g] {
				count++
				continue visitLoop
			}
			currvisits[g] = true
			g.move(lines, &point)
		}
	}
	fmt.Println(count)
}

func findGuard(lines []string) Guard {
	for i, line := range lines {
		for j, char := range line {
			if char != '.' && char != '#' {
				return Guard{Point{i, j}, char, 1}
			}
		}
	}
	return Guard{Point{0, 0}, 0, 0}
}
