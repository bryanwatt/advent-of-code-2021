package main

import (
	//"adventofcode/shared"
	"testing"
)

func TestMainFunctionRuns(t *testing.T) {
	main()
}

func TestTotaldisplayOutput(t *testing.T) {
	input := []string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"}

	res := totaldisplayOutput(input)

	if res != 5355 {
		t.Error("fail")
	}
}
