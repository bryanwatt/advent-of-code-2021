package shared

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		errorOut(err)
	}
	return file
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		errorOut(err)
	}
}

func errorOut(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadFileLineToStrArr(fileName string) ([]string, error) {
	file := openFile(fileName)
	defer closeFile(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// readLines reads a whole file into memory
// and returns an integer slice of its lines.
func ReadFileLineToIntArr(fileName string) ([]int, error) {
	file := openFile(fileName)
	defer closeFile(file)

	var intArr []int
	scanner := bufio.NewScanner(file)

	lineNo := 0
	for scanner.Scan() {
		intVal, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return nil, fmt.Errorf("could not convert %v to int value, line", err)
		}

		intArr = append(intArr, intVal)
		lineNo++
	}
	return intArr, scanner.Err()
}

// readLines reads a whole file into memory
// and returns an two integer slice of its lines.
func ReadFileLineToIntGrid(fileName string) ([][]int, error) {
	file := openFile(fileName)
	defer closeFile(file)

	intGrid := [][]int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// string to int array
		textLine := scanner.Text()
		intVals := make([]int, len(textLine))
		for i := 0; i < len(textLine); i++ {
			intVals[i], _ = strconv.Atoi(string(textLine[i]))
		}

		intGrid = append(intGrid, intVals)
	}
	return intGrid, scanner.Err()
}
