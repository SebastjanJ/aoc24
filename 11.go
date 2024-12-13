package main

import (
	"fmt"
	"math"
)

func Day11() {
	// lines := getInputIntsAll(0)[0]
	lines := getInputIntsAll(11)[0]
	partonetwo11(lines)
}

func partonetwo11(stones []int) {
	count25 := 0
	count75 := 0
	cache := make(map[[2]int]int)

	for _, stone := range stones {
		key := [2]int{stone, 25}
		if val, exists := cache[key]; exists {
			count25 += val
		} else {
			stonecount := breakStone(stone, 25, cache)
			cache[key] = stonecount
			count25 += stonecount
		}

		key = [2]int{stone, 75}
		if val, exists := cache[key]; exists {
			count75 += val
			continue
		} else {
			stonecount := breakStone(stone, 75, cache)
			cache[key] = stonecount
			count75 += stonecount
		}
	}
	fmt.Println(count25)
	fmt.Println(count75)
}

func breakStone(stone, i int, cache map[[2]int]int) int {
	if i == 0 {
		return 1
	}
	key := [2]int{stone, i}
	if val, exists := cache[key]; exists {
		return val
	}
	var count int
	if stone == 0 {
		count = breakStone(1, i-1, cache)
	} else if len(fmt.Sprint(stone))%2 == 0 {
		curr := stone
		magic := int(math.Pow10(len(fmt.Sprint(curr)) / 2))
		count = breakStone(curr/magic, i-1, cache) + breakStone(curr%magic, i-1, cache)
	} else {
		count = breakStone(stone*2024, i-1, cache)
	}
	cache[key] = count
	return count
}
