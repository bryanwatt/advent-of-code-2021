package main

import (
	"adventofcode/shared"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputExample, _ := shared.ReadFileLineToStrArr("day5_example.txt")
	inputPuzzle, _ := shared.ReadFileLineToStrArr("day5_puzzle.txt")

	exampleRes := getNumberOfPointsWithMultipleOverlap(inputExample)
	puzzleRes := getNumberOfPointsWithMultipleOverlap(inputPuzzle)

	fmt.Printf("Example output: %v\n", exampleRes)
	fmt.Printf("Puzzle output: %v\n", puzzleRes)
}

func getNumberOfPointsWithMultipleOverlap(inputLines []string) int {
	positionGrid := buildInputPositions(inputLines)

	// Get the max value for a square grid
	gridSize := getMaxPositionVal(positionGrid) + 1

	// Create a grid for counting the positions
	// X is cols, Y is rows <- caught me out, but they are always equal it seems
	countGrid := make([][]int, gridSize)
	for i := range countGrid {
		countGrid[i] = make([]int, gridSize)
	}

	greaterThanOneCount := 0
	// Loop through all items, and increment each range counter
	for i := 0; i < len(positionGrid); i++ {
		// Get the x/y inc amounts (+1, -1, 0)
		xInc := getIncAmount(positionGrid[i][0][0], positionGrid[i][1][0])
		yInc := getIncAmount(positionGrid[i][0][1], positionGrid[i][1][1])

		// Start the positions at the first value co-ords
		xPos := positionGrid[i][0][0]
		yPos := positionGrid[i][0][1]

		for {
			// Chech the Value and increment if needed
			if countGrid[xPos][yPos] == 1 {
				greaterThanOneCount++
			}
			countGrid[xPos][yPos] = countGrid[xPos][yPos] + 1

			// Check if we have reached the end value co-ords
			if xPos == positionGrid[i][1][0] && yPos == positionGrid[i][1][1] {
				break
			}

			// Increment the Positions
			xPos = xPos + xInc
			yPos = yPos + yInc
		}
	}

	return greaterThanOneCount
}

func getIncAmount(startPos int, endPos int) int {
	if startPos > endPos {
		return -1
	}
	if endPos > startPos {
		return 1
	}
	return 0
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func getMaxPositionVal(positionGrid [][2][2]int) int {
	maxVal := 0

	for _, position := range positionGrid {
		maxVal = max(position[0][0], maxVal)
		maxVal = max(position[1][0], maxVal)
		maxVal = max(position[0][1], maxVal)
		maxVal = max(position[1][0], maxVal)
	}
	return maxVal
}

func buildInputPositions(inputLines []string) [][2][2]int {
	// Return an array of start and end co-ordinate pairs
	returnList := make([][2][2]int, len(inputLines))

	for i := 0; i < len(inputLines); i++ {
		line := inputLines[i]
		splits := strings.Split(line, " -> ")
		from := strings.Split(splits[0], ",")
		to := strings.Split(splits[1], ",")
		returnList[i] = [2][2]int{
			{getIntVal(from[0]), getIntVal(from[1])},
			{getIntVal(to[0]), getIntVal(to[1])},
		}
	}

	return returnList
}

func getIntVal(intString string) int {
	res, _ := strconv.Atoi(intString)
	return res
}
