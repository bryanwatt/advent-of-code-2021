package main

import (
	"adventofcode/shared"
	"fmt"
)

func main() {
	Day7("Example", "day7_example.txt")
	Day7("Puzzle", "day7_puzzle.txt")
}

func Day7(desc string, fileName string) {
	crabPosStr, _ := shared.ReadFileLineToStrArr(fileName)
	crabPosIntArr := shared.StringToIntArr(crabPosStr[0])

	position, fuelCost := getEfficientPosition(crabPosIntArr)

	fmt.Printf("%v output: Most Efficient Position: %v, Fuel Cost: %v \n", desc, position, fuelCost)
}

func getEfficientPosition(crabPositions []int) (int, int) {
	// Get the max horizontal Pos of the crabs
	maxX := 0

	// Create a map of the positions and crab counts
	posMap := make(map[int]int)
	for _, pos := range crabPositions {
		if pos > maxX {
			maxX = pos
		}
		posMap[pos] = posMap[pos] + 1
	}

	effFuelCost := 0
	effPos := 0
	// Iterate through each position, and calculate the total fuel of all crabs to get there
	for x := 0; x <= maxX; x++ {
		posFuelCost := 0
		for k, v := range posMap {
			// Distance to value
			dtv := Abs(k - x)

			// Calc the fuel cost to pos for 1 crab
			moveCost := 0
			for x := 1; x <= dtv; x++ {
				moveCost += x
			}

			fuelCost := moveCost * v
			posFuelCost += fuelCost

			// If it is already less efficient, break this loop early
			if posFuelCost > effFuelCost && effFuelCost != 0 {
				break
			}
		}

		if effFuelCost == 0 || posFuelCost < effFuelCost {
			effPos = x
			effFuelCost = posFuelCost
		}
	}

	return effPos, effFuelCost
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
