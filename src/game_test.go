package main

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	g := NewGame()
	if g.CurrentPlayer() != PlayerX {
		t.Error("Expected Player X to start first")
	}
}

func TestUltimatePlay(t *testing.T) {
	g := NewGame()

	symbol, ok := g.Play(1, 1, 1, 1)
	if !ok || symbol != "X" {
		t.Errorf("Expected valid move for X, got %v, %s", ok, symbol)
	}

	if g.NextBigRow != 1 || g.NextBigCol != 1 {
		t.Errorf("Expected NextBig to be [1,1], got [%d,%d]", g.NextBigRow, g.NextBigCol)
	}

	_, ok = g.Play(0, 0, 0, 0)
	if ok {
		t.Error("Should not allow move outside of the target big board")
	}

	symbol, ok = g.Play(1, 1, 0, 0)
	if !ok || symbol != "O" {
		t.Errorf("Expected valid move for O in [1,1][0,0], got %v, %s", ok, symbol)
	}

	if g.NextBigRow != 0 || g.NextBigCol != 0 {
		t.Errorf("Expected NextBig to be [0,0], got [%d,%d]", g.NextBigRow, g.NextBigCol)
	}
}

func TestSmallBoardWin(t *testing.T) {
	g := NewGame()

	g.Boards[0][0].Cells[0][0] = PlayerX
	g.Boards[0][0].Cells[0][1] = PlayerX
	g.Boards[0][0].Cells[0][2] = PlayerX

	won, winner := checkLine(g.Boards[0][0].Cells)
	if !won || winner != PlayerX {
		t.Errorf("Small board [0,0] should be won by X")
	}
}

func TestBigWin(t *testing.T) {
	g := NewGame()

	g.Boards[0][0].Winner = PlayerX
	g.Boards[0][1].Winner = PlayerX
	g.Boards[0][2].Winner = PlayerX

	won, winner := g.CheckBigWin()
	if !won || winner != PlayerX {
		t.Errorf("Big win expected for X")
	}
}
