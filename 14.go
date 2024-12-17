package main

import (
	"fmt"
	"regexp"
	"sort"
)

type Robot struct {
	p Point
	v Point
}

func (r Robot) String() string {
	return r.p.String() + ", " + r.v.String()
}

var (
	// width  = 11
	// height = 7
	width  = 101
	height = 103
)

func Day14() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(14)
	robots := getRobots(lines)
	robotsCopy := make([]Robot, len(robots))
	copy(robotsCopy, robots)
	fmt.Println(countQuadrants(robots))
	fmt.Println(findTree(robotsCopy))
}

func checkForLine(robots []Robot) bool {
	sort.Slice(robots, func(i, j int) bool {
		if robots[i].p.i == robots[j].p.i {
			return robots[i].p.j < robots[j].p.j
		}
		return robots[i].p.i < robots[j].p.i
	})

	icount := 1
	for i := 0; i < len(robots)-1; i++ {
		if robots[i].p.i == robots[i+1].p.i && robots[i].p.j == robots[i+1].p.j-1 {
			icount++
			if icount == 20 {
				return true
			}
		} else {
			icount = 1
		}
	}
	return false
}

func findTree(robots []Robot) int {
	for i := 1; true; i++ {
		simulateRobots(robots)
		if checkForLine(robots) {
			return i
		}
	}
	return 0
}

func countQuadrants(robots []Robot) int {
	for i := 0; i < 100; i++ {
		simulateRobots(robots)
	}
	count := 1
	quadrants := make(map[int]int)
	for _, r := range robots {
		if r.p.i < height/2 && r.p.j < width/2 {
			quadrants[1]++
		} else if r.p.i < height/2 && r.p.j > width/2 {
			quadrants[2]++
		} else if r.p.i > height/2 && r.p.j < width/2 {
			quadrants[3]++
		} else if r.p.i > height/2 && r.p.j > width/2 {
			quadrants[4]++
		}
	}
	for _, v := range quadrants {
		count *= v
	}
	return count
}

func simulateRobots(robots []Robot) {
	for j := 0; j < len(robots); j++ {
		robots[j].p = robots[j].p.add(robots[j].v)
		robots[j].p = Point{i: (robots[j].p.i + height) % height, j: (robots[j].p.j + width) % width}
	}
}

func getRobots(lines []string) []Robot {
	robots := make([]Robot, 0)
	num := regexp.MustCompile(`[-]?\d+`)
	for _, line := range lines {
		args := num.FindAllString(line, -1)
		robots = append(robots, Robot{
			p: Point{j: strToInt(args[0]), i: strToInt(args[1])},
			v: Point{j: strToInt(args[2]), i: strToInt(args[3])},
		})
	}
	return robots
}
