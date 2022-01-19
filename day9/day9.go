package main

import (
	"adventofcode/shared"
	"fmt"
	"sort"
)

func main() {
	Day9("Example", "day9_example.txt")
	Day9("Puzzle", "day9_puzzle.txt")
}

func Day9(desc string, fileName string) {
	inputGrid, _ := shared.ReadFileLineToIntGrid(fileName)
	outputRes := calculateRiskLevel(inputGrid)

	fmt.Printf("Part1: %v Result output: %v \n", desc, outputRes)

	basinSize := calc3BiggestBasinsMultiplied(inputGrid)
	fmt.Printf("Part2: %v Result output: %v \n", desc, basinSize)
}

func Day92(desc string, fileName string) {
	inputGrid, _ := shared.ReadFileLineToIntGrid(fileName)
	outputRes := calculateRiskLevel(inputGrid)

	fmt.Printf("Part2: %v Result output: %v \n", desc, outputRes)
}

func calculateRiskLevel(intputGrid [][]int) int {
	riskLevelTotal := 0

	lowestPointList := getLowestPoints(intputGrid)

	for _, lowestPoint := range lowestPointList {
		val := intputGrid[lowestPoint[0]][lowestPoint[1]]
		riskLevelTotal = riskLevelTotal + 1 + val
	}

	return riskLevelTotal
}

func calc3BiggestBasinsMultiplied(intputGrid [][]int) int {
	lowestPointList := getLowestPoints(intputGrid)

	basinSizes := make([]int, 0)

	for _, lowestPoint := range lowestPointList {
		visitedValues := make(map[string]bool)
		basinCoOrdinates := getBasinCoOrginates(lowestPoint, intputGrid, visitedValues)
		basinSizes = append(basinSizes, len(basinCoOrdinates))
	}

	sort.Ints(basinSizes)

	return basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]
}

func getLowestPoints(intputGrid [][]int) [][2]int {
	lowestPointList := [][2]int{}

	for row := 0; row < len(intputGrid); row++ {
		for col := 0; col < len(intputGrid[row]); col++ {
			valLowest := true
			posValue := intputGrid[row][col]

			neighbours := getNeighbourCoOrds([2]int{row, col}, intputGrid)

			for _, neighbour := range neighbours {
				if intputGrid[neighbour[0]][neighbour[1]] <= posValue {
					valLowest = false
					continue
				}
			}

			if valLowest {
				point := [2]int{row, col}
				lowestPointList = append(lowestPointList, point)
			}

		}
	}

	return lowestPointList
}

func getNeighbourCoOrds(coords [2]int, intputGrid [][]int) [][2]int {
	returnCoOrdList := [][2]int{}

	numRows := len(intputGrid)
	numCols := len(intputGrid[coords[0]])

	// CoOrd := {row, col}
	if coords[0] > 0 {
		// Add the neighbour above
		returnCoOrdList = append(returnCoOrdList, [2]int{coords[0] - 1, coords[1]})
	}
	if coords[0] < numRows-1 {
		// Add the neighbour below
		returnCoOrdList = append(returnCoOrdList, [2]int{coords[0] + 1, coords[1]})
	}

	if coords[1] > 0 {
		// Add the neighbour left
		returnCoOrdList = append(returnCoOrdList, [2]int{coords[0], coords[1] - 1})
	}

	if coords[1] < numCols-1 {
		// Add the neighbour right
		returnCoOrdList = append(returnCoOrdList, [2]int{coords[0], coords[1] + 1})
	}

	return returnCoOrdList
}

func getBasinCoOrginates(coords [2]int, intputGrid [][]int, visitedCordsValues map[string]bool) [][2]int {
	// Get all the neighbours of this cord
	neighbours := getNeighbourCoOrds(coords, intputGrid)

	returnList := [][2]int{}

	for _, neighbour := range neighbours {
		// Ensure we are not visiting this value again
		mapKey := fmt.Sprintf("%v-%v", neighbour[0], neighbour[1])
		_, alreadyVisited := visitedCordsValues[mapKey]

		if !alreadyVisited {
			if intputGrid[neighbour[0]][neighbour[1]] < 9 {
				neighbourCoOrd := [2]int{neighbour[0], neighbour[1]}
				returnList = append(returnList, neighbourCoOrd)

				visitedCordsValues[mapKey] = true

				returnList = append(returnList, getBasinCoOrginates(neighbourCoOrd, intputGrid, visitedCordsValues)...)
			}
		}
	}
	return returnList
}
