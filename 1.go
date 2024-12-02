package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func Day1() {
	var left []int
	var right []int

	lines := getInputIntsAll(1)

	for _, line := range lines {
		left = append(left, line[0])
		right = append(right, line[1])
	}

	sort.Ints(left)
	sort.Ints(right)

	partone1(left, right)
	parttwo1(left, right)

}

func partone1(left []int, right []int) {
	diff := 0
	for i := 0; i < len(left); i++ {
		diff += int(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println("Part One:", diff)

}

func parttwo1(left []int, right []int) {
	simi := 0
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
