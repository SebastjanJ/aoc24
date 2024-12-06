package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	before int
	after  int
}

type Update struct {
	pages []int
}

func Day5() {
	lines := getInputStrings(5)
	rules, updates := parseInput(lines)
	fmt.Println(partOne(rules, updates))
	fmt.Println(partTwo(rules, updates))
}

func partOne(rules []Rule, updates []Update) int {
	sum := 0
	for _, update := range updates {
		if isCorrectlyOrdered(update, rules) {
			sum += getMiddle(update)
		}
	}
	return sum
}

func partTwo(rules []Rule, updates []Update) int {
	sum := 0
	for _, update := range updates {
		if !isCorrectlyOrdered(update, rules) {
			ordered := reorderUpdate(update, rules)
			sum += getMiddle(ordered)
		}
	}
	return sum
}

func parseInput(lines []string) ([]Rule, []Update) {
	var rules []Rule
	var updates []Update

	i := 0
	for ; i < len(lines) && len(lines[i]) > 0; i++ {
		parts := strings.Split(lines[i], "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules = append(rules, Rule{before, after})
	}

	for _, line := range lines[i+1:] {
		var pages []int
		for _, num := range strings.Split(line, ",") {
			page, _ := strconv.Atoi(num)
			pages = append(pages, page)
		}
		updates = append(updates, Update{pages})
	}

	return rules, updates
}

func isCorrectlyOrdered(update Update, rules []Rule) bool {
	for i := 0; i < len(update.pages); i++ {
		for j := i + 1; j < len(update.pages); j++ {
			if violatesRule(update.pages[i], update.pages[j], rules) {
				return false
			}
		}
	}
	return true
}

func violatesRule(before, after int, rules []Rule) bool {
	for _, rule := range rules {
		if rule.after == before && rule.before == after {
			return true
		}
	}
	return false
}

func reorderUpdate(update Update, rules []Rule) Update {
	pages := make([]int, len(update.pages))
	copy(pages, update.pages)

	changed := true
	for changed {
		changed = false
		for i := 0; i < len(pages); i++ {
			for j := i + 1; j < len(pages); j++ {
				if violatesRule(pages[i], pages[j], rules) {
					pages[i], pages[j] = pages[j], pages[i]
					changed = true
				}
			}
		}
	}

	return Update{pages}
}

func getMiddle(update Update) int {
	return update.pages[len(update.pages)/2]
}
