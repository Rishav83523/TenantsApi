package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (app *application) routes() http.Handler { 
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.GET("/tenants", app.getTenants)
		v1.POST("/tenants", app.createTenant)
		v1.GET("/tenants/:id", app.getTenantByID)
		v1.POST("/tenants/:id/projects", app.createProject)
		v1.GET("/tenants/:id/projects", app.getProjectsByTenantID)
		v1.POST("/projects/:id/services", app.createService)
		v1.GET("/projects/:id/services", app.getServices)
		v1.GET("/services/:id/deployments", app.getDeploymentsByServiceID)
		v1.POST("/services/:id/deployments", app.createDeployment)

	}
	return g
}