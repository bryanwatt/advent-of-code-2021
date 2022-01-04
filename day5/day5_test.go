package main

import (
	"adventofcode/shared"
	"testing"
)

func TestMainFunctionRuns(t *testing.T) {
	main()
}

func TestExample(t *testing.T) {
	// Part 1:
	// exp := 5
	// Part 2:
	exp := 12

	testInputString, _ := shared.ReadFileLineToStrArr("day5_example.txt")

	res := getNumberOfPointsWithMultipleOverlap(testInputString)

	if res != exp {
		t.Errorf("Puzzle output expected %v, got %v\n", exp, res)
	}
}

func TestPuzzle(t *testing.T) {
	// Part 1: exp := 7674
	// Part 2:
	exp := 20898

	testInputString, _ := shared.ReadFileLineToStrArr("day5_puzzle.txt")

	res := getNumberOfPointsWithMultipleOverlap(testInputString)

	if res != exp {
		t.Errorf("Puzzle output expected %v, got %v\n", exp, res)
	}
}

func TestBuildInputPositions(t *testing.T) {
	lines := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
	}

	res := buildInputPositions(lines)

	if len(res) != 3 {
		t.Errorf("Return list should be 3 times long, got %v", len(res))
	}

	if len(res[0]) != 2 {
		t.Errorf("Return list should be of a pair of values, got %v", len(res[0]))
	}

	if len(res[0][1]) != 2 {
		t.Errorf("Each value should be a pair of ints, start and end, got %v", len(res[0][1]))
	}

	if res[1][0][0] != 8 || res[1][0][1] != 0 {
		t.Errorf("Line 2 first value should be [8, 0], got %v", res[1][0])
	}
	if res[2][1][0] != 3 || res[2][1][1] != 4 {
		t.Errorf("Line 3 second value should be [3, 4], got %v", res[2][1])
	}
}
