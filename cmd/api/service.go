package main

import (
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getServices godoc
// @Summary List services by project
// @Description Get all services for a project id
// @Tags Services
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {array} database.Service
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /projects/{id}/services [get]
func (app *application)  getServices(c *gin.Context) { 
    id,err := strconv.Atoi(c.Param("id"))

	if err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant id"})
		return
	}

	services, err := app.models.Services.GetServicesByProjectID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
		return
	}

	if services == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No services found for this project"})
		return
	}

	c.JSON(http.StatusOK, services)


}

// createService godoc
// @Summary Create service
// @Description Create a new service for a project
// @Tags Services
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param payload body createServiceRequest true "Service payload"
// @Success 201 {object} database.Service
// @Failure 400 {object} errorResponse
// @Router /projects/{id}/services [post]
func (app *application) createService(c *gin.Context) { 
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project id"})
		return
	}

	var service database.Service

	if err := c.ShouldBindJSON(&service); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = app.models.Services.Create(id, service.Name, service.Type, service.Language)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create service"})
		return	
	}

	c.JSON(http.StatusCreated, service)

	



}