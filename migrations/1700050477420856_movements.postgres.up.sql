CREATE TABLE IF NOT EXISTS movements (
    id serial PRIMARY KEY,
    date_order DATE NOT NULL,
    organization_unit_id INTEGER,
    office_id INTEGER,
    recipient_user_id INTEGER,
    description TEXT,
    file_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
