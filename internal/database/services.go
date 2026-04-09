package database

import (
	"context"
	"database/sql"
	"time"
)

type ServiceModel struct { 
	DB *sql.DB
}

type Service struct { 
	ID int
Name string
Type string
Language string
ProjectID int
CreatedAt string
UpdatedAt string
}


func (m *ServiceModel) GetServicesByProjectID(projectID int) ([]*Service, error) { 
	query := "SELECT * FROM services WHERE project_id = $1"

	rows, err := m.DB.Query(query, projectID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	services := []*Service{}
	for rows.Next() {
		var service Service
		err := rows.Scan(&service.ID, &service.Name, &service.Type, &service.Language, &service.ProjectID, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			return nil, err
		}
		services = append(services, &service)
	}

	return services, nil
}


func (m *ServiceModel) Create(projectID int, name string, serviceType string, language string) error { 
   
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO services (name, type, language, project_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := m.DB.ExecContext(ctx, query, name, serviceType, language, projectID, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil

}


