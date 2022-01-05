package main

import (
	"adventofcode/shared"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputExample, _ := shared.ReadFileLineToStrArr("day6_example.txt")
	inputPuzzle, _ := shared.ReadFileLineToStrArr("day6_puzzle.txt")

	example18Res := getFishCount(inputExample[0], 18)
	example80Res := getFishCount(inputExample[0], 80)
	puzzleRes := getFishCount(inputPuzzle[0], 80)
	puzzle256Res := getFishCount(inputPuzzle[0], 256)

	fmt.Printf("Example 18 days output: %v\n", example18Res)
	fmt.Printf("Example 80 days output: %v\n", example80Res)
	fmt.Printf("Puzzle 80 days output: %v\n", puzzleRes)
	fmt.Printf("Puzzle 256 days output: %v\n", puzzle256Res)
}

func getFishCount(startingFish string, daysToCalc int) int {
	fishArray := getIntArrayFromInput(startingFish)

	// Set up a fish "map" to store the number of fish at each age
	fm := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	// Add the initial fish day counter
	for _, dayVal := range fishArray {
		fm[dayVal]++
	}

	for i := 0; i < daysToCalc; i++ {
		// Keep swapping "down" the values,
		// although add the zero count to the prev 7 count to be the new 6
		t0, t1, t2, t3, t4, t5, t6, t7, t8 := fm[0], fm[1], fm[2], fm[3], fm[4], fm[5], fm[6], fm[7], fm[8]
		fm[8] = t0
		fm[7] = t8
		fm[6] = t7 + t0
		fm[5] = t6
		fm[4] = t5
		fm[3] = t4
		fm[2] = t3
		fm[1] = t2
		fm[0] = t1
	}

	return fm[0] + fm[1] + fm[2] + fm[3] + fm[4] + fm[5] + fm[6] + fm[7] + fm[8]
}

func getIntArrayFromInput(input string) []int {
	numsStr := strings.Split(input, ",")
	numArray := []int{}
	for i := 0; i < len(numsStr); i++ {
		intVal, _ := strconv.Atoi(numsStr[i])
		numArray = append(numArray, intVal)
	}
	return numArray
}
