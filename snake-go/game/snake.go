package trisnake

import tl "github.com/JoelOtter/termloop"

func NewSnake() *Snake {
	snake := new(Snake)
	snake.Entity = tl.NewEntity(5, 5, 1, 1)
	snake.Direction = right
	snake.Bodylength = []Coordinates{
		{1, 6}, // Tail
		{2, 6}, // Body
		{3, 6}, // Head
	}

	return snake
}

func (snake *Snake) Head() *Coordinates {
	return &snake.Bodylength[len(snake.Bodylength)-1]
}

func (snake *Snake) BorderCollision() bool {
	return gs.ArenaEntity.Contains(*snake.Head())
}
func (snake *Snake) FoodCollision() bool {
	return gs.FoodEntity.Contains(*snake.Head())
}
func (snake *Snake) SnakeCollision() bool {
	return snake.Contains()
}

func (snake *Snake) Contains() bool {
	for i := 0; i < len(snake.Bodylength)-1; i++ {
		if *snake.Head() == *&snake.Bodylength[i] {
			return true
		}
	}
	return false
}

func (snake *Snake) Draw(screen *tl.Screen) {
	sHead := *snake.Head()
	switch snake.Direction {
	case up:
		sHead.Y--
	case down:
		sHead.Y++
	case left:
		sHead.X--
	case right:
		sHead.X++
	}

	if snake.FoodCollision() {
		switch gs.FoodEntity.Emoji {
		case 'R':
			switch ts.GameDifficulty {
			case easy:
				if gs.FPS-3 <= 8 {
					UpdateScore(5)
				} else {
					gs.FPS -= 3
					UpdateScore(5)
					UpdateFPS()
				}
			case normal:
				if gs.FPS-2 <= 12 {
					UpdateScore(5)
				} else {
					gs.FPS -= 2
					UpdateScore(5)
					UpdateFPS()
				}
			case hard:
				if gs.FPS-1 <= 20 {
					UpdateScore(5)
				} else {
					gs.FPS--
					UpdateScore(5)
					UpdateFPS()
				}
			}
			snake.Bodylength = append(snake.Bodylength, sHead)
		case 'S':
			switch ts.GameDifficulty {
			case easy:
				gs.FPS++
			case normal:
				gs.FPS += 3
			case hard:
				gs.FPS += 5
			}
			UpdateFPS()
		default:
			UpdateScore(1)
			snake.Bodylength = append(snake.Bodylength, sHead)
		}
		gs.FoodEntity.MoveFood()
	} else {
		snake.Bodylength = append(snake.Bodylength[1:], sHead)
	}

	snake.SetPosition(sHead.X, sHead.Y)

	if snake.BorderCollision() || snake.SnakeCollision() {
		Gameover()
	}

	for _, c := range snake.Bodylength {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: CheckSelectedColor(counterSnake),
			Ch: '░',
		})
	}
}
