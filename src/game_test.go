package main

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	g := NewGame()
	if g.CurrentPlayer() != PlayerX {
		t.Error("Expected Player X to start first")
	}
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if g.Board[r][c] != None {
				t.Errorf("Expected empty board at [%d][%d]", r, c)
			}
		}
	}
}

func TestPlay(t *testing.T) {
	g := NewGame()

	// Test valid move
	symbol, ok := g.Play(0, 0)
	if !ok || symbol != "X" {
		t.Errorf("Expected valid move for X, got %v, %s", ok, symbol)
	}
	if g.Board[0][0] != PlayerX {
		t.Error("Board at [0][0] should be PlayerX")
	}

	// Test move to occupied cell
	symbol, ok = g.Play(0, 0)
	if ok || symbol != "" {
		t.Error("Should not allow move to occupied cell")
	}

	// Test next turn
	symbol, ok = g.Play(1, 1)
	if !ok || symbol != "O" {
		t.Errorf("Expected valid move for O, got %v, %s", ok, symbol)
	}
	if g.Board[1][1] != PlayerO {
		t.Error("Board at [1][1] should be PlayerO")
	}
}

func TestCheckWin(t *testing.T) {
	tests := []struct {
		name      string
		moves     [][2]int // sequence of moves (row, col)
		hasWinner bool
		winner    Player
		isDraw    bool
	}{
		{
			name:      "Row Win (First)",
			moves:     [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}},
			hasWinner: true,
			winner:    PlayerX,
		},
		{
			name:      "Column Win (First)",
			moves:     [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}},
			hasWinner: true,
			winner:    PlayerX,
		},
		{
			name:      "Diagonal Win",
			moves:     [][2]int{{0, 0}, {0, 1}, {1, 1}, {0, 2}, {2, 2}},
			hasWinner: true,
			winner:    PlayerX,
		},
		{
			name: "Draw",
			moves: [][2]int{
				{0, 0}, {0, 1}, {0, 2},
				{1, 1}, {1, 0}, {1, 2},
				{2, 1}, {2, 0}, {2, 2},
			},
			hasWinner: false,
			isDraw:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGame()
			for _, move := range tt.moves {
				g.Play(move[0], move[1])
			}
			hasWinner, winner, isDraw := g.CheckWin()
			if hasWinner != tt.hasWinner {
				t.Errorf("hasWinner: expected %v, got %v", tt.hasWinner, hasWinner)
			}
			if winner != tt.winner {
				t.Errorf("winner: expected %v, got %v", tt.winner, winner)
			}
			if isDraw != tt.isDraw {
				t.Errorf("isDraw: expected %v, got %v", tt.isDraw, isDraw)
			}
		})
	}
}
