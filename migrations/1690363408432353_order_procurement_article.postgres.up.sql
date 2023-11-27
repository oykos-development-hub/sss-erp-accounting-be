CREATE TABLE IF NOT EXISTS order_procurement_articles (
    id serial PRIMARY KEY,
    order_id INTEGER NOT NULL,
    article_id INTEGER,
    title TEXT,
    description TEXT,
    amount INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES order_lists (id) ON UPDATE CASCADE ON DELETE CASCADE
);
