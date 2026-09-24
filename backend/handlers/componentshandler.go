package handlers

import (
	"errors"
	"fmt"
	"pcb-backend/backend/models"
	"pcb-backend/backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComponent(c *gin.Context) {
	var req models.CreateComponentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"message": "validation failed", "error": err.Error()})
		return
	}

	err = service.ComponentCreationService(req.ToModel())
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInput):
			c.JSON(400, gin.H{"message": "component_name cannot be blank"})
		case errors.Is(err, service.ErrProjectNotFound):
			c.JSON(404, gin.H{"message": "project not found"})
		case errors.Is(err, service.ErrComponentExists):
			c.JSON(409, gin.H{"message": "component with this name already exists in this project"})
		default:
			fmt.Println("create component error:", err)
			c.JSON(500, gin.H{"message": "Failed to create component"})
		}
		return
	}

	c.JSON(200, gin.H{"message": "Component created successfully"})
}

func GetComponent(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Query("project_id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "project_id is required and must be a number"})
		return
	}

	components, err := service.GetComponentService(projectID)
	if err != nil {
		if errors.Is(err, service.ErrProjectNotFound) {
			c.JSON(404, gin.H{"message": "project not found"})
			return
		}
		fmt.Println("get components error:", err)
		c.JSON(500, gin.H{"message": "Failed to get the components"})
		return
	}

	c.JSON(200, components)
}

func GetByIdComponent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Id"})
		return
	}

	component, err := service.GetByIdComponentService(id)
	if err != nil {
		if errors.Is(err, service.ErrComponentNotFound) {
			c.JSON(404, gin.H{"message": "Component not found"})
			return
		}
		fmt.Println("get component error:", err)
		c.JSON(500, gin.H{"message": "Failed to get the component"})
		return
	}
	c.JSON(200, component)
}

func UpdateComponent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Id"})
		return
	}

	var req models.ComponentFields
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"message": "validation failed", "error": err.Error()})
		return
	}

	err = service.UpdateComponentService(id, req.ToModel())
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInput):
			c.JSON(400, gin.H{"message": "component_name cannot be blank"})
		case errors.Is(err, service.ErrComponentNotFound):
			c.JSON(404, gin.H{"message": "Component not found"})
		case errors.Is(err, service.ErrComponentExists):
			c.JSON(409, gin.H{"message": "another component with this name already exists in this project"})
		default:
			fmt.Println("update component error:", err)
			c.JSON(500, gin.H{"message": "Failed to update component"})
		}
		return
	}
	c.JSON(200, gin.H{"message": "Component updated successfully"})
}

func DeleteComponent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Id"})
		return
	}

	err = service.DeleteComponentService(id)
	if err != nil {
		if errors.Is(err, service.ErrComponentNotFound) {
			c.JSON(404, gin.H{"message": "Component not found"})
			return
		}
		fmt.Println("delete component error:", err)
		c.JSON(500, gin.H{"message": "Failed to delete component"})
		return
	}
	c.JSON(200, gin.H{"message": "Component deleted successfully"})
}
