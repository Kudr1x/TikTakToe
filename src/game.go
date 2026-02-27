package main

type Player uint

const (
	None Player = iota
	PlayerX
	PlayerO
)

func (p Player) String() string {
	switch p {
	case PlayerX:
		return "X"
	case PlayerO:
		return "O"
	default:
		return ""
	}
}

type Game struct {
	Board         [3][3]Player
	IsPlayerOTurn bool
	LockBoard     bool
	ScoreX        int
	ScoreO        int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Restart() {
	g.Board = [3][3]Player{}
	g.IsPlayerOTurn = false
	g.LockBoard = false
}

func (g *Game) ResetScore() {
	g.ScoreX = 0
	g.ScoreO = 0
}

func (g *Game) CurrentPlayer() Player {
	if g.IsPlayerOTurn {
		return PlayerO
	}
	return PlayerX
}

func (g *Game) Play(row, col int) (symbol string, ok bool) {
	if g.LockBoard || row < 0 || row > 2 || col < 0 || col > 2 || g.Board[row][col] != None {
		return "", false
	}

	currentPlayer := g.CurrentPlayer()
	g.Board[row][col] = currentPlayer
	g.IsPlayerOTurn = !g.IsPlayerOTurn

	return currentPlayer.String(), true
}

func (g *Game) CheckWin() (hasWinner bool, winner Player, isDraw bool) {
	lines := [8][3][2]int{
		// Rows
		{{0, 0}, {0, 1}, {0, 2}}, {{1, 0}, {1, 1}, {1, 2}}, {{2, 0}, {2, 1}, {2, 2}},
		// Cols
		{{0, 0}, {1, 0}, {2, 0}}, {{0, 1}, {1, 1}, {2, 1}}, {{0, 2}, {1, 2}, {2, 2}},
		// Diagonals
		{{0, 0}, {1, 1}, {2, 2}}, {{0, 2}, {1, 1}, {2, 0}},
	}

	for _, line := range lines {
		p1 := g.Board[line[0][0]][line[0][1]]
		p2 := g.Board[line[1][0]][line[1][1]]
		p3 := g.Board[line[2][0]][line[2][1]]

		if p1 != None && p1 == p2 && p2 == p3 {
			return true, p1, false
		}
	}

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if g.Board[r][c] == None {
				return false, None, false
			}
		}
	}

	return false, None, true
}
