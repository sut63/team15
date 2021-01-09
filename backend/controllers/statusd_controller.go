package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
    "github.com/team15/app/ent/statusd"
)

type StatusdController struct {
	client *ent.Client
	router gin.IRouter
}

type Statusd struct {
	Statusdname string
}

// CreateStatusd handles POST requests for adding statusd entities
// @Summary Create statusd
// @Description Create statusd
// @ID create-statusd
// @Accept   json
// @Produce  json
// @Param statusd body ent.Statusd true "Statusd entity"
// @Success 200 {object} ent.Statusd
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusds [post]
func (ctl *StatusdController) CreateStatusd(c *gin.Context) {
	obj := Statusd{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "statusd binding failed",
		})
		return
	}

	s, err := ctl.client.Statusd.
		Create().
		SetStatusdname(obj.Statusdname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// GetStatusd handles GET requests to retrieve a statusd entity
// @Summary Get a statusd entity by ID
// @Description get statusd by ID
// @ID get-statusd
// @Produce  json
// @Param id path int true "Statusd ID"
// @Success 200 {object} ent.Statusd
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusds/{id} [get]
func (ctl *StatusdController) GetStatusd(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := ctl.client.Statusd.
		Query().
		Where(statusd.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListStatusd handles request to get a list of statusd entities
// @Summary List statusd entities
// @Description list statusd entities
// @ID list-statusd
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Statusd
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusds [get]
func (ctl *StatusdController) ListStatusd(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	statusds, err := ctl.client.Statusd.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, statusds)
}

// DeleteStatusd handles DELETE requests to delete a statusd entity
// @Summary Delete a statusd entity by ID
// @Description get statusd by ID
// @ID delete-statusd
// @Produce  json
// @Param id path int true "Statusd ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statusd/{id} [delete]
func (ctl *StatusdController) DeleteStatusd(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Statusd.
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

// NewStatusdController creates and registers handles for the statusd controller
func NewStatusdController(router gin.IRouter, client *ent.Client) *StatusdController {
	sc := &StatusdController{
		client: client,
		router: router,
	}

	sc.register()

	return sc

}

func (ctl *StatusdController) register() {
	statusds := ctl.router.Group("/statusds")

	statusds.POST("", ctl.CreateStatusd)
	statusds.GET(":id", ctl.GetStatusd)
	statusds.GET("", ctl.ListStatusd)
	statusds.DELETE("", ctl.DeleteStatusd)

}