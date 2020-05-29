package mathch

import (
	"errors"
)

func test() {
	if !Solve("7+7=111", "7+7=14") {
		panic(errors.New("test 1 failed"))
	}
	for i := 0; i < 10000; i++ {
		ex := RandGenerate(1)
		pp, _ := MoveMatch(ex, true)
		if len(pp) == 0 {
			panic(errors.New("unsolvable"))
		}
	}
	for i := 0; i < 10000; i++ {
		ex := RandGenerate(2)
		pp, _ := MoveMatch(ex, true)
		if len(pp) == 0 {
			panic(errors.New("unsolvable"))
		}
	}
	for i := 0; i < 10000; i++ {
		ex := RandGenerate(3)
		pp, _ := MoveMatch(ex, true)
		if len(pp) == 0 {
			panic(errors.New("unsolvable"))
		}
	}

}
