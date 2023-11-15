CREATE TABLE IF NOT EXISTS movement_articles (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    movement_id INTEGER REFERENCES movements(id) ON DELETE CASCADE,
    description TEXT,
    amount INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
