CREATE TABLE IF NOT EXISTS services (
    id integer PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    type TEXT,
    language TEXT,
    project_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
)