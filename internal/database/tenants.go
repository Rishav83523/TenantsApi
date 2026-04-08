package database

import (
	"context"
	"database/sql"
	"time"
)


type TenatModel struct { 
	DB *sql.DB
}

type Tenat struct { 
	ID int
	Name string
	Email string
	CreatedAt string
	UpdatedAt string
}


func (m *TenatModel) GetAll() ([]*Tenat, error) { 
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM tenats"
    
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tenats := []*Tenat{}
	for rows.Next() {
		var tenat Tenat
		err := rows.Scan(&tenat.ID, &tenat.Name, &tenat.Email, &tenat.CreatedAt, &tenat.UpdatedAt)	
		if err != nil {
			return nil, err
		}	
		tenats = append(tenats, &tenat)
	}

	return tenats, nil


}


func (m *TenatModel) Create(name, email string)  error { 
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := "INSERT INTO tenats (name, email, created_at, updated_at) VALUES (?, ?, ?, ?)"

	_, err := m.DB.ExecContext(ctx, query, name, email, time.Now(), time.Now())
	return err
}

func (m *TenatModel) GetByID(id int) (*Tenat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, name, email, created_at, updated_at FROM tenats WHERE id = ?"

	var tenat Tenat
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&tenat.ID,
		&tenat.Name,
		&tenat.Email,
		&tenat.CreatedAt,
		&tenat.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &tenat, nil
}



