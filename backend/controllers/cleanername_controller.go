package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/cleanername"
)

// CleanernameController defines the struct for the cleanername controller
type CleanernameController struct {
	client *ent.Client
	router gin.IRouter
}

type Cleanername struct {
	cleanername string
}

// ListCleanername handles request to get a list of cleanername entities
// @Summary List cleanername entities
// @Description list cleanername entities
// @ID list-cleanername
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Cleanername
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cleanernames [get]
func (ctl *CleanernameController) ListCleanername(c *gin.Context) {
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

	cleanernames, err := ctl.client.Cleanername.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cleanernames)
}

// GetCleanername handles GET requests to retrieve a cleanername entity
// @Summary Get a cleanername entity by ID
// @Description get cleanername by ID
// @ID get-cleanername
// @Produce  json
// @Param id path int true "Cleanername ID"
// @Success 200 {object} ent.Cleanername
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cleanernames/{id} [get]
func (ctl *CleanernameController) GetCleanername(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Cleanername.
		Query().
		Where(cleanername.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// NewCleanernameController creates and registers handles for the cleanername controller
func NewCleanernameController(router gin.IRouter, client *ent.Client) *CleanernameController {
	sc := &CleanernameController{
		client: client,
		router: router,
	}

	sc.register()

	return sc

}

func (ctl *CleanernameController) register() {
	cleanernames := ctl.router.Group("/cleanernames")

	cleanernames.GET(":id", ctl.GetCleanername)
	cleanernames.GET("", ctl.ListCleanername)

}
