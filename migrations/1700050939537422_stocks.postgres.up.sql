CREATE TABLE IF NOT EXISTS stocks (
    id serial PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    amount INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
