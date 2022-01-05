package shared

import (
	"strconv"
	"strings"
)

func StringToIntArr(input string) []int {
	numsStr := strings.Split(input, ",")
	numArray := []int{}
	for i := 0; i < len(numsStr); i++ {
		intVal, _ := strconv.Atoi(numsStr[i])
		numArray = append(numArray, intVal)
	}
	return numArray
}
