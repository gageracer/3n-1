package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	startArg, endArg := 1, 5
	showResults := true
	matrix := false
	args := os.Args[1:]
	//fmt.Println(len(args))

	if len(args) > 0 {
		if args[0] == "no" || args[0] == "h" || args[0] == "n" {
			showResults = false
		} else if args[0] == "m" {
			matrix = true
			showResults = false
		} else {
			startArg, _ = strconv.Atoi(args[0])
			startArg = min(startArg, 1)
			endArg = startArg * 2
		}
	}
	if len(args) > 1 {
		if args[1] == "no" || args[1] == "h" || args[1] == "n" {
			showResults = false
		} else if args[1] == "m" {
			matrix = true
			showResults = false
		} else {
			endArg, _ = strconv.Atoi(args[1])
			endArg = min(endArg, startArg)
		}
	}
	if len(args) > 2 {
		if args[2] == "no" || args[2] == "h" || args[2] == "n" {
			showResults = false
		} else if args[2] == "m" {
			matrix = true
			showResults = false
		}
		// fmt.Println(showResults)
	}

	for start := startArg; start <= endArg; start++ {
		round := 1

		for num := start; num != 1; round++ {
			if num%2 == 0 {
				num = num / 2
			} else {
				num = (3 * num) + 1
			}
			if matrix {
				fmt.Print(round, num)
			}
		}
		if showResults {
			fmt.Printf("%v took %v rounds to reach the loop.\n",
				start, round)
			fmt.Println("------------------------------------")
		}
	}
	fmt.Printf("Last number is %v.\n", startArg)
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}