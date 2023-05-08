package controllers

import (
	"go-example/models"
	"sync"

	"github.com/gin-gonic/gin"
)

type PictureController interface {
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	pictures []models.Picture
}

func NewPictureController() PictureController {
	return &controller{
		pictures: make([]models.Picture, 0),
	}
}

type generator struct {
	counter int
	mtx     sync.Mutex
}

func (g *generator) getNextId() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	g.counter++
	return g.counter
}

var g *generator = &generator{}

func (c *controller) GetAll(ctx *gin.Context) {
	ctx.JSON(200, c.pictures)
}

func (c *controller) Update(ctx *gin.Context) {
	var picToUpdate models.Picture
	err := ctx.ShouldBindUri(&picToUpdate)
	if err != nil {
		ctx.String(400, "bad request %v", err)
		return
	}
	if err := ctx.ShouldBindJSON(&picToUpdate); err != nil {
		ctx.String(400, "bad request %v", err)
		return
	}

	for idx, pic := range c.pictures {
		if pic.Id == picToUpdate.Id {
			c.pictures[idx] = picToUpdate
			ctx.String(200, "success, picture with id %d has been updated", picToUpdate.Id)
			return
		}
	}
	ctx.String(400, "bad request cannot find picture with %d to update", picToUpdate)
}

func (c *controller) Create(ctx *gin.Context) {
	pic := models.Picture{Id: g.getNextId()}
	if err := ctx.BindJSON(&pic); err != nil {
		ctx.String(400, "bad request %v", err)
		return
	}
	c.pictures = append(c.pictures, pic)
	ctx.String(200, "success, new picture id is %d", pic.Id)
}

func (c *controller) Delete(ctx *gin.Context) {
	var picToDelete models.Picture
	if err := ctx.ShouldBindUri(&picToDelete); err != nil {
		ctx.String(400, "bad request %v", err)
		return
	}
	for idx, pic := range c.pictures {
		if pic.Id == picToDelete.Id {
			c.pictures = append(c.pictures[0:idx], c.pictures[idx+1:]...)
			ctx.String(200, "success, picture with id %d has been deleted", picToDelete.Id)
			return
		}
	}
	ctx.String(400, "bad request cannot picture with %d to delete", picToDelete.Id)
}
