package handlers

import (
	"database/sql"
	"pcb-backend/backend/models"
	"pcb-backend/backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateIssues(c *gin.Context) {
	var issues models.Issue
	err := c.ShouldBindJSON(&issues)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	err = service.CreateIssuesService(issues)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to create issue", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Placement Issue created"})

}

func GetAllIssues(c *gin.Context) {

	issues, err := service.GetAllIssuesService()
	if err != nil {
		c.JSON(500, gin.H{"message": "Placement Issue not found"})
		return
	}
	c.JSON(200, issues)

}

func GetIssuesById(c *gin.Context) {
	idparam := c.Param("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	data, err := service.GetIssuesByIdService(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Placement issue not found"})
		return
	}
	c.JSON(200, data)
}

func UpdateIssueById(c *gin.Context) {
	idparam := c.Param("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	var update_data models.Issue
	err = c.ShouldBindJSON(&update_data)
	if err != nil {
		c.JSON(400, gin.H{"Message": "Invalid Request"})
		return
	}
	err = service.UpdateIssueByIdService(id, update_data)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"message": "Placement issue not found"})
			return
		}
		c.JSON(500, gin.H{"message": "Internal Error"})
		return
	}

	c.JSON(200, gin.H{"message": "Placement Issue updated successfully"})
}

func DeleteIssue(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	err = service.DeleteIssueService(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(400, gin.H{"message": "issue not found"})
			return
		}
		c.JSON(500, gin.H{"message": "Placement Issues not able to delete"})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
