package main

import (
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getTenants godoc
// @Summary List tenants
// @Description Get all tenants
// @Tags Tenants
// @Produce json
// @Success 200 {array} database.Tenat
// @Failure 500 {object} errorResponse
// @Router /tenants [get]
func (app *application) getTenants(c *gin.Context) { 
    tenants, err := app.models.Tenants.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tenants"})
		return
	}
	c.JSON(http.StatusOK, tenants)
}

// createTenant godoc
// @Summary Create tenant
// @Description Create a new tenant
// @Tags Tenants
// @Accept json
// @Produce json
// @Param payload body createTenantRequest true "Tenant payload"
// @Success 201 {object} database.Tenat
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /tenants [post]
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

// getTenantByID godoc
// @Summary Get tenant by id
// @Description Get a tenant by tenant id
// @Tags Tenants
// @Produce json
// @Param id path int true "Tenant ID"
// @Success 200 {object} database.Tenat
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /tenants/{id} [get]
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