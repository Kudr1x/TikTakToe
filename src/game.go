package main

type Player uint

const (
	None Player = iota
	PlayerX
	PlayerO
)

type Game struct {
	Board         [3][3]Player
	IsPlayerOTurn bool
	LockBoard     bool
	WinState      bool
	ScoreX        int
	ScoreO        int
}

func NewGame() *Game {
	return &Game{
		Board: [3][3]Player{},
	}
}

func (g *Game) Restart() {
	g.Board = [3][3]Player{}
	g.IsPlayerOTurn = false
	g.LockBoard = false
	g.WinState = false
}

func (g *Game) ResetScore() {
	g.ScoreX = 0
	g.ScoreO = 0
}

func (g *Game) Play(x, y int) (string, bool) {
	if g.LockBoard || g.Board[x][y] != None {
		return "", false
	}

	symbol := "X"
	if g.IsPlayerOTurn {
		g.Board[x][y] = PlayerO
		symbol = "O"
	} else {
		g.Board[x][y] = PlayerX
	}

	g.IsPlayerOTurn = !g.IsPlayerOTurn
	return symbol, true
}

func (g *Game) CheckWin() (hasWinner bool, winner Player, isDraw bool) {
	// Check columns
	for i := 0; i < 3; i++ {
		if g.Board[i][0] != None && g.Board[i][0] == g.Board[i][1] && g.Board[i][0] == g.Board[i][2] {
			return true, g.Board[i][0], false
		}
	}
	// Check rows
	for i := 0; i < 3; i++ {
		if g.Board[0][i] != None && g.Board[0][i] == g.Board[1][i] && g.Board[0][i] == g.Board[2][i] {
			return true, g.Board[0][i], false
		}
	}
	// Check diagonals
	if g.Board[1][1] != None {
		if (g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2]) ||
			(g.Board[0][2] == g.Board[1][1] && g.Board[1][1] == g.Board[2][0]) {
			return true, g.Board[1][1], false
		}
	}

	// Check draw
	full := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == None {
				full = false
				break
			}
		}
	}

	if full {
		return false, None, true
	}

	return false, None, false
}
