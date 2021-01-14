package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/petrule"
)

type PetruleController struct {
	client *ent.Client
	router gin.IRouter
}

type Petrule struct {
	Petrule string
}

// CreatePetrule handles POST requests for adding petrule entities
// @Summary Create petrule
// @Description Create petrule
// @ID create-petrule
// @Accept   json
// @Produce  json
// @Param petrule body ent.Petrule true "Petrule entity"
// @Success 200 {object} ent.Petrule
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /petrules [post]
func (ctl *PetruleController) CreatePetrule(c *gin.Context) {
	obj := Petrule{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "petrule binding failed",
		})
		return
	}

	pr, err := ctl.client.Petrule.
		Create().
		SetPetrule(obj.Petrule).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pr)
}

// GetPetrule handles GET requests to retrieve a petrule entity
// @Summary Get a petrule entity by ID
// @Description get petrule by ID
// @ID get-petrule
// @Produce  json
// @Param id path int true "Petrule ID"
// @Success 200 {object} ent.Petrule
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /petrules/{id} [get]
func (ctl *PetruleController) GetPetrule(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	pr, err := ctl.client.Petrule.
		Query().
		Where(petrule.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pr)
}

// ListPetrule handles request to get a list of petrule entities
// @Summary List petrule entities
// @Description list petrule entities
// @ID list-petrule
// @Produce json
// @Success 200 {array} ent.Petrule
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /petrules [get]
func (ctl *PetruleController) ListPetrule(c *gin.Context) {
	petrules, err := ctl.client.Petrule.
		Query().
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, petrules)
}

// DeletePetrule handles DELETE requests to delete a petrule entity
// @Summary Delete a petrule entity by ID
// @Description get petrule by ID
// @ID delete-petrule
// @Produce  json
// @Param id path int true "Petrule ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /petrule/{id} [delete]
func (ctl *PetruleController) DeletePetrule(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Petrule.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewPetruleController creates and registers handles for the petrule controller
func NewPetruleController(router gin.IRouter, client *ent.Client) *PetruleController {
	prc := &PetruleController{
		client: client,
		router: router,
	}

	prc.register()

	return prc

}

func (ctl *PetruleController) register() {
	petrules := ctl.router.Group("/petrules")

	petrules.POST("", ctl.CreatePetrule)
	petrules.GET(":id", ctl.GetPetrule)
	petrules.GET("", ctl.ListPetrule)
	petrules.DELETE("", ctl.DeletePetrule)

}
