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

func (m *ProjectModel) GetByTenantID(tenantID int) ([]*Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, name, description, tenant_id, created_at, updated_at FROM projects WHERE tenant_id = $1 ORDER BY id"

	rows, err := m.DB.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := []*Project{}
	for rows.Next() {
		var project Project
		err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.TenantID, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}


func (m *ProjectModel) Create(tenantID int, name string, description string) (*Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
     

	query := "INSERT INTO projects (name, description, tenant_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, description, tenant_id, created_at, updated_at"

	var project Project
	err := m.DB.QueryRowContext(ctx, query, name, description, tenantID, time.Now(), time.Now()).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.TenantID,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &project, nil
}