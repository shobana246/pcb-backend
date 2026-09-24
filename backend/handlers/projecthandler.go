package handlers

import (
	"errors"
	"fmt"
	"pcb-backend/backend/models"
	"pcb-backend/backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProjectCreation(c *gin.Context) {
	var project models.ProjectRequest
	err := c.ShouldBindJSON(&project)
	if err != nil {
		c.JSON(400, gin.H{"message": "validation failed", "error": err.Error()})
		return
	}

	err = service.ProjectCreationService(project.ToModel())
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInput):
			c.JSON(400, gin.H{"message": "project_name and revision cannot be blank"})
		case errors.Is(err, service.ErrUserNotFound):
			c.JSON(400, gin.H{"message": "updated_by user does not exist"})
		case errors.Is(err, service.ErrProjectExists):
			c.JSON(409, gin.H{"message": "project with this name and revision already exists"})
		default:
			fmt.Println("create project error:", err)
			c.JSON(500, gin.H{"message": "Failed to create project"})
		}
		return
	}

	c.JSON(200, gin.H{"message": "Project successfully created"})
}

func GetAllProjects(c *gin.Context) {
	project, err := service.GetAllServiceProject()
	if err != nil {
		fmt.Println("get all projects error:", err)
		c.JSON(500, gin.H{"message": "Failed to fetch the projects"})
		return
	}
	c.JSON(200, project)
}

func GetProjectById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Id"})
		return
	}

	getData, err := service.GetProjectByIdService(id)
	if err != nil {
		if errors.Is(err, service.ErrProjectNotFound) {
			c.JSON(404, gin.H{"message": "Project not found"})
			return
		}
		fmt.Println("get project error:", err)
		c.JSON(500, gin.H{"message": "Failed to fetch the project"})
		return
	}
	c.JSON(200, getData)
}

func UpdateProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Id"})
		return
	}

	var update_data models.ProjectRequest
	err = c.ShouldBindJSON(&update_data)
	if err != nil {
		c.JSON(400, gin.H{"message": "validation failed", "error": err.Error()})
		return
	}

	err = service.ProjectUpdateService(id, update_data.ToModel())
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInput):
			c.JSON(400, gin.H{"message": "project_name and revision cannot be blank"})
		case errors.Is(err, service.ErrUserNotFound):
			c.JSON(400, gin.H{"message": "updated_by user does not exist"})
		case errors.Is(err, service.ErrProjectExists):
			c.JSON(409, gin.H{"message": "another project with this name and revision already exists"})
		case errors.Is(err, service.ErrProjectNotFound):
			c.JSON(404, gin.H{"message": "Project not found"})
		default:
			fmt.Println("update project error:", err)
			c.JSON(500, gin.H{"message": "Failed to update project"})
		}
		return
	}
	c.JSON(200, gin.H{"message": "project updated successfully"})
}

func DeleteProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Id"})
		return
	}

	err = service.ProjectDeleteService(id)
	if err != nil {
		if errors.Is(err, service.ErrProjectNotFound) {
			c.JSON(404, gin.H{"message": "Project not found"})
			return
		}
		fmt.Println("delete project error:", err)
		c.JSON(500, gin.H{"message": "Failed to delete project"})
		return
	}

	c.JSON(200, gin.H{"message": "Project Deleted Successfully"})
}
