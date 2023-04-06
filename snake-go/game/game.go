package trisnake

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	tl "github.com/JoelOtter/termloop"
)

func StartGame() {
	sg = tl.NewGame()
	ts := NewTitlescreen()
	ts.AddEntity(ts.Logo)

	for _, v := range ts.OptionsText {
		ts.AddEntity(v)
	}

	sg.Screen().SetFps(10)
	sg.Screen().SetLevel(ts)
	sg.Start()
}

func NewTitlescreen() *Titlescreen {
	ts = new(Titlescreen)
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	logofile, _ := ioutil.ReadFile("util/titlescreen-logo.txt")
	ts.Logo = tl.NewEntityFromCanvas(10, 3, tl.CanvasFromString(string(logofile)))

	ts.GameDifficulty = normal
	ts.OptionsText = []*tl.Text{
		tl.NewText(10, 15, "Press ENTER to start!", tl.ColorWhite, tl.ColorBlack),
		tl.NewText(10, 17, "Press INSERT for options!", tl.ColorWhite, tl.ColorBlack),
	}

	return ts
}

func NewOptionsscreen() *Gameoptionsscreen {
	gop = new(Gameoptionsscreen)
	gop.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	gop.ColorPanelBackground = tl.NewRectangle(43, 3, 33, 21, tl.ColorWhite)
	gop.DifficultyBackground = tl.NewRectangle(5, 3, 33, 10, tl.ColorWhite)
	gop.ObjectBackground = tl.NewRectangle(5, 15, 33, 9, tl.ColorWhite)

	gop.StartText = tl.NewText(2, 1, "Press ENTER to start", tl.ColorWhite, tl.ColorBlack)
	gop.CurrentDifficultyText = tl.NewText(6, 4, fmt.Sprintf("Current difficulty: %s", Difficulty), tl.ColorBlack, tl.ColorWhite)
	gop.CurrentColorObjectText = tl.NewText(44, 4, fmt.Sprintf("Current Object: %s", ColorObject), tl.ColorBlack, tl.ColorWhite)
	gop.ColorSelectedIcon = tl.NewText(73, 10, "■", tl.ColorBlack, tl.ColorWhite)

	gop.ColorPanelOptions = []string{
		"Use ↑ ↓ to change colors",
		"White",
		"Red",
		"Green",
		"Blue",
		"Yellow",
		"Magenta",
		"Cyan",
	}

	gop.DifficultyOptions = []*tl.Text{
		tl.NewText(6, 7, fmt.Sprintf("Press F1 for Easy(8 speed)"), tl.ColorBlack, tl.ColorWhite),
		tl.NewText(6, 9, fmt.Sprintf("Press F2 for Normal(12 speed)"), tl.ColorBlack, tl.ColorWhite),
		tl.NewText(6, 11, fmt.Sprintf("Press F3 for Hard(25 speed)"), tl.ColorBlack, tl.ColorWhite),
	}

	gop.ColorObjectOptions = []*tl.Text{
		tl.NewText(6, 16, fmt.Sprintf("Press F4 for Snake (Colors)"), tl.ColorBlack, tl.ColorWhite),
		tl.NewText(6, 18, fmt.Sprintf("Press F6 for Arena (Colors)"), tl.ColorBlack, tl.ColorWhite),
	}

	gop.AddEntity(gop.DifficultyBackground)
	gop.AddEntity(gop.ColorPanelBackground)
	gop.AddEntity(gop.ObjectBackground)
	gop.AddEntity(gop.CurrentDifficultyText)
	gop.AddEntity(gop.CurrentColorObjectText)
	gop.AddEntity(gop.ColorSelectedIcon)
	gop.AddEntity(gop.StartText)

	for _, v := range gop.DifficultyOptions {
		gop.AddEntity(v)
	}

	y := 6
	for _, vv := range gop.ColorPanelOptions {
		var i *tl.Text
		if y == 6 {
			i = tl.NewText(44, y, vv, tl.ColorBlack, tl.ColorWhite)
			gop.AddEntity(i)
			y += 2
		} else {
			y += 2
			i = tl.NewText(44, y, vv, tl.ColorBlack, tl.ColorWhite)
			gop.AddEntity(i)
		}
	}
	for _, vvv := range gop.ColorObjectOptions {
		gop.AddEntity(vvv)
	}
	return gop
}

