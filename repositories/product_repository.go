package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(nameFilter string) ([]model.Product, error) {
	query := "SELECT id, name, price, stock, category_id FROM products"
	args := []interface{}{}

	if nameFilter != "" {
		query += " WHERE name ILIKE $1"
		args = append(args, "%"+nameFilter+"%")
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *model.Product) error {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	return err
}

// GetByID - ambil produk by ID
func (repo *ProductRepository) GetByID(id int) (*model.Product, error) {
	query := "SELECT id, name, price, stock, category_id FROM products WHERE id = $1"

	var p model.Product
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// GetByIDWithCategory - ambil produk by ID dengan JOIN ke categories
func (repo *ProductRepository) GetByIDWithCategory(id int) (*model.ProductWithCategory, error) {
	query := `
		SELECT p.id, p.name, p.price, p.stock, p.category_id, c.name 
		FROM products p 
		LEFT JOIN categories c ON p.category_id = c.id 
		WHERE p.id = $1`

	var p model.ProductWithCategory
	var categoryName sql.NullString
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &categoryName)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	if categoryName.Valid {
		p.CategoryName = categoryName.String
	}

	return &p, nil
}

func (repo *ProductRepository) Update(product *model.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
