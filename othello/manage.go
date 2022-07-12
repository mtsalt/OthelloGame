package othello

type GameInfo struct {
	battleNum     int
	boardInfo     *BoardInfo // ゲーム回数で、盤面リストを選びだそうとしたが不要
	numOfWinWhite int
	numOfWinBlack int
	numOfDraw     int
}

func NewMatch() *GameInfo {
	g := new(GameInfo)
	g.boardInfo = NewBoard()
	return g
}

func (g *GameInfo) NextMatch() {
	g.incrementBattleNum()
	g.boardInfo = NewBoard()
	g.incrementWinNum()
}

func (g *GameInfo) UpdateBoard(x int, y int) error {
	// NewGameと、Updateのみでゲームを実現する
	// 何が起きているかは、ルーティング処理には見せない
	// 例外発生可能性はあるので、これをルーティング側に知らせることができる仕組みが必要
	// ゲームを始める、盤面データを記録する（将来の話し）、ゲームを終了する（この構造体をメモリから削除する）という処理は
	// このファイルよりも上の処理で実施する

	// error を返せば異常状態なので、処理終了にする
	// 最終的には、以下のようになるのが理想

	b := g.boardInfo
	err := b.updateBoardInfo(x, y)
	if err != nil {
		return err
	}

	return nil
}

func (g *GameInfo) Pass() {
	b := g.boardInfo
	b.incrementPassCount()
}

func (g *GameInfo) GameIsOver() bool {
	return g.boardInfo.gameIsOver()
}

func (g *GameInfo) incrementWinNum() {
	w := g.boardInfo.winner()
	switch w {
	case BLACK:
		g.numOfWinBlack++
	case WHITE:
		g.numOfWinWhite++
	case EMPTY:
		g.numOfDraw++
	}
}

func (g *GameInfo) incrementBattleNum() {
	g.battleNum++
}
