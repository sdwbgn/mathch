package main

import (
	"testing"
)

func TestEngine(t *testing.T) {
	if !Solve("7+7=111", "7+7=14") {
		t.Errorf("Simple Solve Failed: 7+7=111 (Ans: 7+7=14)")
	}
	if !Solve("12-4=1", "12-11=1") {
		t.Errorf("Simple Solve Failed: 12-4=1 (Ans: 12-11=1)")
	}
	for i := 0; i < 10000; i++ {
		ex := RandGenerate(1)
		pp, _ := MoveMatch(ex, true)
		if len(pp) == 0 {
			t.Errorf("Couldn't Solve: %v", ex)
		}
	}
	for i := 0; i < 10000; i++ {
		ex := RandGenerate(2)
		pp, _ := MoveMatch(ex, true)
		if len(pp) == 0 {
			t.Errorf("Couldn't Solve: %v", ex)
		}
	}
	for i := 0; i < 10000; i++ {
		ex := RandGenerate(3)
		pp, _ := MoveMatch(ex, true)
		if len(pp) == 0 {
			t.Errorf("Couldn't Solve: %v", ex)
		}
	}

}
