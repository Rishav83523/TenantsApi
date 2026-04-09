package main

import (
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) getDeploymentsByServiceID(c *gin.Context) { 
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {	
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service id"})
		return
	}

	deployments, err := app.models.Deployments.GetByServiceID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve deployments"})
		return
	}

	if len(deployments) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No deployments found for this service"})
		return
	}

	c.JSON(http.StatusOK, deployments)
}


func (app *application) createDeployment(c *gin.Context) { 
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service id"})
		return
	}

	var deployment database.Deployment
	if err := c.ShouldBindJSON(&deployment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}


	createdDeployment, err := app.models.Deployments.Create(id, deployment.Version, deployment.Status, deployment.Environment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create deployment"})
		return
	}

	c.JSON(http.StatusCreated, createdDeployment)
}