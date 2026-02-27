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

type SmallBoard struct {
	Cells  [3][3]Player
	Winner Player
	IsFull bool
}

type Game struct {
	Boards        [3][3]SmallBoard
	BigWinner     Player
	IsPlayerOTurn bool
	NextBigRow    int // -1 означает свободный ход в любом квадрате
	NextBigCol    int // -1 означает свободный ход в любом квадрате
	ScoreX        int
	ScoreO        int
	IsGameOver    bool
}

func NewGame() *Game {
	g := &Game{}
	g.Restart()
	return g
}

func (g *Game) Restart() {
	g.Boards = [3][3]SmallBoard{}
	g.BigWinner = None
	g.IsPlayerOTurn = false
	g.NextBigRow = -1
	g.NextBigCol = -1
	g.IsGameOver = false
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

func (g *Game) Play(bigRow, bigCol, smallRow, smallCol int) (symbol string, ok bool) {
	if g.IsGameOver {
		return "", false
	}

	if g.NextBigRow != -1 && (bigRow != g.NextBigRow || bigCol != g.NextBigCol) {
		return "", false
	}

	targetBoard := &g.Boards[bigRow][bigCol]

	if targetBoard.Winner != None || targetBoard.Cells[smallRow][smallCol] != None {
		return "", false
	}

	currentPlayer := g.CurrentPlayer()
	targetBoard.Cells[smallRow][smallCol] = currentPlayer
	symbol = currentPlayer.String()

	if won, winner := checkLine(targetBoard.Cells); won {
		targetBoard.Winner = winner
	} else if isBoardFull(targetBoard.Cells) {
		targetBoard.IsFull = true
	}

	if won, winner := g.CheckBigWin(); won {
		g.BigWinner = winner
		g.IsGameOver = true
	}

	nextBoard := &g.Boards[smallRow][smallCol]
	if nextBoard.Winner != None || nextBoard.IsFull {
		g.NextBigRow = -1
		g.NextBigCol = -1
	} else {
		g.NextBigRow = smallRow
		g.NextBigCol = smallCol
	}

	g.IsPlayerOTurn = !g.IsPlayerOTurn
	return symbol, true
}

func checkLine(board [3][3]Player) (bool, Player) {
	lines := [8][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}}, {{1, 0}, {1, 1}, {1, 2}}, {{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}}, {{0, 1}, {1, 1}, {2, 1}}, {{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}}, {{0, 2}, {1, 1}, {2, 0}},
	}

	for _, line := range lines {
		p1 := board[line[0][0]][line[0][1]]
		p2 := board[line[1][0]][line[1][1]]
		p3 := board[line[2][0]][line[2][1]]

		if p1 != None && p1 == p2 && p2 == p3 {
			return true, p1
		}
	}
	return false, None
}

func isBoardFull(board [3][3]Player) bool {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[r][c] == None {
				return false
			}
		}
	}
	return true
}

func (g *Game) CheckBigWin() (bool, Player) {
	var bigState [3][3]Player
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			bigState[r][c] = g.Boards[r][c].Winner
		}
	}
	return checkLine(bigState)
}
