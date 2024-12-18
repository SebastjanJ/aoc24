package main

import (
	"fmt"
	"slices"
	"strings"
)

func Day15() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(15)
	warehouse, moves, robot := parseData(lines)

	robot2 := Point{robot.i, robot.j * 2}
	warehouse2 := getPart2Warehouse(warehouse)

	moveRobot(warehouse, moves, robot)
	fmt.Println(gps(warehouse))
	moveRobot(warehouse2, moves, robot2)
	fmt.Println(gps(warehouse2))
}

func gps(warehouse [][]rune) int {
	sum := 0
	for i := 1; i < len(warehouse)-1; i++ {
		for j := 1; j < len(warehouse[i])-1; j++ {
			if rune(warehouse[i][j]) == 'O' || rune(warehouse[i][j]) == '[' {
				sum += i*100 + j
			}
		}
	}
	return sum
}

func moveRobot(warehouse [][]rune, moves string, robot Point) {
	directions := map[rune]Point{
		'<': {0, -1},
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
	}

	for _, move := range moves {
		dir := directions[move]
		moveInDirection(warehouse, &robot, dir)
	}
}

func moveInDirection(warehouse [][]rune, robot *Point, dir Point) {
	ni, nj := robot.i+dir.i, robot.j+dir.j

	if warehouse[ni][nj] == '#' {
		return
	}
	if warehouse[ni][nj] == '.' {
		robot.i, robot.j = ni, nj
		return
	}
	if warehouse[ni][nj] == 'O' {
		start := Point{ni, nj}
		end := start
		for warehouse[end.i+dir.i][end.j+dir.j] == 'O' {
			end.i += dir.i
			end.j += dir.j
		}
		if warehouse[end.i+dir.i][end.j+dir.j] == '.' {
			warehouse[end.i+dir.i][end.j+dir.j] = 'O'
			warehouse[start.i][start.j] = '.'
			robot.i, robot.j = ni, nj
		}
	}
	// part2
	if warehouse[ni][nj] == '[' || warehouse[ni][nj] == ']' {
		if dir.i == 0 {
			// LEVO DESNO
			start := Point{ni, nj}
			end := start.add(dir)
			for warehouse[end.i][end.j+dir.j] == '[' || warehouse[end.i][end.j+dir.j] == ']' {
				end.j += 2 * dir.j
			}
			if warehouse[end.i][end.j+dir.j] == '.' {
				for j := end.j; j != start.j-2*dir.j; j -= dir.j {
					warehouse[ni][j+dir.j] = warehouse[ni][j]
				}
				robot.i, robot.j = ni, nj
			}
		} else {
			// GOR DOL
			stack := []Point{{ni, nj}}
			allstack := []Point{}
			if warehouse[ni][nj] == '[' {
				stack = append(stack, Point{ni, nj + 1})
			} else {
				stack = append(stack, Point{ni, nj - 1})
			}
			for len(stack) > 0 {
				curr := stack[0]
				ni, nj := curr.i+dir.i, curr.j
				if warehouse[ni][nj] == '#' {
					return
				}
				if warehouse[ni][nj] == '[' {
					stack = append(stack, Point{ni, nj})
					stack = append(stack, Point{ni, nj + 1})
				} else if warehouse[ni][nj] == ']' {
					stack = append(stack, Point{ni, nj})
					stack = append(stack, Point{ni, nj - 1})
				}

				if !slices.Contains(allstack, curr) {
					allstack = append(allstack, curr)
				}
				stack = stack[1:]
			}
			for i := len(allstack) - 1; i >= 0; i-- {
				warehouse[allstack[i].i+dir.i][allstack[i].j] = warehouse[allstack[i].i][allstack[i].j]
				warehouse[allstack[i].i][allstack[i].j] = '.'
			}
			robot.i, robot.j = ni, nj
		}
	}
}

func getPart2Warehouse(warehouse [][]rune) [][]rune {
	newWarehouse := make([][]rune, len(warehouse))

	for i, row := range warehouse {
		newWarehouse[i] = make([]rune, len(warehouse[0])*2)
		for j, tile := range row {
			switch tile {
			case '#':
				newWarehouse[i][2*j] = '#'
				newWarehouse[i][2*j+1] = '#'
			case 'O':
				newWarehouse[i][2*j] = '['
				newWarehouse[i][2*j+1] = ']'
			case '.':
				newWarehouse[i][2*j] = '.'
				newWarehouse[i][2*j+1] = '.'
			}
		}
	}
	return newWarehouse
}

func printWarehouse(warehouse [][]rune, robot Point) {
	for i, row := range warehouse {
		for j, tile := range row {
			if i == robot.i && j == robot.j {
				fmt.Print("@")
			} else {
				fmt.Print(string(tile))
			}
		}
		fmt.Println()
	}
}

func parseData(lines []string) ([][]rune, string, Point) {
	var warehouse [][]rune
	var moves string
	var robot Point
	for i, line := range lines {
		if len(line) == 0 {
			moves = strings.Join(lines[i+1:], "")
			break
		}
		if strings.Contains(line, "@") {
			robot = Point{i, strings.Index(line, "@")}
			line = strings.Replace(line, "@", ".", -1)
		}

		warehouse = append(warehouse, []rune(line))
	}
	return warehouse, moves, robot
}
