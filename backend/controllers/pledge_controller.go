package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/pledge"
)

type PledgeController struct {
	client *ent.Client
	router gin.IRouter
}

type Pledge struct {
	Provision string
}

// CreatePledge handles POST requests for adding pledge entities
// @Summary Create pledge
// @Description Create pledge
// @ID create-pledge
// @Accept   json
// @Produce  json
// @Param pledge body ent.Pledge true "Pledge entity"
// @Success 200 {object} ent.Pledge
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pledges [post]
func (ctl *PledgeController) CreatePledge(c *gin.Context) {
	obj := Pledge{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pledge binding failed",
		})
		return
	}

	p, err := ctl.client.Pledge.
		Create().
		SetProvision(obj.Provision).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, p)
}

// GetPledge handles GET requests to retrieve a pledge entity
// @Summary Get a pledge entity by ID
// @Description get pledge by ID
// @ID get-pledge
// @Produce  json
// @Param id path int true "Pledge ID"
// @Success 200 {object} ent.Pledge
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pledges/{id} [get]
func (ctl *PledgeController) GetPledge(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := ctl.client.Pledge.
		Query().
		Where(pledge.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListPledge handles request to get a list of pledge entities
// @Summary List pledge entities
// @Description list pledge entities
// @ID list-pledge
// @Produce json
// @Success 200 {array} ent.Pledge
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pledges [get]
func (ctl *PledgeController) ListPledge(c *gin.Context) {
	pledges, err := ctl.client.Pledge.
		Query().
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pledges)
}

// DeletePledge handles DELETE requests to delete a pledge entity
// @Summary Delete a pledge entity by ID
// @Description get pledge by ID
// @ID delete-pledge
// @Produce  json
// @Param id path int true "Pledge ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pledge/{id} [delete]
func (ctl *PledgeController) DeletePledge(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Pledge.
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

// NewPledgeController creates and registers handles for the pledge controller
func NewPledgeController(router gin.IRouter, client *ent.Client) *PledgeController {
	pc := &PledgeController{
		client: client,
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *PledgeController) register() {
	pledges := ctl.router.Group("/pledges")

	pledges.POST("", ctl.CreatePledge)
	pledges.GET(":id", ctl.GetPledge)
	pledges.GET("", ctl.ListPledge)
	pledges.DELETE("", ctl.DeletePledge)

}
