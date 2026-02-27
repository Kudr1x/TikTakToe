package main

import (
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
	buttons [9]*widget.Button
	scoreX  *widget.Label
	scoreO  *widget.Label
}

func NewTicTacToeUI() *TicTacToeUI {
	a := app.New()
	w := a.NewWindow("TikTakToe")
	w.Resize(fyne.NewSize(300, 300))

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

	content := container.NewGridWithRows(4, topBar, board[0], board[1], board[2])
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

func (ui *TicTacToeUI) initPlayBoard() [3]*fyne.Container {
	for i := 0; i < 9; i++ {
		idx := i
		row, col := i/3, i%3
		ui.buttons[idx] = widget.NewButton("", func() {
			ui.handleMove(idx, row, col)
		})
	}

	return [3]*fyne.Container{
		container.NewGridWithColumns(3, ui.buttons[0], ui.buttons[1], ui.buttons[2]),
		container.NewGridWithColumns(3, ui.buttons[3], ui.buttons[4], ui.buttons[5]),
		container.NewGridWithColumns(3, ui.buttons[6], ui.buttons[7], ui.buttons[8]),
	}
}

func (ui *TicTacToeUI) handleMove(idx, row, col int) {
	symbol, ok := ui.game.Play(row, col)
	if !ok {
		return
	}

	ui.buttons[idx].SetText(symbol)
	ui.checkGameState()
}

func (ui *TicTacToeUI) checkGameState() {
	hasWinner, winner, isDraw := ui.game.CheckWin()
	if !hasWinner && !isDraw {
		return
	}

	var msg string
	if hasWinner {
		if winner == PlayerX {
			ui.game.ScoreX++
			ui.scoreX.SetText(strconv.Itoa(ui.game.ScoreX))
			msg = "Победа игрока X. "
		} else {
			ui.game.ScoreO++
			ui.scoreO.SetText(strconv.Itoa(ui.game.ScoreO))
			msg = "Победа игрока O. "
		}
	} else {
		msg = "Ничья. "
	}

	dialog.ShowConfirm("Игра окончена", msg+"Начать заново?", func(b bool) {
		if b {
			ui.restartGame()
		} else {
			ui.game.LockBoard = true
		}
	}, ui.window)
}

func (ui *TicTacToeUI) restartGame() {
	ui.game.Restart()
	for _, btn := range ui.buttons {
		btn.SetText("")
	}
}

func (ui *TicTacToeUI) resetScores() {
	ui.game.ResetScore()
	ui.scoreX.SetText("0")
	ui.scoreO.SetText("0")
}

func (ui *TicTacToeUI) Run() {
	ui.window.ShowAndRun()
}
