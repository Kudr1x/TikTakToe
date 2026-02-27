package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type TicTacToeUI struct {
	game    *Game
	window  fyne.Window
	buttons [3][3][3][3]*widget.Button
	scoreX  *widget.Label
	scoreO  *widget.Label
	status  *widget.Label
}

func NewTicTacToeUI() *TicTacToeUI {
	a := app.New()
	w := a.NewWindow("Ultimate Tic-Tac-Toe")
	w.Resize(fyne.NewSize(500, 500))

	ui := &TicTacToeUI{
		game:   NewGame(),
		window: w,
	}

	ui.setupContent()
	return ui
}

func (ui *TicTacToeUI) setupContent() {
	topBar := ui.initTopBar()
	board := ui.initPlayBoard()
	ui.status = widget.NewLabel("Ход игрока X")

	content := container.NewBorder(
		container.NewVBox(topBar, ui.status),
		nil, nil, nil,
		board,
	)
	ui.window.SetContent(content)
}

func (ui *TicTacToeUI) initTopBar() *fyne.Container {
	ui.scoreX = widget.NewLabel("0")
	ui.scoreO = widget.NewLabel("0")

	return container.NewGridWithColumns(3,
		container.NewGridWithRows(2, widget.NewLabel("Игрок X"), ui.scoreX),
		container.NewGridWithRows(2,
			widget.NewButton("Сброс", ui.resetScores),
			widget.NewButton("Заново", ui.restartGame),
		),
		container.NewGridWithRows(2, widget.NewLabel("Игрок O"), ui.scoreO),
	)
}

func (ui *TicTacToeUI) initPlayBoard() *fyne.Container {
	bigGrid := container.NewGridWithColumns(3)

	for br := 0; br < 3; br++ {
		for bc := 0; bc < 3; bc++ {
			smallGrid := container.NewGridWithColumns(3)
			for sr := 0; sr < 3; sr++ {
				for sc := 0; sc < 3; sc++ {
					bRow, bCol, sRow, sCol := br, bc, sr, sc
					btn := widget.NewButton("", func() {
						ui.handleMove(bRow, bCol, sRow, sCol)
					})
					ui.buttons[br][bc][sr][sc] = btn
					smallGrid.Add(btn)
				}
			}

			card := widget.NewCard("", "", smallGrid)
			bigGrid.Add(card)
		}
	}

	return bigGrid
}

func (ui *TicTacToeUI) handleMove(br, bc, sr, sc int) {
	symbol, ok := ui.game.Play(br, bc, sr, sc)
	if !ok {
		return
	}

	ui.buttons[br][bc][sr][sc].SetText(symbol)
	ui.updateBoardState()
	ui.checkGameState()
}

func (ui *TicTacToeUI) updateBoardState() {
	player := ui.game.CurrentPlayer()
	ui.status.SetText(fmt.Sprintf("Ход игрока %s", player.String()))

	for br := 0; br < 3; br++ {
		for bc := 0; bc < 3; bc++ {
			board := ui.game.Boards[br][bc]
			isPossibleBoard := ui.game.NextBigRow == -1 || (br == ui.game.NextBigRow && bc == ui.game.NextBigCol)

			for sr := 0; sr < 3; sr++ {
				for sc := 0; sc < 3; sc++ {
					btn := ui.buttons[br][bc][sr][sc]

					switch {
					case board.Winner != None:
						btn.SetText(board.Winner.String())
						btn.Disable()
					case board.Cells[sr][sc] != None, !isPossibleBoard:
						btn.Disable()
					default:
						btn.Enable()
					}

					if isPossibleBoard && board.Winner == None && !ui.game.IsGameOver {
						btn.Importance = widget.HighImportance
					} else {
						btn.Importance = widget.MediumImportance
					}
					btn.Refresh()
				}
			}
		}
	}
}

func (ui *TicTacToeUI) checkGameState() {
	won, winner := ui.game.CheckBigWin()
	if !won && !ui.isDraw() {
		return
	}

	var msg string
	if won {
		if winner == PlayerX {
			ui.game.ScoreX++
			ui.scoreX.SetText(strconv.Itoa(ui.game.ScoreX))
			msg = "Победа игрока X во всей игре! "
		} else {
			ui.game.ScoreO++
			ui.scoreO.SetText(strconv.Itoa(ui.game.ScoreO))
			msg = "Победа игрока O во всей игре! "
		}
	} else {
		msg = "Ничья в игре! "
	}

	dialog.ShowConfirm("Игра окончена", msg+"Начать заново?", func(b bool) {
		if b {
			ui.restartGame()
		}
	}, ui.window)
}

func (ui *TicTacToeUI) isDraw() bool {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if ui.game.Boards[r][c].Winner == None && !ui.game.Boards[r][c].IsFull {
				return false
			}
		}
	}
	return true
}

func (ui *TicTacToeUI) restartGame() {
	ui.game.Restart()
	for br := 0; br < 3; br++ {
		for bc := 0; bc < 3; bc++ {
			for sr := 0; sr < 3; sr++ {
				for sc := 0; sc < 3; sc++ {
					btn := ui.buttons[br][bc][sr][sc]
					btn.SetText("")
					btn.Enable()
					btn.Importance = widget.MediumImportance
				}
			}
		}
	}
	ui.status.SetText("Ход игрока X")
}

func (ui *TicTacToeUI) resetScores() {
	ui.game.ResetScore()
	ui.scoreX.SetText("0")
	ui.scoreO.SetText("0")
}

func (ui *TicTacToeUI) Run() {
	ui.window.ShowAndRun()
}
