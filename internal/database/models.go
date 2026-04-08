package database

import "database/sql"


type Models struct {
	Tenants TenatModel
	Services ServiceModel
	Projects ProjectModel
	Deployments DeploymentModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Tenants: TenatModel{DB: db},
		Services: ServiceModel{DB: db},	
		Projects: ProjectModel{DB: db},
		Deployments: DeploymentModel{DB: db},
	}
}

