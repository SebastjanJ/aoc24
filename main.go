package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day := os.Args[1]
	switch day {
	case "1":
		Day1()
	case "2":
		Day2()
	case "3":
		Day3()
	case "4":
		Day4()
	case "5":
		Day5()
	case "6":
		Day6()
	case "7":
		Day7()
	default:
		fmt.Println("Day not implemented")
	}
}
