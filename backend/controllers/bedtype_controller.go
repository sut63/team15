package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/bedtype"
)

type BedtypeController struct {
	client *ent.Client
	router gin.IRouter
}

type Bedtype struct {
	Bedtypename string
}

// CreateBedtype handles POST requests for adding bedtype entities
// @Summary Create bedtype
// @Description Create bedtype
// @ID create-bedtype
// @Accept   json
// @Produce  json
// @Param bedtype body ent.Bedtype true "Bedtype entity"
// @Success 200 {object} ent.Bedtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bedtypes [post]
func (ctl *BedtypeController) CreateBedtype(c *gin.Context) {
	obj := Bedtype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bedtype binding failed",
		})
		return
	}

	bt, err := ctl.client.Bedtype.
		Create().
		SetBedtypename(obj.Bedtypename).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, bt)
}

// GetBedtype handles GET requests to retrieve a bedtype entity
// @Summary Get a bedtype entity by ID
// @Description get bedtype by ID
// @ID get-bedtype
// @Produce  json
// @Param id path int true "Bedtype ID"
// @Success 200 {object} ent.Bedtype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bedtypes/{id} [get]
func (ctl *BedtypeController) GetBedtype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	bt, err := ctl.client.Bedtype.
		Query().
		Where(bedtype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bt)
}

// ListBedtype handles request to get a list of bedtype entities
// @Summary List bedtype entities
// @Description list bedtype entities
// @ID list-bedtype
// @Produce json
// @Success 200 {array} ent.Bedtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bedtypes [get]
func (ctl *BedtypeController) ListBedtype(c *gin.Context) {
	bedtypes, err := ctl.client.Bedtype.
		Query().
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bedtypes)
}

// DeleteBedtype handles DELETE requests to delete a bedtype entity
// @Summary Delete a bedtype entity by ID
// @Description get bedtype by ID
// @ID delete-bedtype
// @Produce  json
// @Param id path int true "Bedtype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bedtype/{id} [delete]
func (ctl *BedtypeController) DeleteBedtype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Bedtype.
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

// NewBedtypeController creates and registers handles for the bedtype controller
func NewBedtypeController(router gin.IRouter, client *ent.Client) *BedtypeController {
	btc := &BedtypeController{
		client: client,
		router: router,
	}

	btc.register()

	return btc

}

func (ctl *BedtypeController) register() {
	bedtypes := ctl.router.Group("/bedtypes")

	bedtypes.POST("", ctl.CreateBedtype)
	bedtypes.GET(":id", ctl.GetBedtype)
	bedtypes.GET("", ctl.ListBedtype)
	bedtypes.DELETE("", ctl.DeleteBedtype)

}
