CREATE TABLE IF NOT EXISTS stock_order_articles (
    id serial PRIMARY KEY,
    stock_id INTEGER REFERENCES stocks(id) ON DELETE CASCADE,
    article_id INTEGER REFERENCES order_procurement_articles(id) ON DELETE CASCADE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
