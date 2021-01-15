package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/situation"
)

type SituationController struct {
	client *ent.Client
	router gin.IRouter
}

type Situation struct {
	Situationname string
}

// CreateSituation handles POST requests for adding situation entities
// @Summary Create situation
// @Description Create situation
// @ID create-situation
// @Accept   json
// @Produce  json
// @Param situation body ent.Situation true "Situation entity"
// @Success 200 {object} ent.Situation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /situations [post]
func (ctl *SituationController) CreateSituation(c *gin.Context) {
	obj := Situation{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "situation binding failed",
		})
		return
	}

	si, err := ctl.client.Situation.
		Create().
		SetSituationname(obj.Situationname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, si)
}

// GetSituation handles GET requests to retrieve a situation entity
// @Summary Get a situation entity by ID
// @Description get situation by ID
// @ID get-situation
// @Produce  json
// @Param id path int true "Situation ID"
// @Success 200 {object} ent.Situation
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /situations/{id} [get]
func (ctl *SituationController) GetSituation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	si, err := ctl.client.Situation.
		Query().
		Where(situation.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, si)
}

// ListSituation handles request to get a list of situation entities
// @Summary List situation entities
// @Description list situation entities
// @ID list-situation
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Situation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /situations [get]
func (ctl *SituationController) ListSituation(c *gin.Context) {
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

	situations, err := ctl.client.Situation.
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

	c.JSON(200, situations)
}

// DeleteSituation handles DELETE requests to delete a situation entity
// @Summary Delete a situation entity by ID
// @Description get situation by ID
// @ID delete-situation
// @Produce  json
// @Param id path int true "Situation ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /situation/{id} [delete]
func (ctl *SituationController) DeleteSituation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Situation.
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

// NewSituationController creates and registers handles for the situation controller
func NewSituationController(router gin.IRouter, client *ent.Client) *SituationController {
	sic := &SituationController{
		client: client,
		router: router,
	}

	sic.register()

	return sic

}

func (ctl *SituationController) register() {
	situations := ctl.router.Group("/situations")

	situations.POST("", ctl.CreateSituation)
	situations.GET(":id", ctl.GetSituation)
	situations.GET("", ctl.ListSituation)
	situations.DELETE("", ctl.DeleteSituation)

}
