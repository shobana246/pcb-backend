package handlers

import (
	"database/sql"
	"pcb-backend/backend/models"
	"pcb-backend/backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComponent(c *gin.Context) {
	var component models.Component
	err := c.ShouldBindJSON(&component)
	if err != nil {
		c.JSON(400, gin.H{"messgae": "Invalid request"})
		return
	}
	err = service.ComponentCreationService(component)
	if err != nil {
		c.JSON(404, gin.H{"message": "Failed to create component", "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Component Created successfully"})
}

func GetComponent(c *gin.Context) {
	idparam := c.Query("project_id")
	ProjectId, err := strconv.Atoi(idparam)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	get_data, err := service.GetComponentService(ProjectId)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to get the component", "Error": err.Error()})
		return
	}
	if len(get_data) == 0 {
		c.JSON(404, gin.H{"message": "No components found for this project"})
		return
	}

	c.JSON(200, get_data)

}

func GetByIdComponent(c *gin.Context) {
	idparam := c.Param("id")
	Id, err := strconv.Atoi(idparam)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	get_data, err := service.GetByIdComponentService(Id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"message": "Component not found"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to get the component", "Error": err.Error()})
		return
	}
	c.JSON(200, get_data)

}

func UpdateComponent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"message": "Id not found"})
		return
	}

	var component models.Component
	err = c.ShouldBindJSON(&component)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	err = service.UpdateComponentService(id, component)
	if err != nil {
		c.JSON(404, gin.H{"message": "Component not found"})
		return
	}
	c.JSON(200, gin.H{"Message": "Component Updated Successfully"})
}

func DeleteComponent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	err = service.DeleteComponentService(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Component not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Component Deleted"})
}
