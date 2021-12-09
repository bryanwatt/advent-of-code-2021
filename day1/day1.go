package main

import (
	"adventofcode/shared"
	"fmt"
)

func main() {
	fmt.Println("Also check the unit tests with -v to see test output")
	loadAndCalculateSlidingDepthIncreases("day1_example.txt", "Part 1: Example", 1)
	loadAndCalculateSlidingDepthIncreases("day1_puzzle.txt", "Part 1: Puzzle ", 1)
	loadAndCalculateSlidingDepthIncreases("day1_example.txt", "Part 2: Example", 3)
	loadAndCalculateSlidingDepthIncreases("day1_puzzle.txt", "Part 2: Puzzle ", 3)
}

func loadAndCalculateSlidingDepthIncreases(fileName string, descriptor string, window int) {
	fileValues, err := shared.ReadFileLineToIntArr(fileName)

	if err != nil {
		panic((err))
	}

	result := calcConsecutiveSlidingDepthIncreases(fileValues, window)

	fmt.Printf("%v: %v increases of %v window value(s)\n", descriptor, result, window)
}

func calcConsecutiveSlidingDepthIncreases(intArr []int, slideScale int) int {
	if len(intArr) < 1 {
		return 0
	}

	cnt := 0

	for i := 1; i < len(intArr)-slideScale+1; i++ {
		// previous window total
		prev := sumArrayPoints(intArr, i-1, slideScale)
		// current window total
		curr := sumArrayPoints(intArr, i, slideScale)

		if curr > prev {
			cnt++
		}
	}

	return cnt
}

func sumArrayPoints(intArr []int, startIndex int, noToTake int) int {
	tot := 0
	for i := 0; i < noToTake; i++ {
		tot += intArr[startIndex+i]
	}
	return tot
}
