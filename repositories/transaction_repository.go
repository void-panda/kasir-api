package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/model"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []model.CheckoutItem) (*model.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]model.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow(
			"SELECT name, price, stock FROM products WHERE id = $1",
			item.ProductID,
		).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		if stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s (id: %d)", productName, item.ProductID)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec(
			"UPDATE products SET stock = stock - $1 WHERE id = $2",
			item.Quantity,
			item.ProductID,
		)
		if err != nil {
			return nil, err
		}

		details = append(details, model.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	var createdAt time.Time
	err = tx.QueryRow(
		"INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id, created_at	",
		totalAmount,
	).Scan(&transactionID, &createdAt)
	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID
		err = tx.QueryRow(
			"INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING id",
			transactionID,
			details[i].ProductID,
			details[i].Quantity,
			details[i].Subtotal,
		).Scan(&details[i].ID)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &model.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		CreatedAt:   createdAt,
		Details:     details,
	}, nil
}

func (repo *TransactionRepository) GetTodaySummary() (*model.SalesSummary, error) {
	summary := &model.SalesSummary{}

	// Total Revenue & Total Transaksi
	queryTotal := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(id) 
		FROM transactions 
		WHERE created_at::date = CURRENT_DATE`
	err := repo.db.QueryRow(queryTotal).Scan(&summary.TotalRevenue, &summary.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// Produk Terlaris
	queryTopProduct := `
		SELECT p.name, SUM(td.quantity) as total_qty
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at::date = CURRENT_DATE
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1`
	err = repo.db.QueryRow(queryTopProduct).Scan(&summary.ProdukTerlaris.Nama, &summary.ProdukTerlaris.QtyTerjual)
	if err == sql.ErrNoRows {
		summary.ProdukTerlaris.Nama = "-"
		summary.ProdukTerlaris.QtyTerjual = 0
	} else if err != nil {
		return nil, err
	}

	return summary, nil
}
