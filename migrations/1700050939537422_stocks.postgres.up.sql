CREATE TABLE IF NOT EXISTS stocks (
    id serial PRIMARY KEY,
    title TEXT,
    description TEXT,
    year TEXT,
    exception BOOLEAN default false,
    amount INTEGER NOT NULL,
    vat_percentage INTEGER,
    net_price FLOAT,
    organization_unit_id INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
