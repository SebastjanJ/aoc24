package main

import (
	"fmt"
	"strconv"
)

func Day9() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(9)
	files, spaces := getFilesAndSpace(lines[0])
	blocks := getBlocks(files, spaces)
	fmt.Println(partone9(blocks))
	fmt.Println(parttwo9(blocks))

}

func partone9(blocks []int) int {
	checksum := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == -1 {
			j := len(blocks) - 1
			for ; blocks[j] == -1; j-- {
			}
			checksum += blocks[j] * i
			blocks = blocks[:j]
			continue
		}
		checksum += blocks[i] * i
	}
	return checksum
}

func parttwo9(blocks []int) int {
	checksum := 0

blockslopp:
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == -2 {
			continue
		}
		if blocks[i] == -1 {
			j := i
			for ; blocks[j] == -1; j++ {
			}
			spacecount := j - i
			lastcount := 1
			last := len(blocks) - 2
			for ; last > i; last-- {
				if blocks[last+1] == -1 || blocks[last+1] == -2 {
					lastcount = 1
					continue
				}
				if blocks[last] == blocks[last+1] {
					lastcount++
				} else {
					if lastcount <= spacecount {
						for k := i; k < i+lastcount; k++ {
							checksum += blocks[last+k-i+1] * k
							blocks[last+k-i+1] = -2
						}
						i = i + lastcount - 1

						continue blockslopp
					}
					lastcount = 1
				}
			}
			continue
		}

		checksum += blocks[i] * i
	}
	return checksum
}

func getBlocks(files, spaces []int) []int {
	blocks := make([]int, 0)
	for i, file := range files {
		for j := 0; j < file; j++ {
			blocks = append(blocks, i)
		}
		if len(spaces) > 0 {
			space := spaces[0]
			for j := 0; j < space; j++ {
				blocks = append(blocks, -1)
			}
			spaces = spaces[1:]
		}
	}
	return blocks
}

func getFilesAndSpace(line string) ([]int, []int) {
	files := []int{}
	spaces := []int{}

	for i, char := range line {
		num, _ := strconv.Atoi(string(char))
		if i%2 == 0 {
			files = append(files, num)
		} else {
			spaces = append(spaces, num)
		}
	}
	return files, spaces
}
