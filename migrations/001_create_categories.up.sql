-- Migration: Create categories table
-- Description: Membuat tabel categories

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

-- Insert sample categories
INSERT INTO categories (name, description) VALUES 
    ('Makanan', 'Kategori makanan dan minuman'),
    ('Elektronik', 'Kategori barang elektronik'),
    ('Pakaian', 'Kategori pakaian dan aksesoris')
ON CONFLICT DO NOTHING;