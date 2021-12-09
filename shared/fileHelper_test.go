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
