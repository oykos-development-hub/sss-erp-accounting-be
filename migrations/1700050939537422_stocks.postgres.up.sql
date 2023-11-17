CREATE TABLE IF NOT EXISTS stocks (
    id serial PRIMARY KEY,
    article_id INTEGER,
    amount INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
