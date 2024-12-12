package main

import "fmt"

func Day10() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(10)
	findTrailHeads(lines)

}

var (
	visited = make(map[Point]bool)
)

func findTrailHeads(lines []string) {
	partone := 0
	parttwo := 0
	for i, line := range lines {
		for j, num := range line {
			if num == '0' {
				parttwo += findTrails(lines, num+1, i, j, len(lines), len(lines[0]))
				partone += len(visited)
				visited = make(map[Point]bool)
			}
		}
	}
	fmt.Println(partone)
	fmt.Println(parttwo)
}

func findTrails(lines []string, curr rune, i, j, m, n int) int {
	if curr > '9' {
		visited[Point{i, j}] = true
		return 1
	}
	count := 0
	if i > 0 && rune(lines[i-1][j]) == curr {
		count += findTrails(lines, curr+1, i-1, j, m, n)
	}
	if j > 0 && rune(lines[i][j-1]) == curr {
		count += findTrails(lines, curr+1, i, j-1, m, n)
	}
	if i < m-1 && rune(lines[i+1][j]) == curr {
		count += findTrails(lines, curr+1, i+1, j, m, n)
	}
	if j < n-1 && rune(lines[i][j+1]) == curr {
		count += findTrails(lines, curr+1, i, j+1, m, n)
	}
	return count
}
