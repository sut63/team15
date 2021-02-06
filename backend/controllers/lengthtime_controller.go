package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/lengthtime"
)

// LengthtimeController defines the struct for the lengthtime controller
type LengthtimeController struct {
	client *ent.Client
	router gin.IRouter
}

type Lengthtime struct {
	lengthtime string
}

// ListLengthtime handles request to get a list of lengthtime entities
// @Summary List lengthtime entities
// @Description list lengthtime entities
// @ID list-lengthtime
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Lengthtime
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /lengthtimes [get]
func (ctl *LengthtimeController) ListLengthtime(c *gin.Context) {
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

	lengthtimes, err := ctl.client.Lengthtime.
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

// GetLengthtime handles GET requests to retrieve a lengthtime entity
// @Summary Get a lengthtime entity by ID
// @Description get lengthtime by ID
// @ID get-lengthtime
// @Produce  json
// @Param id path int true "Lengthtime ID"
// @Success 200 {object} ent.Lengthtime
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /lengthtimes/{id} [get]
func (ctl *LengthtimeController) GetLengthtime(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Lengthtime.
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

// NewLengthtimeController creates and registers handles for the lengthtime controller
func NewLengthtimeController(router gin.IRouter, client *ent.Client) *LengthtimeController {
	sc := &LengthtimeController{
		client: client,
		router: router,
	}

	sc.register()

	return sc

}

func (ctl *LengthtimeController) register() {
	lengthtimes := ctl.router.Group("/lengthtimes")

	lengthtimes.GET(":id", ctl.GetLengthtime)
	lengthtimes.GET("", ctl.ListLengthtime)

}
