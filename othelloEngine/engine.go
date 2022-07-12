package othello

import "fmt"

type OutcomeInfo struct {
	times       int
	whiteWinNum int
	blackWinNum int
	board       [8][8]int
}

func NextBoardState(board [8][8]int) [8][8]int {
	fmt.Println(board)
	return board
}

func EngineMain() {

}
