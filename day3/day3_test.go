package main

import (
	"adventofcode/shared"
	"testing"
)

func TestPart1InputExample(t *testing.T) {
	// This value is given in the example
	expPowerConsumption := 198
	inputRows, _ := shared.ReadFileLineToStrArr("day3_example.txt")

	if inputRows == nil || !(len(inputRows) > 0) {
		t.Error("No input received")
	}

	res := calculatePowerConsumption(inputRows)

	if expPowerConsumption != res {
		t.Errorf("Power Consumption calculated incorrect, expected %v, got %v", expPowerConsumption, res)
	}
}

func TestPart1PuzzleExample(t *testing.T) {
	// This answer was marked correct on submission
	expPowerConsumption := 1071734

	inputRows, _ := shared.ReadFileLineToStrArr("day3_puzzle.txt")

	if inputRows == nil || !(len(inputRows) > 0) {
		t.Error("No input received")
	}

	res := calculatePowerConsumption(inputRows)

	if expPowerConsumption != res {
		t.Errorf("Power Consumption calculated incorrect, expected %v, got %v", expPowerConsumption, res)
	}
}

func TestPart2InputExample(t *testing.T) {
	// This value is given in the example
	expLifeSupportValue := 230
	inputRows, _ := shared.ReadFileLineToStrArr("day3_example.txt")

	if inputRows == nil || !(len(inputRows) > 0) {
		t.Error("No input received")
	}

	res := calculateLifeSupportRating(inputRows)

	if int64(expLifeSupportValue) != res {
		t.Errorf("Life Support Value calculated incorrect, expected %v, got %v", expLifeSupportValue, res)
	}
}

func TestPart2PuzzleExample(t *testing.T) {
	// This answer was marked correct on submission
	expLifeSupportValue := 6124992
	inputRows, _ := shared.ReadFileLineToStrArr("day3_puzzle.txt")

	if inputRows == nil || !(len(inputRows) > 0) {
		t.Error("No input received")
	}

	res := calculateLifeSupportRating(inputRows)

	if int64(expLifeSupportValue) != res {
		t.Errorf("Life Support Value calculated incorrect, expected %v, got %v", expLifeSupportValue, res)
	}
}
