-- Migration: Create products table
-- Description: Membuat tabel products dengan foreign key ke categories

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL DEFAULT 0,
    stock INTEGER NOT NULL DEFAULT 0,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL
);

-- Insert sample products
INSERT INTO products (name, price, stock, category_id) VALUES 
    ('Nasi Goreng', 15000, 50, 1),
    ('Es Teh', 5000, 100, 1),
    ('Mouse Wireless', 75000, 20, 2),
    ('Kaos Polos', 50000, 30, 3)
ON CONFLICT DO NOTHING;