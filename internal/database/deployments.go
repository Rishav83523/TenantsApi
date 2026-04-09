package database

import (
	"database/sql"
)

type DeploymentModel struct {
	DB *sql.DB
}


type Deployment struct {
	ID int
ServiceID int
Version string
Status string
Environment string
DeployedAt string	
}

func (m *DeploymentModel) GetByServiceID(serviceID int) ([]*Deployment, error) {
	query := "SELECT id, service_id, version, status, environment, deployed_at FROM deployments WHERE service_id = $1 ORDER BY id"
	rows, err := m.DB.Query(query, serviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deployments := []*Deployment{}
	for rows.Next() {
		var deployment Deployment
		err := rows.Scan(&deployment.ID, &deployment.ServiceID, &deployment.Version, &deployment.Status, &deployment.Environment, &deployment.DeployedAt)
		if err != nil {
			return nil, err
		}
		deployments = append(deployments, &deployment)
	}

	return deployments, nil
}


func (m *DeploymentModel) Create(serviceID int, version string, status string, environment string) (*Deployment, error) {
	query := "INSERT INTO deployments (service_id, version, status, environment, deployed_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id, service_id, version, status, environment, deployed_at"	
	var deployment Deployment
	err := m.DB.QueryRow(query, serviceID, version, status, environment).Scan(&deployment.ID, &deployment.ServiceID, &deployment.Version, &deployment.Status, &deployment.Environment, &deployment.DeployedAt)
	if err != nil {
		return nil, err
	}
	return &deployment, nil
}