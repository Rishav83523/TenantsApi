CREATE TABLE IF NOT EXISTS deployments (
    id integer PRIMARY KEY AUTOINCREMENT,
    service_id INTEGER,
    version TEXT,
    status TEXT,
    environment TEXT,
    deployed_at TIMESTAMP,
    FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE CASCADE
)