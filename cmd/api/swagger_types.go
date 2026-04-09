package main

type errorResponse struct {
	Error string `json:"error" example:"Invalid request body"`
}

type createTenantRequest struct {
	Name  string `json:"name" example:"Acme Corp"`
	Email string `json:"email" example:"admin@acme.com"`
}

type createProjectRequest struct {
	Name        string `json:"name" example:"Tenant Portal"`
	Description string `json:"description" example:"Customer-facing portal"`
}

type createServiceRequest struct {
	Name     string `json:"name" example:"Auth Service"`
	Type     string `json:"type" example:"backend"`
	Language string `json:"language" example:"go"`
}

type createDeploymentRequest struct {
	Version     string `json:"version" example:"1.0.0"`
	Status      string `json:"status" example:"in test"`
	Environment string `json:"environment" example:"aws"`
}