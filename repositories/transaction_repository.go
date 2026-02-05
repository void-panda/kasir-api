package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/model"
	"time"

	"github.com/lib/pq"
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

	// 1. Dapatkan semua ID produk untuk batch select
	productIDs := make([]int, len(items))
	itemMap := make(map[int]int)
	for i, item := range items {
		productIDs[i] = item.ProductID
		itemMap[item.ProductID] = item.Quantity
	}

	// 2. Batch SELECT produk
	rows, err := tx.Query(
		"SELECT id, name, price, stock FROM products WHERE id = ANY($1)",
		pq.Array(productIDs),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make(map[int]struct {
		Name  string
		Price int
		Stock int
	})
	for rows.Next() {
		var id int
		var p struct {
			Name  string
			Price int
			Stock int
		}
		if err := rows.Scan(&id, &p.Name, &p.Price, &p.Stock); err != nil {
			return nil, err
		}
		products[id] = p
	}

	// 3. Validasi stok dan hitung total
	totalAmount := 0
	details := make([]model.TransactionDetail, 0)
	for _, item := range items {
		p, ok := products[item.ProductID]
		if !ok {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}

		if p.Stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s (id: %d)", p.Name, item.ProductID)
		}

		subtotal := p.Price * item.Quantity
		totalAmount += subtotal

		// Tetap update stok per baris untuk locking ROW (mencegah double sell)
		// Namun bisa juga dioptimasi jika perlu. Dalam case POS, ini biasanya ok.
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
			ProductName: p.Name,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	// 4. INSERT transaction
	var transactionID int
	var createdAt time.Time
	err = tx.QueryRow(
		"INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id, created_at",
		totalAmount,
	).Scan(&transactionID, &createdAt)
	if err != nil {
		return nil, err
	}

	// 5. Batch INSERT transaction details
	// Kita bisa gunakan satu query dengan banyak VALUES
	query := "INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES "
	values := []interface{}{}
	for i, d := range details {
		details[i].TransactionID = transactionID
		n := i * 4
		query += fmt.Sprintf("($%d, $%d, $%d, $%d),", n+1, n+2, n+3, n+4)
		values = append(values, transactionID, d.ProductID, d.Quantity, d.Subtotal)
	}
	query = query[:len(query)-1] // Remove trailing comma
	query += " RETURNING id"

	rows, err = tx.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	detailIdx := 0
	for rows.Next() {
		if err := rows.Scan(&details[detailIdx].ID); err != nil {
			return nil, err
		}
		detailIdx++
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
	err = repo.db.QueryRow(queryTopProduct).
		Scan(&summary.ProdukTerlaris.Nama, &summary.ProdukTerlaris.QtyTerjual)
	if err == sql.ErrNoRows {
		summary.ProdukTerlaris.Nama = "-"
		summary.ProdukTerlaris.QtyTerjual = 0
	} else if err != nil {
		return nil, err
	}

	return summary, nil
}
