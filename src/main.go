package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var buttons [9]*widget.Button
var state = [3][3]uint{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
var flag bool
var winState bool

func main() {
	a := app.New()
	w := a.NewWindow("TikTakToe")
	w.Resize(fyne.NewSize(300, 300))

	for i := 0; i < 9; i++ {
		x, y := getPositionInMatrix(i)
		var button *widget.Button
		button = widget.NewButton("", func() {
			x1, y1 := x, y
			sta := changeState(x1, y1)
			fmt.Println(x1, y1)
			button.SetText(sta)
			check(w)
		})

		buttons[i] = button
	}

	line1 := container.NewGridWithColumns(3, buttons[0], buttons[1], buttons[2])
	line2 := container.NewGridWithColumns(3, buttons[3], buttons[4], buttons[5])
	line3 := container.NewGridWithColumns(3, buttons[6], buttons[7], buttons[8])

	content := container.NewGridWithRows(3, line1, line2, line3)

	w.SetContent(content)
	w.ShowAndRun()
}

func getPositionInMatrix(pos int) (x int, y int) {
	y = pos / 3
	x = pos % 3

	return x, y
}

func changeState(x int, y int) (st string) {
	if state[x][y] != 0 {
		if state[x][y] == 1 {
			return "O"
		} else {
			return "X"
		}
	}

	if flag {
		flag = !flag
		state[x][y] = 1
		return "O"
	} else {
		flag = !flag
		state[x][y] = 2
		return "X"
	}
}

func check(window fyne.Window) {
	checkLeftColumn(window)
	checkMiddleColumn(window)
	checkRightColumn(window)
	checkTopRow(window)
	checkMiddleRow(window)
	checkBottomRow(window)
	checkMainDiagonal(window)
	checkSecondDiagonal(window)
	checkAallDesk(window)
}

func checkLeftColumn(window fyne.Window) bool {
	if state[0][0] == state[0][1] && state[0][0] == state[0][2] && state[0][0] != 0 && state[0][1] != 0 && state[0][2] != 0 {
		if state[0][0] == 1 {
			secondWin(window)
		} else {
			firstWin(window)
		}
		return true
	}

	return false
}

func checkMiddleColumn(window fyne.Window) {
	if state[1][0] == state[1][1] && state[1][0] == state[1][2] && state[1][0] != 0 && state[1][1] != 0 && state[1][2] != 0 {
		if state[1][0] == 1 {
			secondWin(window)
		} else {
			firstWin(window)
		}
	}
}

func checkRightColumn(window fyne.Window) {
	if state[2][0] == state[2][1] && state[2][0] == state[2][2] && state[2][0] != 0 && state[2][1] != 0 && state[2][2] != 0 {
		if state[2][0] == 1 {
			secondWin(window)
		} else {
			firstWin(window)
		}
	}
}

func checkTopRow(window fyne.Window) {
	if state[0][0] == state[1][0] && state[0][0] == state[2][0] && state[0][0] != 0 && state[1][0] != 0 && state[2][0] != 0 {
		if state[0][0] == 1 {
			secondWin(window)
		} else {
			firstWin(window)
		}
	}
}

func checkMiddleRow(window fyne.Window) {
	if state[0][1] == state[1][1] && state[0][1] == state[2][1] && state[0][1] != 0 && state[1][1] != 0 && state[2][1] != 0 {
		if state[0][1] == 1 {
			secondWin(window)
		} else {
			firstWin(window)
		}
	}
}

func checkBottomRow(window fyne.Window) {
	if state[0][2] == state[1][2] && state[0][2] == state[2][2] && state[0][2] != 0 && state[1][2] != 0 && state[2][2] != 0 {
		if state[0][2] == 1 {
			secondWin(window)
		} else {
			firstWin(window)
		}
	}
}

func checkMainDiagonal(window fyne.Window) {
	if state[0][0] == state[1][1] && state[0][0] == state[2][2] && state[0][0] != 0 && state[1][1] != 0 && state[2][2] != 0 {
		if state[0][0] == 1 {
			secondWin(window)
		} else {
			firstWin(window)
		}
	}
}

func checkSecondDiagonal(window fyne.Window) {
	if state[0][2] == state[1][1] && state[0][2] == state[2][0] && state[0][2] != 0 && state[1][1] != 0 && state[2][0] != 0 {
		if state[0][2] == 1 {
			secondWin(window)
		} else {
			firstWin(window)
		}
	}
}

func checkAallDesk(window fyne.Window) {
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
		draw(window)
	}
}

func draw(window fyne.Window) {
	dialog.ShowConfirm("Игра окончена", "Ничья \nНачать заново?", func(b bool) {
		if b != true {
			return
		} else {
			restart()
		}
	}, window)
}

func firstWin(window fyne.Window) {
	dialog.ShowConfirm("Игра окончена", "Победа 1 игрока \nНачать заново?", func(b bool) {
		if b != true {
			return
		} else {
			restart()
		}
	}, window)
	winState = true
}

func secondWin(window fyne.Window) {
	dialog.ShowConfirm("Игра окончена", "Победа 2 игрока \nНачать заново?", func(b bool) {
		if b != true {
			return
		} else {
			restart()
		}
	}, window)
	winState = true
}

func restart() {
	winState = false

	for _, i := range buttons {
		i.SetText("")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			state[i][j] = 0
		}
	}
}
