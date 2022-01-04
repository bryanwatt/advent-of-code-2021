package main

import (
	"adventofcode/shared"
	"fmt"
	"strconv"
	"strings"
)

type bingoNumber struct {
	val    int
	marked bool
}

type bingoBoard struct {
	numbers [5][5]bingoNumber
}

func main() {
	puzzleInput, _ := shared.ReadFileLineToStrArr("day4_puzzle.txt")

	win, lose := playBingoAndGetWinningandLosingBoardTotals(puzzleInput)

	fmt.Printf("Winning Board Total: %v\n", win)
	fmt.Printf("Losing Board Total: %v\n", lose)
}

func playBingoAndGetWinningandLosingBoardTotals(bingoInput []string) (int, int) {
	// Get the Bingo Boards
	bingoBoards := getBingoBoards(bingoInput)

	// Get the bingo numbers
	bingoNumbers := getBingoNumbers(bingoInput)

	boardWinTotals := getBoardWinIndexsAndAmounts(bingoBoards, bingoNumbers)

	winningTotal := 0
	losingTotal := 0
	for x := 0; x < len(bingoBoards); x++ {
		if boardWinTotals[x][0] == 0 {
			winningTotal = boardWinTotals[x][1]
		}
		if boardWinTotals[x][0] == len(bingoBoards)-1 {
			losingTotal = boardWinTotals[x][1]
		}
	}
	return winningTotal, losingTotal
}

func getBoardWinIndexsAndAmounts(bingoBoards []bingoBoard, bingoNumbers []int) map[int][]int {
	boardWinMap := make(map[int][]int)
	boardWinCount := 0

	// Loop through each bingo number sequentially
	for numIndex := 0; numIndex < len(bingoNumbers); numIndex++ {
		bingoNumber := bingoNumbers[numIndex]

		// Loop through each bingo board
		for bIndex := 0; bIndex < len(bingoBoards); bIndex++ {
			// Ensure the board is not already "won"
			_, prs := boardWinMap[bIndex]
			if !prs {
				// Mark the numbers on the board
				markBingoBoardNumbers(&bingoBoards[bIndex], bingoNumber)

				// And check if the board now has bingo
				if boardHasBingo(bingoBoards[bIndex]) {
					// Calc the board total
					tot := calcBingoWinningBoardTotal(bingoBoards[bIndex], bingoNumber)
					// Add the amount and index to the map
					boardWinMap[bIndex] = []int{boardWinCount, tot}
					boardWinCount++

					if boardWinCount == len(bingoBoards) {
						return boardWinMap
					}
				}
			}
		}
	}

	panic("Could not finish all boards")
}

func markBingoBoardNumbers(board *bingoBoard, number int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if board.numbers[x][y].val == number {
				board.numbers[x][y].marked = true
			}
		}
	}
}

func boardHasBingo(board bingoBoard) bool {
	// 5x5 board, check each row/col
	for x := 0; x < 5; x++ {
		rowAllMarked := true
		colAllMarked := true
		for y := 0; y < 5; y++ {
			if !board.numbers[x][y].marked {
				rowAllMarked = false
			}
			if !board.numbers[y][x].marked {
				colAllMarked = false
			}
		}
		if rowAllMarked || colAllMarked {
			return true
		}
	}
	return false
}

func getBingoNumbers(bingoBoardInput []string) []int {
	// Line 1 is the bingo numbers
	numsStr := strings.Split(bingoBoardInput[0], ",")
	numArray := []int{}
	for i := 0; i < len(numsStr); i++ {
		intVal, _ := strconv.Atoi(numsStr[i])
		numArray = append(numArray, intVal)
	}
	return numArray
}

func calcBingoWinningBoardTotal(board bingoBoard, numberWon int) int {
	markedTotal := 0
	for rowi := 0; rowi < 5; rowi++ {
		for coli := 0; coli < 5; coli++ {
			// Only look for UNMARKED
			if !board.numbers[rowi][coli].marked {
				cellVal := board.numbers[rowi][coli].val
				markedTotal += cellVal
			}
		}
	}
	return markedTotal * numberWon
}

func getBingoBoards(bingoInput []string) []bingoBoard {
	var bingoBoards []bingoBoard
	for x := 2; x < len(bingoInput); x = x + 6 {
		var bingoBoard bingoBoard
		boardLine := 0
		// Each line of this board segment
		for y := x; y < x+5; y++ {
			// Get the line numbers
			bnums := getBingoBoardLineNumbers(bingoInput[y])

			// Set up each BingoBoard Number
			for z := 0; z < len(bnums); z++ {
				bingoBoard.numbers[boardLine][z] = bingoNumber{
					val:    bnums[z],
					marked: false,
				}
			}
			boardLine++
		}
		bingoBoards = append(bingoBoards, bingoBoard)
	}
	return bingoBoards
}

func getBingoBoardLineNumbers(line string) []int {
	// Board input are right justified, seperated by a space
	// So Length/spacing is constant
	// There are always 5
	// Like " 0 22  3 83 46"
	var returnaArr = []int{}
	for x := 0; x < len(line); x = x + 3 {
		strVal := line[x : x+2]
		intVal, _ := strconv.Atoi(strings.TrimSpace(strVal))
		returnaArr = append(returnaArr, intVal)
	}
	return returnaArr
}
