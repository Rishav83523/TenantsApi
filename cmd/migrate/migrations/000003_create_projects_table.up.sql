CREATE TABLE IF NOT EXISTS projects (
    id integer PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    tenant_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (tenant_id) REFERENCES tenats(id) ON DELETE CASCADE
)