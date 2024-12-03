package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day3() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(3)
	partone3(lines)
	parttwo3(lines)
}

func partone3(lines []string) {
	sum := 0
	mul := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	num := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		matches := mul.FindAllString(line, -1)
		for _, match := range matches {
			nums := num.FindAllString(match, -1)
			left, _ := strconv.Atoi(nums[0])
			right, _ := strconv.Atoi(nums[1])
			sum += left * right
		}
	}
	fmt.Println(sum)
}

func parttwo3(lines []string) {
	sum := 0
	mul := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	num := regexp.MustCompile(`\d+`)
	do := true
	for _, line := range lines {
		matches := mul.FindAllString(line, -1)
		for _, match := range matches {
			nums := num.FindAllString(match, -1)
			if len(nums) == 0 {
				if match == "do()" {
					do = true
				} else {
					do = false
				}
				continue
			}
			if do {
				left, _ := strconv.Atoi(nums[0])
				right, _ := strconv.Atoi(nums[1])
				sum += left * right
			}
		}
	}
	fmt.Println(sum)
}
