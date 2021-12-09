package main

import (
	"adventofcode/shared"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	commandList, err := shared.ReadFileLineToStrArr("day2_puzzle.txt")

	if err != nil {
		panic(err)
	}

	part1(commandList)
	part2(commandList)
}

func part1(commandList []string) {
	horiz, depth := calculateSubPosition(commandList)

	fmt.Println("Part 1")
	fmt.Printf("Sub Horizontal Position: %v\n", horiz)
	fmt.Printf("Sub Depth: %v\n", depth)
	fmt.Printf("Total multiplied: %v\n", horiz*depth)
}

func part2(commandList []string) {
	horiz, depth, aim := calculateSubPositionWithAim(commandList)

	fmt.Println("Part 2")
	fmt.Printf("Sub Horizontal Position: %v\n", horiz)
	fmt.Printf("Sub Depth: %v\n", depth)
	fmt.Printf("Sub Aim: %v\n", aim)
	fmt.Printf("Total multiplied: %v\n", horiz*depth)
}

func calculateSubPosition(commandList []string) (int, int) {
	horiz := 0
	depth := 0

	for _, command := range commandList {
		h, v := calculateCommandResult(command)
		horiz += h
		depth += v
	}

	return horiz, depth
}

// Starting at 0 0, apply all command effects and calculate the resulting end horizontal and depth positions
func calculateSubPositionWithAim(commandList []string) (int, int, int) {
	horiz := 0
	depth := 0
	aim := 0

	for _, command := range commandList {
		h, v := calculateCommandResult(command)
		horiz += h
		aim += v

		depth += h * aim
	}

	return horiz, depth, aim
}

// Note, vert effect is inverse - INCREASES as it goes down
func calculateCommandResult(command string) (int, int) {
	horizontalEffect := 0
	verticalEffect := 0

	s := strings.Split(command, " ")

	valEffect, err := strconv.Atoi(s[1])
	if err != nil {
		panic(fmt.Errorf("could not convert value %v to an int", s[1]))
	}

	switch s[0] {
	case "forward":
		horizontalEffect = valEffect
	case "up":
		verticalEffect = -valEffect
	case "down":
		verticalEffect = valEffect
	default:
		panic(fmt.Errorf("no idea what %v is", s[0]))
	}

	return horizontalEffect, verticalEffect
}
