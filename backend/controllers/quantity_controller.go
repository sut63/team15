package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/quantity"
)

type QuantityController struct {
	client *ent.Client
	router gin.IRouter
}

type Quantity struct {
	Quantity string
}

// CreateQuantity handles POST requests for adding quantity entities
// @Summary Create quantity
// @Description Create quantity
// @ID create-quantity
// @Accept   json
// @Produce  json
// @Param quantity body ent.Quantity true "Quantity entity"
// @Success 200 {object} ent.Quantity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /quantitys [post]
func (ctl *QuantityController) CreateQuantity(c *gin.Context) {
	obj := Quantity{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "quantity binding failed",
		})
		return
	}

	qu, err := ctl.client.Quantity.
		Create().
		SetQuantity(obj.Quantity).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, qu)
}

// GetQuantity handles GET requests to retrieve a quantity entity
// @Summary Get a quantity entity by ID
// @Description get quantity by ID
// @ID get-quantity
// @Produce  json
// @Param id path int true "Quantity ID"
// @Success 200 {object} ent.Quantity
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /quantitys/{id} [get]
func (ctl *QuantityController) GetQuantity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	qu, err := ctl.client.Quantity.
		Query().
		Where(quantity.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, qu)
}

// ListQuantity handles request to get a list of quantity entities
// @Summary List quantity entities
// @Description list quantity entities
// @ID list-quantity
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Quantity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /quantitys [get]
func (ctl *QuantityController) ListQuantity(c *gin.Context) {
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

	quantitys, err := ctl.client.Quantity.
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

	c.JSON(200, quantitys)
}

// DeleteQuantity handles DELETE requests to delete a quantity entity
// @Summary Delete a quantity entity by ID
// @Description get quantity by ID
// @ID delete-quantity
// @Produce  json
// @Param id path int true "Quantity ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /quantity/{id} [delete]
func (ctl *QuantityController) DeleteQuantity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Quantity.
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

// NewQuantityController creates and registers handles for the quantity controller
func NewQuantityController(router gin.IRouter, client *ent.Client) *QuantityController {
	quc := &QuantityController{
		client: client,
		router: router,
	}

	quc.register()

	return quc

}

func (ctl *QuantityController) register() {
	quantitys := ctl.router.Group("/quantitys")

	quantitys.POST("", ctl.CreateQuantity)
	quantitys.GET(":id", ctl.GetQuantity)
	quantitys.GET("", ctl.ListQuantity)
	quantitys.DELETE("", ctl.DeleteQuantity)

}
