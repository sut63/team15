package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team15/app/ent"
	"github.com/team15/app/ent/employee"
)

type EmployeeController struct {
	client *ent.Client
	router gin.IRouter
}

// GetEmployee handles GET requests to retrieve a employee entity
// @Summary Get a employee entity by ID
// @Description get employee by ID
// @ID get-employee
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} ent.Employee
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employees/{id} [get]
func (ctl *EmployeeController) GetEmployee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	em, err := ctl.client.Employee.
		Query().
		WithJobposition().
		Where(employee.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, em)
}

// ListEmployee handles request to get a list of employee entities
// @Summary List employee entities
// @Description list employee entities
// @ID list-employee
// @Produce json
// @Success 200 {array} ent.Employee
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employees [get]
func (ctl *EmployeeController) ListEmployee(c *gin.Context) {
	employees, err := ctl.client.Employee.
		Query().
		WithJobposition().
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, employees)
}

// DeleteEmployee handles DELETE requests to delete a employee entity
// @Summary Delete a employee entity by ID
// @Description get employee by ID
// @ID delete-employee
// @Produce  json
// @Param id path int true "Employee ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /employee/{id} [delete]
func (ctl *EmployeeController) DeleteEmployee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Employee.
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

// NewEmployeeController creates and registers handles for the employee controller
func NewEmployeeController(router gin.IRouter, client *ent.Client) *EmployeeController {
	emc := &EmployeeController{
		client: client,
		router: router,
	}

	emc.register()

	return emc

}

func (ctl *EmployeeController) register() {
	employees := ctl.router.Group("/employees")

	employees.GET(":id", ctl.GetEmployee)
	employees.GET("", ctl.ListEmployee)
	employees.DELETE("", ctl.DeleteEmployee)

}
