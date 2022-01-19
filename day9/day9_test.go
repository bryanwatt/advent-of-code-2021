package main

import (
	"adventofcode/shared"
	"testing"
)

func TestMainFunctionRuns(t *testing.T) {
	main()
}

func TestPart1ExampleGivenGivesRightResult(t *testing.T) {
	expectedAnswer := 15
	runPart1Test("day9_example.txt", expectedAnswer, t)
}

func TestPart1PuzzleAnswerGivenGivesRightResult(t *testing.T) {
	expectedAnswer := 545 //Obtained after calculating

	runPart1Test("day9_puzzle.txt", expectedAnswer, t)
}

func TestPart2ExampleGivenGivesRightResult(t *testing.T) {
	expectedAnswer := 1134
	runPart2Test("day9_example.txt", expectedAnswer, t)
}

func TestPart2PuzzleAnswerGivenGivesRightResult(t *testing.T) {
	expectedAnswer := 950600 //Obtained after calculating

	runPart2Test("day9_puzzle.txt", expectedAnswer, t)
}

func runPart1Test(fileName string, expectedAnswer int, t *testing.T) {
	inputGrid, _ := shared.ReadFileLineToIntGrid(fileName)
	outputRes := calculateRiskLevel(inputGrid)

	if outputRes != expectedAnswer {
		t.Errorf("%v expected %v, got %v", fileName, expectedAnswer, outputRes)
	}
}

func runPart2Test(fileName string, expectedAnswer int, t *testing.T) {
	inputGrid, _ := shared.ReadFileLineToIntGrid(fileName)
	outputRes := calc3BiggestBasinsMultiplied(inputGrid)

	if outputRes != expectedAnswer {
		t.Errorf("%v expected %v, got %v", fileName, expectedAnswer, outputRes)
	}
}

func TestGetNeighbourCoOrds(t *testing.T) {
	inputGrid, _ := shared.ReadFileLineToIntGrid("day9_example.txt")

	runGetNeighburCoOrdsTest(t, inputGrid, [2]int{0, 0}, 2, [2]int{0, 1}, [2]int{1, 0})
	runGetNeighburCoOrdsTest(t, inputGrid, [2]int{0, 1}, 3, [2]int{0, 0}, [2]int{0, 2}, [2]int{1, 1})
	runGetNeighburCoOrdsTest(t, inputGrid, [2]int{0, 9}, 2, [2]int{0, 8}, [2]int{1, 9})
	runGetNeighburCoOrdsTest(t, inputGrid, [2]int{3, 3}, 4, [2]int{3, 2}, [2]int{3, 4}, [2]int{2, 3}, [2]int{4, 3})
	runGetNeighburCoOrdsTest(t, inputGrid, [2]int{4, 9}, 2)
	runGetNeighburCoOrdsTest(t, inputGrid, [2]int{3, 0}, 3)
}

func runGetNeighburCoOrdsTest(t *testing.T, inputGrid [][]int, coOrd [2]int, expCount int, expectedCord ...[2]int) {
	neighbours := getNeighbourCoOrds(coOrd, inputGrid)

	if len(neighbours) != expCount {
		t.Errorf("CoOrd %v %v expected %v neighbouts, got %v", coOrd[0], coOrd[1], expCount, len(neighbours))
	}

	for _, searchCord := range expectedCord {
		found := false

		for _, neighbour := range neighbours {
			if searchCord[0] == neighbour[0] && searchCord[1] == neighbour[1] {
				found = true
			}
		}

		if !found {
			t.Errorf("Could not find cord %v %v in neighbour list, got %v", searchCord[0], searchCord[1], neighbours)
		}
	}
}

func TestGetBasinCoOrginates(t *testing.T) {
	inputGrid, _ := shared.ReadFileLineToIntGrid("day9_example.txt")

	runGetBasinCoOrginatesTest(inputGrid, [2]int{0, 1}, 3, t)
	runGetBasinCoOrginatesTest(inputGrid, [2]int{0, 9}, 9, t)
	runGetBasinCoOrginatesTest(inputGrid, [2]int{2, 2}, 14, t)
	runGetBasinCoOrginatesTest(inputGrid, [2]int{4, 6}, 9, t)
}

func runGetBasinCoOrginatesTest(inputGrid [][]int, coOrd [2]int, expCount int, t *testing.T) {
	visitedValues := make(map[string]bool)

	resBasinValues := getBasinCoOrginates(coOrd, inputGrid, visitedValues)

	if len(resBasinValues) != expCount {
		t.Errorf("CoOrd %v expected %v count, got %v", coOrd, expCount, len(resBasinValues))
	}
}

func TestCalc3BiggestBasinsMultiplied(t *testing.T) {
	inputGrid, _ := shared.ReadFileLineToIntGrid("day9_example.txt")

	exp := 1134

	res := calc3BiggestBasinsMultiplied(inputGrid)

	if exp != res {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
