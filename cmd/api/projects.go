package main

import (
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getProjectsByTenantID godoc
// @Summary List projects by tenant
// @Description Get all projects for a tenant id
// @Tags Projects
// @Produce json
// @Param id path int true "Tenant ID"
// @Success 200 {array} database.Project
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /tenants/{id}/projects [get]
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

// createProject godoc
// @Summary Create project
// @Description Create a new project for a tenant
// @Tags Projects
// @Accept json
// @Produce json
// @Param id path int true "Tenant ID"
// @Param payload body createProjectRequest true "Project payload"
// @Success 201 {object} database.Project
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /tenants/{id}/projects [post]
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