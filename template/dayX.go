package main

import (
	"adventofcode/shared"
	"fmt"
)

func main() {
	bla, _ := shared.ReadFileLineToStrArr("puzzleOrExample.txt")
	fmt.Print("This is empty class thats wired up for these examples")

	if bla != nil {
		fmt.Print("Bla Bla")
	}

}
