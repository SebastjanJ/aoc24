package main

import (
	"fmt"
	"math"
	"strings"
)

type Historian struct {
	pos   Point
	steps int
}

func testForExample() {
	lines, _ := GetTest[string]()
	bytes := getBytes(lines, 12)
	fmt.Println(bfs(bytes, Point{0, 0}, Point{6, 6}))
	bruteForce(lines, bytes, 12, Point{6, 6})
}

func testForInput() {
	lines := getInputStrings(18)
	bytes := getBytes(lines, 1024)
	fmt.Println(bfs(bytes, Point{0, 0}, Point{70, 70}))
	bruteForce(lines, bytes, 1024, Point{70, 70})
}

func Day18() {
	// testForExample()
	testForInput()
}

func bruteForce(lines []string, bytes map[Point]bool, n int, end Point) {
	for i := n; i < len(lines); i++ {
		args := strings.Split(lines[i], ",")
		curr := Point{strToInt(args[1]), strToInt(args[0])}
		bytes[curr] = true
		minsteps := bfs(bytes, Point{0, 0}, end)
		if minsteps == math.MaxInt {
			fmt.Println(curr.j, ",", curr.i)
			break
		}
	}
}

func bfs(bytes map[Point]bool, start, end Point) int {
	queue := []Historian{{start, 0}}
	visited := make(map[Point]bool)
	visited[start] = true
	minsteps := math.MaxInt

	directions := []Point{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for len(queue) > 0 {
		h := queue[0]
		queue = queue[1:]

		if h.pos == end {
			minsteps = min(minsteps, h.steps)
		}

		for _, dir := range directions {
			neighbor := Point{h.pos.i + dir.i, h.pos.j + dir.j}
			if neighbor.i >= 0 && neighbor.i <= end.i && neighbor.j >= 0 && neighbor.j <= end.j && !bytes[neighbor] && !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, Historian{neighbor, h.steps + 1})
			}
		}
	}
	return minsteps
}

func getBytes(lines []string, n int) map[Point]bool {
	bytes := make(map[Point]bool)
	for i := 0; i < n; i++ {
		args := strings.Split(lines[i], ",")
		bytes[Point{strToInt(args[1]), strToInt(args[0])}] = true
	}
	return bytes
}
