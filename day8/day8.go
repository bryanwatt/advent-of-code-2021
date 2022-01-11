package main

import (
	"adventofcode/shared"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	Day8("Example", "day8_example.txt")
	Day8("Puzzle", "day8_puzzle.txt")
}

func Day8(desc string, fileName string) {
	intputText, _ := shared.ReadFileLineToStrArr(fileName)

	outputRes := totaldisplayOutput(intputText)

	fmt.Printf("%v Result output: %v \n", desc, outputRes)
}

var stdSegmentMap = map[int]string{
	0: "abcefg",
	1: "cf",
	2: "acdeg",
	3: "acdfg",
	4: "bcdf",
	5: "abdfg",
	6: "abdefg",
	7: "acf",
	8: "abcdefg",
	9: "abcdfg",
}

var inputSegmentMap = map[string]int{}

func totaldisplayOutput(inputLines []string) int {
	tot := 0

	for x := 0; x < len(inputLines); x++ {
		inputSplit := strings.Split(inputLines[x], " | ")
		preDigits := strings.Split(inputSplit[0], " ")
		digits := strings.Split(inputSplit[1], " ")

		segMap := make(map[string]string)

		// THIS IS DIRTY filthy CODE!

		// Difference between len2 and len3 will give "a"
		len2 := getSegmentsWithLength(preDigits, 2)[0]
		len3 := getSegmentsWithLength(preDigits, 3)[0]
		len4 := getSegmentsWithLength(preDigits, 4)[0]
		len5s := getSegmentsWithLength(preDigits, 5)
		len6s := getSegmentsWithLength(preDigits, 6)
		len7 := getSegmentsWithLength(preDigits, 7)[0]

		len23Dif := difference(len2, len3)
		if len(len23Dif) != 1 {
			panic("len 2 and 3 should only have 1 diff")
		}
		segMap["a"] = len23Dif[0]

		// a len6 (0,6,9) NOT matching the segments of len2(1) = 6, meaning dif = "c"
		for _, len6val := range len6s {
			len6Dif := differenceIn(len6val, len2)
			if len(len6Dif) > 0 {
				segMap["c"] = len6Dif[0]
			}
		}

		// then the other seg of len2(1) = "f"
		fVal := differenceIn(segMap["c"], len2)
		if len(fVal) != 1 {
			panic("No f Val?")
		}
		segMap["f"] = fVal[0]

		// Now we have ACF
		// A 5 length segment that has all the values of "7" must be 3
		for _, len5Val := range len5s {
			len5Diff := differenceIn(len5Val, len3)

			if len(len5Diff) == 0 {
				// this is a 3
				// the value both 3 & 4 DONT have is e
				threeFourTogether := addTogether(len5Val, len4)

				// The value not in there (use 8s as it has it all), is e
				threeFourEightDif := differenceIn(threeFourTogether, len7)
				if len(threeFourEightDif) != 1 {
					panic("cant find diff in 3 & 4 to 8")
				}
				segMap["e"] = threeFourEightDif[0]

				// The 3 does not have b & e, but we know what e is now
				threeEightDiff := differenceIn(len5Val, len7)
				if len(threeEightDiff) != 2 {
					panic("threeEightDiff should always be 2")
				}
				if threeEightDiff[0] == segMap["e"] {
					segMap["b"] = threeEightDiff[1]
				} else {
					segMap["b"] = threeEightDiff[0]
				}
			}
		}

		// Now we have ABCEF, solve d&g
		// If we add the segments known for A & E to the what we know is 4, we can get G from 8
		val4AE := addTogether(len4, segMap["a"])
		val4AE = addTogether(val4AE, segMap["e"])
		nineDiff := differenceIn(val4AE, len7)
		if len(nineDiff) != 1 {
			panic("Could not determin G")
		}
		segMap["g"] = nineDiff[0]

		// D is not the item left
		known := segMap["a"] + segMap["b"] + segMap["c"] + segMap["e"] + segMap["f"] + segMap["g"]
		segMap["d"] = differenceIn(known, "abcdefg")[0]

		// Build a list of the keysToValue
		rubbishValMap := make(map[string]int)
		for sc := 0; sc < 10; sc++ {
			// std number segments
			stdSeg := stdSegmentMap[sc]
			badSeg := SortString(getMappedSegments(stdSeg, segMap))
			rubbishValMap[badSeg] = sc
		}

		numberStr := ""
		// Now loop through the digits and work them out
		for y := 0; y < len(digits); y++ {
			sortedSegs := SortString(digits[y])
			numericVal, pres := rubbishValMap[sortedSegs]
			if !pres {
				panic("Oh noes")
			}
			numberStr = numberStr + fmt.Sprint(numericVal)
		}
		intNum, _ := strconv.Atoi(numberStr)
		tot += intNum

	}

	return tot
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func getMappedSegments(segmentKeys string, segmentKeyMap map[string]string) string {
	str := ""
	for x := 0; x < len(segmentKeys); x++ {
		str = str + segmentKeyMap[string(segmentKeys[x])]
	}
	return str
}

func getSegmentsWithLength(entries []string, segLength int) []string {
	returnList := []string{}
	for y := 0; y < len(entries); y++ {
		if len(entries[y]) == segLength {
			returnList = append(returnList, entries[y])
		}
	}
	if len(returnList) == 0 {
		panic("None found with length")
	}
	return returnList
}

func difference(string1 string, string2 string) []string {
	var diff []string

	atob := differenceIn(string1, string2)
	btoa := differenceIn(string2, string1)
	diff = append(diff, atob...)
	diff = append(diff, btoa...)

	return diff
}

func differenceIn(stringToEval string, searchString string) []string {
	slice1 := []rune(searchString)
	slice2 := []rune(stringToEval)

	var diff []string

	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, string(s1))
		}
	}

	return diff
}

func stringContains(string1 string, string2 string) bool {
	for x := 0; x < len(string2); x++ {
		found := false
		for y := 0; y < len(string1); y++ {
			if string2[x] == string1[y] {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func addTogether(string1 string, string2 string) string {
	returnString := string1
	// Add each value of string 2 not found in string 1
	for x := 0; x < len(string2); x++ {
		found := false
		for y := 0; y < len(string1); y++ {
			if string2[x] == string1[y] {
				found = true
			}
		}
		if !found {
			returnString = returnString + string(string2[x])
		}
	}
	return returnString
}
