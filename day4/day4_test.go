package main

import (
	"adventofcode/shared"
	"testing"
)

func TestMainFunctionRuns(t *testing.T) {
	main()
}

func TestGetBingoBoardLineNumbersReturnsCorrectArray(t *testing.T) {
	testData := []struct {
		in  string
		exp []int
	}{
		{"50 98 55 14 47", []int{50, 98, 55, 14, 47}},
		{" 0 22  3 83 46", []int{0, 22, 3, 83, 46}},
	}

	for _, testCase := range testData {
		var res = getBingoBoardLineNumbers(testCase.in)
		if len(res) != 5 {
			t.Errorf("Line %v should always give 5 item array, got %v", testCase.in, len(res))
		}

		for x := 0; x < 5; x++ {
			if res[x] != testCase.exp[x] {
				t.Errorf("Item %v expected %v, got %v", x, testCase.exp[x], res[x])
			}
		}
	}
}

func TestGetBingoBoardsReturnsCorrectlySetUpBoards(t *testing.T) {
	testInputString, _ := shared.ReadFileLineToStrArr("day4_example.txt")

	boards := getBingoBoards(testInputString)

	if len(boards) != 3 {
		t.Errorf("There are supposed to be 3 board in the example input, got %v", len(boards))
	}

	// Spot check a few values
	if boards[0].numbers[0][1].val != 13 {
		t.Errorf("Board 1's first row second number should be 13, got %v", boards[0].numbers[0][1].val)
	}

	// Spot check a few values
	if boards[2].numbers[4][4].val != 7 {
		t.Errorf("Board 3's last row last number should be 7, got %v", boards[2].numbers[4][4].val)
	}
}

func TestCalcBingoWinningBoardTotal(t *testing.T) {
	testBoard := buildTestWinningBoard()

	numberWon := 24

	expectedRes := 4512

	actual := calcBingoWinningBoardTotal(testBoard, numberWon)

	if expectedRes != actual {
		t.Errorf("Winning Board total incorrect, expected %v, got %v", expectedRes, actual)
	}
}

func buildTestWinningBoard() bingoBoard {
	boardInput := []string{
		"Bla",
		"",
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	}
	testBoard := getBingoBoards(boardInput)[0]
	// Set the marked ones
	testBoard.numbers[0][0].marked = true
	testBoard.numbers[0][1].marked = true
	testBoard.numbers[0][2].marked = true
	testBoard.numbers[0][3].marked = true
	testBoard.numbers[0][4].marked = true
	testBoard.numbers[1][3].marked = true
	testBoard.numbers[2][2].marked = true
	testBoard.numbers[3][1].marked = true
	testBoard.numbers[3][4].marked = true
	testBoard.numbers[4][0].marked = true
	testBoard.numbers[4][1].marked = true
	testBoard.numbers[4][4].marked = true

	return testBoard
}

func TestBoardHasBingoWithRow(t *testing.T) {
	testBoard := buildTestWinningBoard()

	res := boardHasBingo(testBoard)

	if !res {
		t.Error("Board should have bingo")
	}
}

func TestBoardHasBingoWithCol(t *testing.T) {
	testBoard := buildTestWinningBoard()
	testBoard.numbers[0][4].marked = false

	// Ensure the test board does not have bingo now
	if boardHasBingo(testBoard) {
		t.Error("the test board should not have a bingo row")
	}

	testBoard.numbers[2][0].marked = true
	testBoard.numbers[2][1].marked = true
	testBoard.numbers[2][2].marked = true
	testBoard.numbers[2][3].marked = true
	testBoard.numbers[2][4].marked = true

	if !boardHasBingo(testBoard) {
		t.Error("Board should have a bingo col")
	}
}

func TestExampleGiven(t *testing.T) {
	testInputString, _ := shared.ReadFileLineToStrArr("day4_example.txt")

	win, _ := playBingoAndGetWinningandLosingBoardTotals(testInputString)

	if win != 4512 {
		t.Errorf("Example bingo input should have resulted in 4512, got %v\n", win)
	}
}

func TestPuzzleExample(t *testing.T) {
	exp := 31424

	testInputString, _ := shared.ReadFileLineToStrArr("day4_puzzle.txt")

	win, _ := playBingoAndGetWinningandLosingBoardTotals(testInputString)

	if win != exp {
		t.Errorf("Puzzle bingo input should have resulted in %v, got %v\n", exp, win)
	}
}

func TestExampleGivenPart2(t *testing.T) {
	testInputString, _ := shared.ReadFileLineToStrArr("day4_example.txt")

	_, lose := playBingoAndGetWinningandLosingBoardTotals(testInputString)

	if lose != 1924 {
		t.Errorf("Example bingo input should have resulted in 1924, got %v\n", lose)
	}
}

func TestPuzzleExamplePart2(t *testing.T) {
	exp := 23042

	testInputString, _ := shared.ReadFileLineToStrArr("day4_puzzle.txt")

	_, lose := playBingoAndGetWinningandLosingBoardTotals(testInputString)

	if lose != exp {
		t.Errorf("Puzzle bingo part 2 input should have resulted in %v, got %v\n", exp, lose)
	}
}