func NewGamescreen() *Gamescreen {
	gs = new(Gamescreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	SetDifficultyFPS()
	gs.Score = 0
	gs.SnakeEntity = NewSnake()
	gs.ArenaEntity = NewArena(70, 25)
	gs.FoodEntity = NewFood()
	gs.SidepanelObject = NewSidepanel()

	gs.AddEntity(gs.FoodEntity)
	gs.AddEntity(gs.SidepanelObject.Background)
	gs.AddEntity(gs.SidepanelObject.ScoreText)
	gs.AddEntity(gs.SidepanelObject.SpeedText)
	gs.AddEntity(gs.SidepanelObject.DifficultyText)
	gs.AddEntity(gs.SnakeEntity)
	gs.AddEntity(gs.ArenaEntity)

	y := 7
	for _, v := range sp.Instructions {
		var i *tl.Text
		y += 2
		i = tl.NewText(70+2, y, v, tl.ColorBlack, tl.ColorWhite)
		gs.AddEntity(i)
	}

	sg.Screen().SetFps(gs.FPS)

	return gs
}

func SetDifficultyFPS() {
	switch ts.GameDifficulty {
	case easy:
		gs.FPS = 8
	case normal:
		gs.FPS = 12
	case hard:
		gs.FPS = 25
	}
}

func NewSidepanel() *Sidepanel {
	sp = new(Sidepanel)
	sp.Instructions = []string{
		"Use ← → ↑ ↓ to move the snake around",
		"Pick up the food to grow bigger",
		"■: 1 point/growth",
		"R: 5 points (removes some speed!)",
		"S: 1 point (increased speed!!)",
	}

	sp.Background = tl.NewRectangle(70+1, 0, 45, 25, tl.ColorWhite)
	sp.ScoreText = tl.NewText(70+2, 1, fmt.Sprintf("Score: %d", gs.Score), tl.ColorBlack, tl.ColorWhite)
	sp.SpeedText = tl.NewText(70+2, 3, fmt.Sprintf("Speed: %.0f", gs.FPS), tl.ColorBlack, tl.ColorWhite)
	sp.DifficultyText = tl.NewText(70+2, 5, fmt.Sprintf("Difficulty: %s", Difficulty), tl.ColorBlack, tl.ColorWhite)
	return sp
}

func Gameover() {
	gos := new(Gameoverscreen)
	gos.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	logofile, _ := ioutil.ReadFile("util/gameover-logo.txt")
	gos.Logo = tl.NewEntityFromCanvas(10, 3, tl.CanvasFromString(string(logofile)))
	gos.Finalstats = []*tl.Text{
		tl.NewText(10, 13, fmt.Sprintf("Score: %d", gs.Score), tl.ColorWhite, tl.ColorBlack),
		tl.NewText(10, 15, fmt.Sprintf("Speed: %.0f", gs.FPS), tl.ColorWhite, tl.ColorBlack),
		tl.NewText(10, 17, fmt.Sprintf("Difficulty: %s", Difficulty), tl.ColorWhite, tl.ColorBlack),
	}
	gos.OptionsBackground = tl.NewRectangle(45, 12, 45, 7, tl.ColorWhite)
	gos.OptionsText = []*tl.Text{
		tl.NewText(47, 13, "Press \"Home\" to restart!", tl.ColorBlack, tl.ColorWhite),
		tl.NewText(47, 15, "Press \"Delete\" to quit!", tl.ColorBlack, tl.ColorWhite),
		tl.NewText(47, 17, "Press \"Spacebar\" to save your score!", tl.ColorBlack, tl.ColorWhite),
	}

	for _, v := range gos.Finalstats {
		gos.AddEntity(v)
	}
	gos.AddEntity(gos.Logo)
	gos.AddEntity(gos.OptionsBackground)

	for _, vv := range gos.OptionsText {
		gos.AddEntity(vv)
	}

	sg.Screen().SetLevel(gos)
}

func UpdateScore(amount int) {
	gs.Score += amount
	sp.ScoreText.SetText(fmt.Sprintf("Score: %d", gs.Score))
}

func UpdateFPS() {
	sg.Screen().SetFps(gs.FPS)
	sp.ScoreText.SetText(fmt.Sprintf("Score: %.0f", gs.FPS))
}

func RestartGame() {
	gs.RemoveEntity(gs.SnakeEntity)
	gs.RemoveEntity(gs.FoodEntity)

	gs.SnakeEntity = NewSnake()
	gs.FoodEntity = NewFood()

	SetDifficultyFPS()
	gs.Score = 0

	sp.ScoreText.SetText(fmt.Sprintf("Score: %d", gs.Score))
	sp.SpeedText.SetText(fmt.Sprintf("Speed: %.0f", gs.FPS))

	gs.AddEntity(gs.SnakeEntity)
	gs.AddEntity(gs.FoodEntity)
	sg.Screen().SetFps(gs.FPS)
	sg.Screen().SetLevel(gs)
}

func SaveHighScore(score int, speed float64, difficulty string) {
	var newRow []byte
	datetime := time.Now().Local()
	newRow = []byte(fmt.Sprintf("\n|" + fmt.Sprintf("%s", datetime.Format("2006-01-01 15:04:05")) + " | " + fmt.Sprintf("%d", score) + " | " + fmt.Sprintf("%.0f", speed) + " | " + difficulty + "| "))

	if _, existerr := os.Stat("HIGHSCORES.md"); existerr != nil {
		if os.IsNotExist(existerr) {
			_, err := os.Create("HIGHSCORES.md")
			if err != nil {
				log.Fatal("Create file error: ", err)
			}
		} else {
			log.Fatal("other file error: ", existerr)
		}
	}

	f, err := os.OpenFile("HIGHSCORES.md", os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	_, err2 := f.Write(newRow)
	if err2 != nil {
		log.Fatalf("Error writing to file: %s", err2)
	}

}
