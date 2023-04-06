package trisnake

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

var (
	insideborderW = 70 - 1
	insideborderH = 25 - 1
)

func NewFood() *Food {
	food := new(Food)
	food.Entity = tl.NewEntity(1, 1, 1, 1)
	food.MoveFood()

	return food
}

func (food *Food) MoveFood() {
	NewX := RandomInsideArena(insideborderW, 1)
	NewY := RandomInsideArena(insideborderH, 1)

	food.Foodposition.X = NewX
	food.Foodposition.Y = NewY
	food.Emoji = RandomFood()

	food.SetPosition(food.Foodposition.X, food.Foodposition.Y)
}

func RandomFood() rune {
	emoji := []rune{
		'R',
		'■',
		'■',
		'■',
		'■',
		'■',
		'■',
		'■',
		'■',
		'■',
		'■',
		'S',
	}

	rand.Seed(time.Now().UnixNano())

	return emoji[rand.Intn(len(emoji))]
}

func (food *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(food.Foodposition.X, food.Foodposition.Y, &tl.Cell{
		Ch: food.Emoji,
	})
}

func (food *Food) Contains(c Coordinates) bool {
	return c.X == food.Foodposition.X && c.Y == food.Foodposition.Y
}

func RandomInsideArena(iMax, iMin int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(iMax-iMin) + iMin
}
