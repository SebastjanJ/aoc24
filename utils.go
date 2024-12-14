package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func insert(slice []int, i int, value int) []int {
	// Grow the slice by one element
	slice = append(slice, 0)
	// Shift elements to make room
	copy(slice[i+1:], slice[i:])
	// Insert the value
	slice[i] = value
	return slice
}

func getInputIntsAll(day int) [][]int {
	var lines []string
	var ints [][]int
	if day == 0 {
		lines, _ = GetTest[string]()
	} else {
		lines = getInputStrings(day)
	}
	for _, line := range lines {
		var intline []int
		args := strings.Fields(line)
		for _, arg := range args {
			argint, _ := strconv.Atoi(arg)
			intline = append(intline, argint)
		}
		ints = append(ints, intline)
	}
	return ints
}

func getInputStrings(day int) []string {
	client, err := aocutil.NewInputFromFile("session.txt")
	if err != nil {
		log.Fatalf("Failed to initialize AoC client: %v", err)
	}

	lines, err := client.Strings(2024, day)
	if err != nil {
		log.Fatalf("Failed to fetch input: %v", err)
	}

	// Example processing of input
	return lines
}

func getInputInts(day int) []int {
	client, err := aocutil.NewInputFromFile("session.txt")
	if err != nil {
		log.Fatalf("Failed to initialize AoC client: %v", err)
	}

	lines, err := client.Ints(2024, day)
	if err != nil {
		log.Fatalf("Failed to fetch input: %v", err)
	}

	// Example processing of input
	return lines
}
func getInputFloats(day int) []float64 {
	client, err := aocutil.NewInputFromFile("session.txt")
	if err != nil {
		log.Fatalf("Failed to initialize AoC client: %v", err)
	}

	lines, err := client.Floats(2024, day)
	if err != nil {
		log.Fatalf("Failed to fetch input: %v", err)
	}

	// Example processing of input
	return lines
}

func GetTest[T any]() ([]T, error) {
	var data []T
	var err error

	// Read example input from a local file
	fileName := fmt.Sprintf("test.txt")
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read example input: %v", err)
	}

	switch any(data).(type) {
	case []string:
		data = any(parseStrings(content)).([]T)
	case []int:
		data = any(parseInts(content)).([]T)
	case []float64:
		data = any(parseFloats(content)).([]T)
	default:
		return nil, fmt.Errorf("unsupported type")
	}
	return data, nil
}

func parseStrings(content []byte) []string {
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func parseInts(content []byte) []int {
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	ints := make([]int, len(lines))
	for i, line := range lines {
		ints[i], _ = strconv.Atoi(line) // Handle errors as needed
		fmt.Println(ints[i])
	}
	return ints
}

func parseFloats(content []byte) []float64 {
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	floats := make([]float64, len(lines))
	for i, line := range lines {
		floats[i], _ = strconv.ParseFloat(line, 64) // Handle errors as needed
	}
	return floats
}
