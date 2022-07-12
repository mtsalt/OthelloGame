package othello

import "errors"

type BoardInfo struct {
	board           [8][8]int
	emptyNum        int
	whiteNum        int
	blackNum        int
	turn            int
	passCount       int
	availablePoints pointList
}

func NewBoard() *BoardInfo {
	b := new(BoardInfo)
	b.board = initBoard()
	b.turn = BLACK
	b.countStone()
	b.computeCanPutPoints()
	return b
}

func (b *BoardInfo) updateBoardInfo(x int, y int) error {
	err := b.putStone(point{x, y}, b.turn)
	if err != nil {
		return err
	}
	b.countStone()
	b.switchTurn()
	b.computeCanPutPoints()
	return nil
}

func initBoard() [8][8]int {
	var b [8][8]int
	b[3][3] = BLACK
	b[3][4] = WHITE
	b[4][3] = WHITE
	b[4][4] = BLACK
	return b
}

func (b *BoardInfo) putStone(p point, colorNum int) error {
	// 盤面更新
	// クライアントから送られてきた座標が、配置可能座標一覧に入っているか確認
	// 最終的に、このメソッドでErrorを返せるようにし、座標に問題ありならゲーム終了にする
	if b.pointInAvailablePoints(p) {
		b.setAndReverse(p)
		return nil
	} else {
		return errors.New("point is invalid.")
	}
}

func (b *BoardInfo) switchTurn() {
	if b.turn == BLACK {
		b.turn = WHITE
	} else if b.turn == WHITE {
		b.turn = BLACK
	} else {
		// error
	}
}

func (b *BoardInfo) pointInAvailablePoints(p point) bool {
	return b.availablePoints.Contain(p)
}

func (b *BoardInfo) setAndReverse(p point) {

}

func (b *BoardInfo) computeCanPutPoints() {
	b.availablePoints = pointList{}
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if b.canPutStone(x, y) {
				b.availablePoints.Add(point{x, y})
			}
		}
	}
}

func (b *BoardInfo) canPutStone(x int, y int) bool {

	board := b.board
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == EMPTY && b.isAvailable(i, j) {
				return true
			}
		}
	}
	return false
}

func (b *BoardInfo) isAvailable(x int, y int) bool {

	offset := [3]int{-1, 0, 1}
	for yOffset := range offset {
		for xOffset := range offset {
			if b.canReverse(x, y, xOffset, yOffset) {
				return true
			}
		}
	}
	return false
}

func pointIsOnBoard(x int, y int) bool {
	if (0 <= x && x < 8) && (0 <= y && y < 8) {
		return true
	} else {
		return false
	}
}

func (b *BoardInfo) canReverse(x int, y int, xOffset int, yOffset int) bool {

	for n := 0; ; n++ {
		xPoint, yPoint := x+xOffset, y+yOffset
		if !pointIsOnBoard(xPoint, yPoint) {
			return false
		}
	}
}

func (b *BoardInfo) countStone() (int, int, int) {
	var empty, white, black int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			switch b.board[i][j] {
			case EMPTY:
				empty++
			case WHITE:
				white++
			case BLACK:
				black++
			}
		}
	}
	b.emptyNum = empty
	b.whiteNum = white
	b.blackNum = black
	return empty, white, black
}

func (b *BoardInfo) incrementPassCount() {
	b.passCount++
}

func (b *BoardInfo) gameIsOver() bool {

	// 白黒両方の、available positionが空になれば終了
	if b.passCount > 2 {
		// passCountが２になったら終了 → 置く場所がない
		return true
	}
	return false
}

func (b *BoardInfo) winner() int {
	if b.blackNum < b.whiteNum {
		return WHITE
	} else if b.whiteNum < b.blackNum {
		return BLACK
	} else {
		return EMPTY
	}
}
