package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var buttons [9]*widget.Button
var state = [3][3]uint{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
var lockBoard, winState, flag bool
var counterWin1Player, counterWin2Player int
var label1Player, label2Player *widget.Label

func main() {
	a := app.New()
	w := a.NewWindow("TikTakToe")
	w.Resize(fyne.NewSize(300, 300))

	line0 := initTopBarMenu()
	line1, line2, line3 := initPlayBoard(w)

	content := container.NewGridWithRows(4, line0, line1, line2, line3)

	w.SetContent(content)
	w.ShowAndRun()
}

func initPlayBoard(window fyne.Window) (*fyne.Container, *fyne.Container, *fyne.Container) {
	for i := 0; i < 9; i++ {
		x, y := getPositionInMatrix(i)
		var button *widget.Button
		button = widget.NewButton("", func() {
			checkState, sta := changeState(x, y)
			if checkState && !lockBoard {
				button.SetText(sta)
				check(window)
			}
			//fmt.Println(x, y)
		})
		buttons[i] = button
	}

	line1 := container.NewGridWithColumns(3, buttons[0], buttons[1], buttons[2])
	line2 := container.NewGridWithColumns(3, buttons[3], buttons[4], buttons[5])
	line3 := container.NewGridWithColumns(3, buttons[6], buttons[7], buttons[8])

	return line1, line2, line3
}

func initTopBarMenu() *fyne.Container {
	label1Player = widget.NewLabel(strconv.Itoa(counterWin1Player))
	label2Player = widget.NewLabel(strconv.Itoa(counterWin2Player))

	line0 := container.NewGridWithColumns(3, container.NewGridWithRows(2, widget.NewLabel("1 Игрок"), label1Player),
		container.NewGridWithRows(2, widget.NewButton("Сброс", func() {
			counterWin1Player = 0
			counterWin2Player = 0
			label1Player.SetText("0")
			label2Player.SetText("0")
		}), widget.NewButton("Заново", func() {
			restart()
		})), container.NewGridWithRows(2, widget.NewLabel("2 Игрок"), label2Player))

	return line0
}

func getPositionInMatrix(pos int) (x int, y int) {
	y = pos / 3
	x = pos % 3

	return x, y
}

func changeState(x int, y int) (checkState bool, st string) {
	if state[x][y] != 0 {
		return false, ""
	}
	if flag {
		flag = !flag
		state[x][y] = 1
		return true, "O"
	} else {
		flag = !flag
		state[x][y] = 2
		return true, "X"
	}
}

func check(window fyne.Window) {
	checkColumns(window)
	checkRows(window)
	checkDiagonals(window)
	checkAllDesk(window)
}

func checkColumns(window fyne.Window) {
	for i := 0; i < 3; i++ {
		if state[i][0] == state[i][1] && state[i][0] == state[i][2] && state[i][0] != 0 && state[i][1] != 0 && state[i][2] != 0 {
			endGame(true, state[0][i] == 1, window)
		}
	}
}

func checkRows(window fyne.Window) {
	for i := 0; i < 3; i++ {
		if state[0][i] == state[1][i] && state[0][i] == state[2][i] && state[0][i] != 0 && state[1][i] != 0 && state[2][i] != 0 {
			endGame(true, state[0][i] == 1, window)
		}
	}
}

func checkDiagonals(window fyne.Window) {
	for i := 0; i <= 2; i += 2 {
		if state[0][i] == state[1][1] && state[0][i] == state[2][2-i] && state[0][i] != 0 && state[1][1] != 0 && state[2][2-i] != 0 {
			endGame(true, state[0][i] == 1, window)
		}
	}
}

func endGame(status bool, player bool, window fyne.Window) {
	var msg string
	if status {
		if player {
			counterWin2Player++
			label2Player.SetText(strconv.Itoa(counterWin2Player))
			msg = "Победа 2 игрока"
		} else {
			counterWin1Player++
			label1Player.SetText(strconv.Itoa(counterWin1Player))
			msg = "Победа 1 игрока"
		}
	} else {
		msg = "Ничья"
	}

	dialog.ShowConfirm("Игра окончена", msg+"\nначать заново?", func(b bool) {
		if b != true {
			lockBoard = true
		} else {
			restart()
		}
	}, window)
	winState = true
}

func checkAllDesk(window fyne.Window) {
	counter := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if state[i][j] == 0 {
				return
			} else {
				counter++
			}
		}
	}

	if counter == 9 && !winState {
		endGame(false, false, window)
	}
}

func restart() {
	winState, flag, lockBoard = false, false, false

	for _, i := range buttons {
		i.SetText("")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			state[i][j] = 0
		}
	}
}
