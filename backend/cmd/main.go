package main

import (
	"fmt"

	"pcb-backend/backend/database"
	"pcb-backend/backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.POST("/projects", handlers.ProjectCreation)
	router.GET("/projects", handlers.GetAllProjects)
	router.GET("/projects/:id", handlers.GetProjectById)
	router.PUT("/projects/:id", handlers.UpdateProject)
	router.DELETE("/projects/:id", handlers.DeleteProject)

	router.POST("/components", handlers.CreateComponent)
	router.GET("/components", handlers.GetComponent)
	router.GET("/components/:id", handlers.GetByIdComponent)
	router.PUT("/components/:id", handlers.UpdateComponent)
	router.DELETE("/components/:id", handlers.DeleteComponent)

	router.POST("/issues", handlers.CreateIssues)
	router.GET("/issues", handlers.GetAllIssues)
	router.GET("/issues/:id", handlers.GetIssuesById)
	router.PUT("/issues/:id", handlers.UpdateIssueById)
	router.DELETE("/issues/:id", handlers.DeleteIssue)

	fmt.Println("server start to run.....")
	router.Run(":8080")
}
