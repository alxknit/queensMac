package game

import (
	"fmt"
	"strings"
	"time"
)

// Game ...
// this is the main function on this package
func Game(i int, short bool) {
	shortDisplay := short
	board := make([]int, i)
	start := time.Now()
	solve(board, shortDisplay)
	elapsed := time.Since(start)
	fmt.Println()
	fmt.Printf("Time elapsed %s\n", elapsed)
	// display(board)
}

func solve(board []int, shortDisplay bool) {
	// Variables
	size := len(board)
	solutions := 0
	pos := 0
	queens := 0
	fromValue := 1
	var res bool

	for {
		res = line(board, pos, fromValue)
		if res {
			queens++
			if queens == size {
				solutions++
				fmt.Println()
				fmt.Printf("Solution no %v\n", solutions)
				fmt.Println(board)
				if !shortDisplay {
					display(board)
				}

				board[pos] = 0
				queens = queens - 2
				pos--
				fromValue = board[pos]
				fromValue++
				board[pos] = 0
				continue
			} else {
				fromValue = 1
				pos++
			} // queens == size

		} // if res
		if !res {
			// backtrack
			pos--
			if pos == -1 {
				if solutions == 0 {
					fmt.Println("No solutions found")
				}
				break
			} // pos == -1 => no solution
			fromValue = board[pos]
			fromValue++
			board[pos] = 0
			queens--
		} // !res
	} // for ever
	fmt.Printf("Number of solutions = %v\n", solutions)

} // func solve

func line(board []int, pos int, fromValue int) bool {
	maxValue := len(board) + 1
	if fromValue == maxValue {
		return false
	}
	start := fromValue
	for count := start; count < maxValue; count++ {
		check := checkPossition(board, pos, count)
		// fmt.Println(count, check)
		if check {
			board[pos] = count
			return true
		}
	}
	return false
} // func

func display(board []int) {
	// display the array in a chess way
	// Get:
	//	board = array
	// Out:
	// 	nothing real...
	var str strings.Builder
	size := len(board)
	for row := 0; row < size; row++ {
		for column := 0; column < size; column++ {
			if row+1 == board[column] {
				// fmt.Print("Q ")
				fmt.Fprintf(&str, "Q ")
			} else {
				//fmt.Print(". ")
				fmt.Fprintf(&str, ". ")
			} //if
		} // for column
		//fmt.Println()
		fmt.Fprintf(&str, "\n")
	} // for row
	fmt.Println(str.String())
}

func checkPossition(board []int, pos int, value int) bool {
	/*
		Get's
			the board as referance
			pos of where we need to search for
			value to search for
		Return's
			bool if valid

	*/
	countMinus := 0
	countPlus := 0
	if 0 == pos {
		return true
	}
	for x := pos - 1; x > -1; x-- {
		countMinus++
		countPlus++
		valueOnBoard := board[x]
		if valueOnBoard == value {
			return false
		}
		if value+countPlus == valueOnBoard {
			return false
		}
		if value-countMinus == valueOnBoard {
			return false
		}
	} // for
	return true
}
