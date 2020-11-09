package game

import (
	"testing"
)

func TestLine001(t *testing.T) {
	board := []int{1, 0, 0, 0}
	result := line(board, 1, -1)
	result2:=line(board,2,-1)

	if !result {
		t.Errorf("%v %v", board,result2)
	}

}



func TestCheckPossition01(t *testing.T) {
	//checkPossition(board []int, pos int, value int) bool
	board := []int{0, 0, 0}
	pos := 0
	value := 1
	wait := true
	res := checkPossition(board, pos, value)
	if res != wait {
		t.Errorf("On Board %v in pos %v can't be in value %v", board, pos, value)
	}
} // func
func TestCheckPossition02(t *testing.T) {
	//checkPossition(board []int, pos int, value int) bool
	board := []int{1, 0, 0}
	pos := 1
	value := 1
	wait := false
	res := checkPossition(board, pos, value)
	if res != wait {
		t.Errorf("On Board %v in pos %v can't be in value %v", board, pos, value)
	}
} // func

func TestCheckPossition03(t *testing.T) {
	//checkPossition(board []int, pos int, value int) bool
	board := []int{2, 4, 1, 0}
	pos := 4
	value := 3
	wait := true
	res := checkPossition(board, pos, value)
	if res != wait {
		t.Errorf("On Board %v in pos %v can't be in value %v", board, pos, value)
	}
} // func


func TestCompareSlices(t *testing.T) {
	in := []int{0, 0}
	out := []int{0, 0}
	resBool, resInt := compareSlices(in, out)

	if !resBool {
		t.Errorf("In %v not = wiht Out %v on index %v", in, out, resInt)
	}

}

// Test canceled
// --------------------------------------------------------------------------
// func TestPreparePoolMulty(t *testing.T) {
// 	cases := [][][]int{
// 		{{0, 0, 0}, {1, 2, 3}},
// 		{{0, 0}, {1, 2}},
// 		{{0, 0, 0, 0}, {1, 2, 3, 4}},
// 	}
// 	for counter := 0; counter < len(cases); counter++ {
// 		in := cases[counter][0]
// 		out := cases[counter][1]
// 		res := make([]int, len(in))
// 		prePool := preparePool(in, res)
// 		resBool, resInt := compareSlices(out, prePool)
// 		if !resBool {
// 			t.Errorf("In %v not = wiht Out %v on index %v", in, out, resInt)
// 		}
// 	}
// --------------------------------------------------------------------------
// func TestPreparePool(t *testing.T) {
// 	in := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	out := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	res := make([]int, len(in))
// 	prePool := preparePool(in, res)
// 	resBool, resInt := compareSlices(out, prePool)
// 	if !resBool {
// 		t.Errorf("In %v not = wiht Out %v on index %v", in, out, resInt)
// 	}
// }
// --------------------------------------------------------------------------
