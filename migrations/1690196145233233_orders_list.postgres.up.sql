CREATE TABLE IF NOT EXISTS order_lists (
    id serial PRIMARY KEY,
    date_order DATE NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    public_procurement_id INTEGER NOT NULL,
    supplier_id INTEGER,
    status TEXT NOT NULL,
    date_system DATE,
    invoice_date DATE,
    invoice_number INTEGER,
    organization_unit_id INTEGER NOT NULL,
    office_id INTEGER,
    recipient_user_id INTEGER NOT NULL,
    description_recipient TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
