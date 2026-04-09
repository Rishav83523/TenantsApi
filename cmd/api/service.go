package main

import (
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)



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