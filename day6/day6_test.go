package main

import (
	"testing"
)

func TestMainFunctionRuns(t *testing.T) {
	main()
}

func TestExampleGiven(t *testing.T) {
	exp18Days := 26
	exp80Days := 5934
	exp256Days := 26984457539

	input := "3,4,3,1,2"

	res18Days := getFishCount(input, 18)
	res80Days := getFishCount(input, 80)
	res256Days := getFishCount(input, 256)

	if res18Days != exp18Days {
		t.Errorf("Example output for 18 day fish grown expected %v, got %v\n", exp18Days, res18Days)
	}

	if res80Days != exp80Days {
		t.Errorf("Example output for 80 day fish grown expected %v, got %v\n", exp80Days, res80Days)
	}

	if res256Days != exp256Days {
		t.Errorf("Example output for 256 day fish grown expected %v, got %v\n", exp256Days, res256Days)
	}
}
