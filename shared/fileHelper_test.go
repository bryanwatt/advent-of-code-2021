package shared

import (
	"testing"
)

func TestReadFileLineToStrArrGetStringResults(t *testing.T) {
	fileName := "file1_test.txt"
	fileStringArr, err := ReadFileLineToStrArr(fileName)

	if fileStringArr == nil || err != nil {
		t.Error("Failed to get a result from file")
	}

	if len(fileStringArr) != 3 {
		t.Errorf("Number of lines incorrect, got: %d expected: 3", len(fileStringArr))
	}
}

func TestReadFileLineToIntArrGetsIntResults(t *testing.T) {
	fileName := "file2_test.txt"
	fileIntArray, err := ReadFileLineToIntArr(fileName)

	if fileIntArray == nil || err != nil {
		t.Error("Failed to get a result from file")
	}

	if len(fileIntArray) != 5 {
		t.Errorf("Number of items incorrect, got: %d expected: 5", len(fileIntArray))
	}

	if fileIntArray[2] != 300 {
		t.Errorf("Third item in the array is incorrect, got: %d expected: 300", fileIntArray[2])
	}
}

func TestReadFileLineToIntGrid(t *testing.T) {
	fileName := "file3_test.txt"
	fileIntGrid, err := ReadFileLineToIntGrid(fileName)

	if fileIntGrid == nil || err != nil {
		t.Error("Failed to get a result from file")
	}

	if len(fileIntGrid) != 4 {
		t.Errorf("Number of items incorrect, got: %d expected: 5", len(fileIntGrid))
	}

	if len(fileIntGrid[2]) != 10 {
		t.Errorf("Second dimenension lenth got: %d expected: 10", len(fileIntGrid[2]))
	}

	if fileIntGrid[3][3] != 7 {
		t.Errorf("3 3 value got: %d expected: 7", fileIntGrid[3][3])
	}
}
