package handlers

import (
	"database/sql"
	"fmt"
	"pcb-backend/backend/models"
	"pcb-backend/backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProjectCreation(c *gin.Context) {

	var project models.Project

	err := c.ShouldBindJSON(&project)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}

	err = service.ProjectCreationService(project)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to create project", "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Project successfully created "})

}

func GetAllProjects(c *gin.Context) {
	project, err := service.GetAllServiceProject()
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to fetch the projects"})
		return
	}
	c.JSON(200, project)
}

func GetProjectById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Id"})
		return
	}

	getData, err := service.GetProjectByIdService(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Project not found"})
		return
	}
	c.JSON(200, getData)

}

func UpdateProject(c *gin.Context) {
	fmt.Println("reached the update handler--->")
	idparam := c.Param("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	var update_data models.Project
	err = c.ShouldBindJSON(&update_data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid resquest"})
		return
	}

	err = service.ProjectUpdateService(id, update_data)

	if err != nil {
		fmt.Println("erre in the service calling func", err)
		c.JSON(404, gin.H{"message": "project not found or update failed"})
		return
	}
	c.JSON(200, gin.H{"message": "project updated successfully"})

}

func DeleteProject(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid id"})
		return
	}

	err = service.ProjectDeleteService(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"message": "Project not found"})
			return
		}
		c.JSON(400, gin.H{"message": "Project not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Project Deleted Successfully"})

}
