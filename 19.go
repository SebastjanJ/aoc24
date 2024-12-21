package main

import (
	"fmt"
	"strings"
)

func Day19() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(19)
	designs, towels := getTowels(lines)
	cycleTowels(designs, towels)
}

func cycleTowels(designs []string, towels []string) {
	count1 := 0
	count2 := 0
	for _, towel := range towels {
		perms := checkTowel(designs, towel, make(map[string]int))
		if perms > 0 {
			count1++
		}
		count2 += perms
	}
	fmt.Println(count1)
	fmt.Println(count2)
}

func checkTowel(designs []string, towel string, memo map[string]int) int {
	if len(towel) == 0 {
		return 1
	}
	if val, exists := memo[towel]; exists {
		return val
	}
	count := 0
	for _, design := range designs {
		if strings.HasPrefix(towel, design) {
			count += checkTowel(designs, towel[len(design):], memo)
			memo[towel] = count

		}
	}
	return count
}

func getTowels(lines []string) ([]string, []string) {
	designs := make([]string, 0)
	designs = strings.Split(lines[0], ", ")
	return designs, lines[2:]
}
