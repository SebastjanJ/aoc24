package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(1)
	var left []int
	var right []int
	diff := 0
	simi := 0

	for _, line := range lines {
		args := strings.Fields(line)
		leftint, _ := strconv.Atoi(args[0])
		left = append(left, leftint)
		rightint, _ := strconv.Atoi(args[1])
		right = append(right, rightint)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		diff += int(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println("Part One:", diff)

	for _, num := range left {
		i, found := slices.BinarySearch(right, num)
		if found {
			count := 0
			for ; right[i] == num; i++ {
				count += 1
			}
			simi += count * num
		}
	}
	fmt.Println("Part Two:", simi)
}
