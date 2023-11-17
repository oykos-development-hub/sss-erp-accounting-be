CREATE TABLE IF NOT EXISTS movement_articles (
    id serial PRIMARY KEY,
    article_id INTEGER,
    movement_id INTEGER REFERENCES movements(id) ON DELETE CASCADE,
    amount INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
