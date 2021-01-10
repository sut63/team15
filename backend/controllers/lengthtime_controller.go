package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/lengthtime"
)

// LengthTimeController defines the struct for the lengthtime controller
type LengthTimeController struct {
	client *ent.Client
	router gin.IRouter
}

type LengthTime struct {
	lengthtime string
}

// ListLengthTime handles request to get a list of lengthtime entities
// @Summary List lengthtime entities
// @Description list lengthtime entities
// @ID list-lengthtime
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.LengthTime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /lengthtimes [get]
func (ctl *LengthTimeController) ListLengthTime(c *gin.Context) {
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

	lengthtimes, err := ctl.client.LengthTime.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, lengthtimes)
}

// GetLengthTime handles GET requests to retrieve a lengthtime entity
// @Summary Get a lengthtime entity by ID
// @Description get lengthtime by ID
// @ID get-lengthtime
// @Produce  json
// @Param id path int true "LengthTime ID"
// @Success 200 {object} ent.LengthTime
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /lengthtimes/{id} [get]
func (ctl *LengthTimeController) GetLengthTime(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.LengthTime.
		Query().
		Where(lengthtime.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// NewLengthTimeController creates and registers handles for the lengthtime controller
func NewLengthTimeController(router gin.IRouter, client *ent.Client) *LengthTimeController {
	sc := &LengthTimeController{
		client: client,
		router: router,
	}

	sc.register()

	return sc

}

func (ctl *LengthTimeController) register() {
	lengthtimes := ctl.router.Group("/lengthtimes")

	lengthtimes.GET(":id", ctl.GetLengthTime)
	lengthtimes.GET("", ctl.ListLengthTime)

}
