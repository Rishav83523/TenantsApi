package main

import (
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (app *application) getTenants(c *gin.Context) { 
    tenants, err := app.models.Tenants.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tenants"})
		return
	}
	c.JSON(http.StatusOK, tenants)
}


func (app *application) createTenant(c *gin.Context) { 
var tenant database.Tenat

if err := c.ShouldBindJSON(&tenant); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	return
}

err := app.models.Tenants.Create(tenant.Name, tenant.Email)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tenant"})
	return

}

c.JSON(http.StatusCreated, tenant)

}

func (app *application) getTenantByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant id"})
		return
	}

	tenant, err := app.models.Tenants.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
		return
	}

	c.JSON(http.StatusOK, tenant)
}