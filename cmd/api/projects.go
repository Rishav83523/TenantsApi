package main

import (
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) getProjectsByTenantID(c *gin.Context) { 
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant id"})
		return
	}

	projects, err := app.models.Projects.GetByTenantID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects"})
		return
	}

	if len(projects) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No projects found for this tenant"})
		return
	}

	c.JSON(http.StatusOK, projects)

}


func (app *application) createProject(c *gin.Context) { 
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant id"})
		return
	}

	var project database.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdProject, err := app.models.Projects.Create(id, project.Name, project.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, createdProject)



}