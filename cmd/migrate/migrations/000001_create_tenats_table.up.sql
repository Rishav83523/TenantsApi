CREATE TABLE IF NOT EXISTS tenats (
    id integer PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)