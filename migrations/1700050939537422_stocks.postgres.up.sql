CREATE TABLE IF NOT EXISTS stocks (
    id serial PRIMARY KEY,
    article_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    organization_unit_id INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
