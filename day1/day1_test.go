package main

import (
	"adventofcode/shared"
	"testing"
)

func TestSumArrayPointsGetsCorrectValues(t *testing.T) {
	intArr := []int{10, 35, 55, 86, 57, 66, 78, 81, 93, 101}

	// should sum 55,86,57 = 198
	res := sumArrayPoints(intArr, 2, 3)

	if res != 198 {
		t.Errorf("Result should total 198, got %d", res)
	}
}

func TestPart1DepthIncreasesForExample(t *testing.T) {
	intputValues, err := shared.ReadFileLineToIntArr("day1_example.txt")

	if err != nil {
		t.Error(err)
	}

	result := calcConsecutiveSlidingDepthIncreases(intputValues, 1)

	if result != 7 {
		t.Errorf("Example for depth increases should be 7, got %d", result)
	}

	t.Logf("Example depth increases tested returned the expected 7")
}

func TestPart1DepthIncreasesForQuestion(t *testing.T) {
	intputValues, err := shared.ReadFileLineToIntArr("day1_puzzle.txt")

	if err != nil {
		t.Error(err)
	}

	result := calcConsecutiveSlidingDepthIncreases(intputValues, 1)

	t.Logf("Part 1 Puzzle returned %d", result)
}

func TestPart2SlidingWindowDepthInceasesForExample(t *testing.T) {
	intputValues, err := shared.ReadFileLineToIntArr("day1_example.txt")

	if err != nil {
		t.Error(err)
	}

	result := calcConsecutiveSlidingDepthIncreases(intputValues, 3)

	if result != 5 {
		t.Errorf("Example for depth increases should be 5, got %d", result)
	}

	t.Logf("Example depth increases tested returned the expected 5")
}

func TestPart2SlidingWindowDepthIncreasesForQuestion(t *testing.T) {
	intputValues, err := shared.ReadFileLineToIntArr("day1_puzzle.txt")

	if err != nil {
		t.Error(err)
	}

	result := calcConsecutiveSlidingDepthIncreases(intputValues, 3)

	t.Logf("Part 2 Puzzle returned %d", result)
}
