package main

import "fmt"

func Day8() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(8)
	antenas := getAtenas(lines)
	m := len(lines)
	n := len(lines[0])
	partonetwo8(antenas, m, n)
}

func partonetwo8(antenas map[rune][]Point, m, n int) {
	antinodes1 := make(map[Point]bool)
	antinodes2 := make(map[Point]bool)
	for _, points := range antenas {
		for i, point := range points {
			if len(points) < 2 {
				continue
			}
			for j, otherpoint := range points {
				if i == j {
					continue
				}
				// partone
				antinode := Point{point.i - 2*(point.i-otherpoint.i), point.j - 2*(point.j-otherpoint.j)}
				if !antinodes1[antinode] && antinode.i >= 0 && antinode.i < m && antinode.j >= 0 && antinode.j < n {
					antinodes1[antinode] = true
				}
				// parttwo
				distance := Point{otherpoint.i - point.i, otherpoint.j - point.j}
				curr := otherpoint
				for ; curr.i >= 0 && curr.i < m && curr.j >= 0 && curr.j < n; curr = curr.add(distance) {
					antinodes2[curr] = true
				}
			}
		}
	}
	fmt.Println(len(antinodes1))
	fmt.Println(len(antinodes2))
}

func getAtenas(lines []string) map[rune][]Point {
	antenas := make(map[rune][]Point)
	for i, line := range lines {
		for j, char := range line {
			if char == '.' {
				continue
			}
			if _, exists := antenas[char]; exists {
				antenas[char] = append(antenas[char], Point{i, j})
			} else {
				antenas[char] = []Point{{i, j}}
			}
		}
	}
	return antenas
}
