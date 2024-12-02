package main

import (
	"fmt"
	"math"
	"slices"
)

func Day2() {
	lines := getInputIntsAll(2)
	partone2(lines)
	parttwo2(lines)
}

func partone2(lines [][]int) {
	count := 0
	for _, nums := range lines {
		if checkSafe(nums) {
			count += 1
		}
	}
	fmt.Println("Part One: ", count)
}

func parttwo2(lines [][]int) {
	count := 0
	for _, nums := range lines {
		if checkSafe(nums) {
			count += 1
		} else {
			for i := 0; i < len(nums); i++ {
				tmpnums := append(append([]int(nil), nums[:i]...), nums[i+1:]...)
				if checkSafe(tmpnums) {
					count += 1
					break
				}
			}
		}

	}
	fmt.Println("Part Two: ", count)
}

func checkLevels(a int, b int) bool {
	return a == b || math.Abs(float64(a-b)) > 3
}

func checkSafe(nums []int) bool {
	reversed := append([]int(nil), nums...)
	slices.Reverse(reversed)
	if slices.IsSorted(nums) || slices.IsSorted(reversed) {
		for i := 1; i < len(nums); i++ {
			if checkLevels(nums[i], nums[i-1]) {
				return false
			}
		}
		return true
	}
	return false
}
