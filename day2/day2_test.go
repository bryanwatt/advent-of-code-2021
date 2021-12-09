package main

import (
	"adventofcode/shared"
	"testing"
)

func TestCalculateSubPosition_for_example_provided(t *testing.T) {
	// This is part 1's test testing code arrives at example conclusion
	intputValues, err := shared.ReadFileLineToStrArr("day2_example.txt")

	if err != nil {
		t.Error(err)
	}

	horizPos, depth := calculateSubPosition(intputValues)

	if horizPos != 15 {
		t.Error("Horizontal Position in example should be 15")
	}
	if depth != 10 {
		t.Error("Depth in example should be 10")
	}
}

func TestCalculateSubPositionWithAimfor_example_provided(t *testing.T) {
	// This is part 2's test testing code arrives at example conclusion
	intputValues, err := shared.ReadFileLineToStrArr("day2_example.txt")

	if err != nil {
		t.Error(err)
	}

	horizPos, depth, aim := calculateSubPositionWithAim(intputValues)

	if horizPos != 15 {
		t.Errorf("Horizontal Position in example should be 15, got %v", horizPos)
	}
	if depth != 60 {
		t.Errorf("Depth in example should be 60, got %v", depth)
	}
	if aim != 10 {
		t.Errorf("Depth in example should be 10, got %v", aim)
	}
}

func TestCalculateCommandResult_for_up_command(t *testing.T) {
	command := "up 3"

	horEffect, depthEffect := calculateCommandResult(command)

	if horEffect != 0 {
		t.Error("up command should not effect vertical position")
	}

	if depthEffect != -3 {
		t.Error("up 3 command should decrease the depth position by 3")
	}
}

func TestCalculateCommandResult_for_down_command(t *testing.T) {
	command := "down 5"

	horEffect, depthEffect := calculateCommandResult(command)

	if horEffect != 0 {
		t.Error("down command should not effect vertical position")
	}

	if depthEffect != 5 {
		t.Error("down 5 command should increase the depth position by 5")
	}
}

func TestCalculateCommandResult_for_forward_command(t *testing.T) {
	command := "forward 2"

	horEffect, depthEffect := calculateCommandResult(command)

	if depthEffect != 0 {
		t.Error("forward command should not effect depth position")
	}

	if horEffect != 2 {
		t.Error("forward 2 command should increase the horizontal position by 2")
	}
}
