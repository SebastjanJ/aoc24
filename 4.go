package main

import "fmt"

func Day4() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(4)
	partone4(lines)
	parttwo4(lines)
}

func partone4(lines []string) {
	count := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 'X' {
				count += wordsearch(lines, i, j)
			}
		}
	}
	fmt.Println(count)
}

func parttwo4(lines []string) {
	count := 0
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] == 'A' {
				if mas(lines, i, j) {
					count += 1
				}
			}
		}
	}
	fmt.Println(count)
}

func mas(lines []string, i int, j int) bool {
	return (lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S' ||
		lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M') &&
		(lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S' ||
			lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M')
}

func leftToRight(lines []string, i int, j int) bool {
	return j+3 < len(lines[i]) && lines[i][j:j+4] == "XMAS"
}
func rightToleft(lines []string, i int, j int) bool {
	return j > 2 && lines[i][j-3:j+1] == "SAMX"
}
func wordsearch(lines []string, i int, j int) int {
	count := 0
	if leftToRight(lines, i, j) {
		count += 1
	}
	if rightToleft(lines, i, j) {
		count += 1
	}
	if i > 2 {
		if lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
			count += 1
		}
		if j > 2 && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
			count += 1
		}
		if j < len(lines[i])-3 && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
			count += 1
		}
	}
	if i+3 < len(lines) {
		if lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
			count += 1
		}
		if j > 2 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
			count += 1
		}
		if j < len(lines[i])-3 && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
			count += 1
		}
	}
	return count
}
