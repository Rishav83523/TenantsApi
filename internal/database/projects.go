package database

import (
	"context"
	"database/sql"
	"time"
)

type ProjectModel struct { 
	DB *sql.DB
}


type Project struct { 
	ID int
Name string
Description string
TenantID int
CreatedAt string
UpdatedAt string
}

func (m *ProjectModel) GetProjectByID(id int) (*Project, error) { 
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM projects WHERE id = $1"

	var project Project

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&project.ID, &project.Name, &project.Description, &project.TenantID, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &project, nil
}


func (m *ProjectModel) Create(tenantID int, name string, description string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
     

	query := "INSERT INTO projects (name, description, tenant_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := m.DB.ExecContext(ctx, query, name, description, tenantID, time.Now(), time.Now())
	return err
}