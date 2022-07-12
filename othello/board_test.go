package othello

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitBoard(t *testing.T) {
	expected := [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 1, 0, 0, 0},
		{0, 0, 0, 1, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	actual := initBoard()
	assert.Equal(t, expected, actual)
}

func printBoard(board [8][8]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%d", board[i][j])
		}
		fmt.Println()
	}
}

func printAdjacencyMap(adjacencyMap [8][8]bool) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%t", adjacencyMap[i][j])
		}
		fmt.Println()
	}
}
