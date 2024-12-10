package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day7() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(7)
	partonetwo7(lines)
}

func partonetwo7(lines []string) {
	partonecount := 0
	parttwocount := 0
	for _, line := range lines {
		args := strings.Split(line, ":")
		test, _ := strconv.Atoi(args[0])
		nums := strings.Fields(args[1])
		sum, _ := strconv.Atoi(nums[0])
		if rec(test, nums[1:], sum, false) {
			partonecount += test
		}
		if rec(test, nums[1:], sum, true) {
			parttwocount += test
		}
	}
	fmt.Println("Part one: ", partonecount)
	fmt.Println("Part one: ", parttwocount)
}

func rec(test int, nums []string, sum int, parttwo bool) bool {
	if sum > test {
		return false
	}
	if len(nums) == 0 {
		return sum == test
	}

	numint, _ := strconv.Atoi(nums[0])
	if rec(test, nums[1:], sum*numint, parttwo) {
		return true
	}
	if rec(test, nums[1:], sum+numint, parttwo) {
		return true
	}
	if parttwo && rec(test, nums[1:], sum*int(math.Pow10(len(nums[0])))+numint, parttwo) {
		return true
	}
	return false
}
