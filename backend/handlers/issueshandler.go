package handlers

import (
	"errors"
	"fmt"
	"pcb-backend/backend/models"
	"pcb-backend/backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateIssues(c *gin.Context) {
	var req models.CreateIssueRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"message": "validation failed", "error": err.Error()})
		return
	}

	err = service.CreateIssuesService(req.ToModel())
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInput):
			c.JSON(400, gin.H{"message": "issue description cannot be blank"})
		case errors.Is(err, service.ErrComponentNotFound):
			c.JSON(404, gin.H{"message": "component not found"})
		default:
			fmt.Println("create issue error:", err)
			c.JSON(500, gin.H{"message": "Failed to create issue"})
		}
		return
	}
	c.JSON(200, gin.H{"message": "Placement Issue created"})
}

func GetAllIssues(c *gin.Context) {
	issues, err := service.GetAllIssuesService()
	if err != nil {
		fmt.Println("get issues error:", err)
		c.JSON(500, gin.H{"message": "Failed to get the issues"})
		return
	}
	c.JSON(200, issues)
}

func GetIssuesById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Id"})
		return
	}

	issue, err := service.GetIssuesByIdService(id)
	if err != nil {
		if errors.Is(err, service.ErrIssueNotFound) {
			c.JSON(404, gin.H{"message": "Placement issue not found"})
			return
		}
		fmt.Println("get issue error:", err)
		c.JSON(500, gin.H{"message": "Failed to get the issue"})
		return
	}
	c.JSON(200, issue)
}

func UpdateIssueById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Id"})
		return
	}

	var req models.UpdateIssueRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"message": "validation failed", "error": err.Error()})
		return
	}

	err = service.UpdateIssueByIdService(id, models.Status(req.Status))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrIssueNotFound):
			c.JSON(404, gin.H{"message": "Placement issue not found"})
		case errors.Is(err, service.ErrInvalidTransition):
			c.JSON(409, gin.H{"message": "this status change is not allowed"})
		default:
			fmt.Println("update issue error:", err)
			c.JSON(500, gin.H{"message": "Failed to update issue"})
		}
		return
	}
	c.JSON(200, gin.H{"message": "Placement Issue updated successfully"})
}

func DeleteIssue(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Id"})
		return
	}

	err = service.DeleteIssueService(id)
	if err != nil {
		if errors.Is(err, service.ErrIssueNotFound) {
			c.JSON(404, gin.H{"message": "Placement issue not found"})
			return
		}
		fmt.Println("delete issue error:", err)
		c.JSON(500, gin.H{"message": "Failed to delete issue"})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted successfully"})
}
