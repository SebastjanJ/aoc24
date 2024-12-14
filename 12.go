package main

import (
	"fmt"
	"sort"
)

func Day12() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(12)

	regions := make([][]Point, 0)
	for i, line := range lines {
		for j := range line {
			if visited[Point{i, j}] {
				continue
			}
			regions = append(regions, getRegion(lines, i, j, len(lines), len(lines[0]), lines[i][j]))
		}
	}
	partone12(regions)
	parttwo12(regions)
}

func parttwo12(regions [][]Point) {
	price := 0
	for _, region := range regions {
		area := len(region)
		if area <= 2 {
			price += area * 4
			continue
		}
		sides := topToBottom(region) + leftToRight2(region)
		price += area * sides
	}
	fmt.Println(price)
}

func partone12(regions [][]Point) {
	price := 0
	for _, region := range regions {
		area := len(region)
		perimiter := area * 4
		for _, point := range region {
			for _, otherpoint := range region {
				if point == otherpoint {
					continue
				}
				if point.isNeighbor(otherpoint) {
					perimiter -= 1
				}
			}
		}
		price += area * perimiter
	}
	fmt.Println(price)
}

func getRegion(lines []string, i, j, m, n int, curr byte) []Point {
	point := Point{i, j}
	visited[point] = true

	region := make([]Point, 0)
	region = append(region, point)

	if i > 0 && !visited[Point{i - 1, j}] && lines[i-1][j] == curr {
		region = append(region, getRegion(lines, i-1, j, m, n, curr)...)
	}
	if j > 0 && !visited[Point{i, j - 1}] && lines[i][j-1] == curr {
		region = append(region, getRegion(lines, i, j-1, m, n, curr)...)
	}
	if i < m-1 && !visited[Point{i + 1, j}] && lines[i+1][j] == curr {
		region = append(region, getRegion(lines, i+1, j, m, n, curr)...)
	}
	if j < n-1 && !visited[Point{i, j + 1}] && lines[i][j+1] == curr {
		region = append(region, getRegion(lines, i, j+1, m, n, curr)...)
	}
	return region
}

func countNewSides(a, b map[int]bool) int {
	keys := make([]int, 0, len(a))
	for k := range a {
		if !b[k] {
			keys = append(keys, k)
		}
	}
	if len(keys) == 0 {
		return 0
	}
	count := 1
	sort.Ints(keys)
	for i := 0; i < len(keys)-1; i++ {
		if keys[i]+1 != keys[i+1] {
			count++
		}
	}
	return count
}

func leftToRight2(region []Point) int {
	sides := 0
	sort.Slice(region, func(i, j int) bool {
		return region[i].j < region[j].j
	})
	currj := region[0].j
	curr := make(map[int]bool)
	prev := make(map[int]bool)
	for _, point := range region {
		if point.j != currj {
			sides += countNewSides(curr, prev) + countNewSides(prev, curr)
			prev = curr
			curr = make(map[int]bool)
			currj = point.j
		}
		curr[point.i] = true
	}
	sides += countNewSides(curr, prev) + countNewSides(prev, curr) + countNewSides(curr, make(map[int]bool))
	return sides
}

func topToBottom(region []Point) int {
	sides := 0
	sort.Slice(region, func(i, j int) bool {
		return region[i].i < region[j].i
	})
	curri := region[0].i
	curr := make(map[int]bool)
	prev := make(map[int]bool)
	for _, point := range region {
		if point.i != curri {
			sides += countNewSides(curr, prev) + countNewSides(prev, curr)
			prev = curr
			curr = make(map[int]bool)
			curri = point.i
		}
		curr[point.j] = true
	}
	sides += countNewSides(curr, prev) + countNewSides(prev, curr) + countNewSides(curr, make(map[int]bool))
	return sides
}
