package trisnake

import (
	"fmt"
	"os"

	tl "github.com/JoelOtter/termloop"
	tb "github.com/nsf/termbox-go"
)

var (
	counterSnake = 10
	counterArena = 10
)

func (snake *Snake) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			if snake.Direction != left {
				snake.Direction = right
			}
		case tl.KeyArrowLeft:
			if snake.Direction != right {
				snake.Direction = left
			}
		case tl.KeyArrowUp:
			if snake.Direction != down {
				snake.Direction = up
			}
		case tl.KeyArrowDown:
			if snake.Direction != up {
				snake.Direction = down
			}
		}
	}
}

func (gos *Gameoverscreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyHome:
			RestartGame()
		case tl.KeyDelete:
			tb.Close()
			os.Exit(0)
		case tl.KeySpace:
			SaveHighScore(gs.Score, gs.FPS, Difficulty)
		}
	}
}

func (ts *Titlescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		if event.Key == tl.KeyEnter {
			gs = NewGamescreen()
			sg.Screen().SetLevel(gs)
		}
		if event.Key == tl.KeyInsert {
			gop := NewOptionsscreen()
			sg.Screen().SetLevel(gop)
		}
	}
}

func (g *Gameoptionsscreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyF1:
			ts.GameDifficulty = easy
			Difficulty = "Easy"
			gop.CurrentDifficultyText.SetText(fmt.Sprintf("Current difficulty: %s", Difficulty))
		case tl.KeyF2:
			ts.GameDifficulty = normal
			Difficulty = "Normal"
			gop.CurrentDifficultyText.SetText(fmt.Sprintf("Current difficulty: %s", Difficulty))

		case tl.KeyArrowUp:
			switch ColorObject {
			case "Snake":
				if counterSnake <= 10 {
					return
				}
				counterSnake -= 2
				gop.ColorSelectedIcon.SetPosition(73, counterSnake)

			case "Arena":
				if counterArena <= 10 {
					return
				}
				counterArena -= 2
				gop.ColorSelectedIcon.SetPosition(73, counterArena)
			}
		case tl.KeyArrowDown:
			switch ColorObject {
			case "Snake":
				if counterSnake >= 22 {
					return
				}
				counterSnake += 2
				gop.ColorSelectedIcon.SetPosition(73, counterSnake)

			case "Arena":
				if counterArena >= 22 {
					return
				}
				counterArena += 2
				gop.ColorSelectedIcon.SetPosition(73, counterArena)
			}
		case tl.KeyF3:
			ts.GameDifficulty = hard
			Difficulty = "Hard"
			gop.CurrentDifficultyText.SetText(fmt.Sprintf("Current difficulty: %s", Difficulty))

		case tl.KeyF4:
			ColorObject = "Snake"
			gop.CurrentColorObjectText.SetText(fmt.Sprintf("Current object: %s", ColorObject))

		case tl.KeyF5:
			ColorObject = "Food"
			gop.CurrentColorObjectText.SetText(fmt.Sprintf("Current object: %s", ColorObject))

		case tl.KeyF6:
			ColorObject = "Arena"
			gop.CurrentColorObjectText.SetText(fmt.Sprintf("Current object: %s", ColorObject))

		case tl.KeyEnter:
			gs = NewGamescreen()
			sg.Screen().SetLevel(gs)
		}

	}
}

func CheckSelectedColor(c int) tl.Attr {
	switch c {
	case 10:
		return tl.ColorWhite
	case 12:
		return tl.ColorRed
	case 14:
		return tl.ColorGreen
	case 16:
		return tl.ColorBlue
	case 18:
		return tl.ColorYellow
	case 20:
		return tl.ColorMagenta
	case 22:
		return tl.ColorCyan
	default:
		return tl.ColorDefault
	}
}
