CREATE TABLE IF NOT EXISTS stocks (
    id serial PRIMARY KEY,
    title TEXT,
    description TEXT,
    year TEXT,
    amount INTEGER NOT NULL,
    organization_unit_id INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
