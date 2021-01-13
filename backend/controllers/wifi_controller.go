package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/wifi"
)

type WifiController struct {
	client *ent.Client
	router gin.IRouter
}

type Wifi struct {
	Wifiname string
}

// CreateWifi handles POST requests for adding wifi entities
// @Summary Create wifi
// @Description Create wifi
// @ID create-wifi
// @Accept   json
// @Produce  json
// @Param wifi body ent.Wifi true "Wifi entity"
// @Success 200 {object} ent.Wifi
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /wifis [post]
func (ctl *WifiController) CreateWifi(c *gin.Context) {
	obj := Wifi{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "wifi binding failed",
		})
		return
	}

	s, err := ctl.client.Wifi.
		Create().
		SetWifiname(obj.Wifiname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// GetWifi handles GET requests to retrieve a wifi entity
// @Summary Get a wifi entity by ID
// @Description get wifi by ID
// @ID get-wifi
// @Produce  json
// @Param id path int true "Wifi ID"
// @Success 200 {object} ent.Wifi
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /wifis/{id} [get]
func (ctl *WifiController) GetWifi(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := ctl.client.Wifi.
		Query().
		Where(wifi.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListWifi handles request to get a list of wifi entities
// @Summary List wifi entities
// @Description list wifi entities
// @ID list-wifi
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Wifi
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /wifis [get]
func (ctl *WifiController) ListWifi(c *gin.Context) {
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

	wifis, err := ctl.client.Wifi.
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

	c.JSON(200, wifis)
}

// DeleteWifi handles DELETE requests to delete a wifi entity
// @Summary Delete a wifi entity by ID
// @Description get wifi by ID
// @ID delete-wifi
// @Produce  json
// @Param id path int true "Wifi ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /wifis/{id} [delete]
func (ctl *WifiController) DeleteWifi(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Wifi.
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

// NewWifiController creates and registers handles for the wifi controller
func NewWifiController(router gin.IRouter, client *ent.Client) *WifiController {
	wf := &WifiController{
		client: client,
		router: router,
	}

	wf.register()

	return wf

}

func (ctl *WifiController) register() {
	wifis := ctl.router.Group("/wifis")

	wifis.POST("", ctl.CreateWifi)
	wifis.GET(":id", ctl.GetWifi)
	wifis.GET("", ctl.ListWifi)
	wifis.DELETE("", ctl.DeleteWifi)

}
