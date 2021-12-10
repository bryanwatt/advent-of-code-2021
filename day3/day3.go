package main

import (
	"adventofcode/shared"
	"fmt"
	"strconv"
)

func main() {
	powerValueList, _ := shared.ReadFileLineToStrArr("day3_puzzle.txt")

	pwr := calculatePowerConsumption(powerValueList)

	fmt.Printf("Power Consumption: %v\n", pwr)
}

func calculatePowerConsumption(inputValues []string) int {
	// Power consumption is calculated by multiplying the Gamme & Epsilon values
	// Gamma and Epsilon are calculated from selecting the most common and least common
	// bits respectively in each position in the whole input list
	// The results in a bit arrray, which then gets converted to decimal
	valueArray := make([][]int, len(inputValues[0]))

	for _, line := range inputValues {
		for x := 0; x < len(line); x++ {
			if len(valueArray[x]) == 0 {
				valueArray[x] = make([]int, 2)
			}

			intVal := 0
			if line[x] == 49 {
				intVal = 1
			}
			valueArray[x][intVal] = valueArray[x][intVal] + 1
		}
	}

	gammaBitString, epsilonBitString := calculateGammaAndEpsilonRates(valueArray)
	gammaInt64, _ := strconv.ParseInt(gammaBitString, 2, 32)
	epsilonInt64, _ := strconv.ParseInt(epsilonBitString, 2, 32)

	return int(gammaInt64 * epsilonInt64)
}

func calculateGammaAndEpsilonRates(inputTotals [][]int) (string, string) {
	gamma := ""
	epsilon := ""

	for i := 0; i < len(inputTotals); i++ {
		if inputTotals[i][0] > inputTotals[i][1] {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
	}

	return gamma, epsilon
}

func calculateLifeSupportRating(inputList []string) int64 {
	oxygenGenRatingBits := getBitStringByMostCommonBit(inputList, 0, true)
	co2ScrubberRatingBits := getBitStringByMostCommonBit(inputList, 0, false)

	oxygenGenRating, _ := strconv.ParseInt(oxygenGenRatingBits[0], 2, 32)
	co2ScrubberRating, _ := strconv.ParseInt(co2ScrubberRatingBits[0], 2, 32)

	return oxygenGenRating * co2ScrubberRating
}

func getBitStringByMostCommonBit(inputList []string, pos int, searchMostCommon bool) []string {
	countOfZero := 0
	countOfOne := 0
	valueArray := make([][]string, 2)
	for i := 0; i < len(inputList); i++ {
		switch inputList[i][pos] {
		case 48:
			countOfZero++
			valueArray[0] = append(valueArray[0], inputList[i])
		case 49:
			countOfOne++
			valueArray[1] = append(valueArray[1], inputList[i])
		default:
			panic("Value not 0 or 1")
		}
	}
	bitToUse := 0
	if searchMostCommon {
		if countOfOne >= countOfZero {
			bitToUse = 1
		}
	} else {
		if countOfOne < countOfZero {
			bitToUse = 1
		}
	}

	// If there is only 1 value left, return it
	if len(valueArray[bitToUse]) == 1 {
		return valueArray[bitToUse]
	}

	// If we've gotten to the end of character positions,
	// the remaining values should be the same, and exit
	if pos == len(valueArray[bitToUse][0])-1 {
		return valueArray[bitToUse]
	}

	// Recursively go to the next bit position
	return getBitStringByMostCommonBit(valueArray[bitToUse], pos+1, searchMostCommon)
}
