package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/model"
)

type UserRepository struct {
	db *sql.DB
}

// struct method for : UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// === Repo Functions ===
func (repo *UserRepository) GetAll() ([]model.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]model.User, 0)
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, err
}

func (repo *UserRepository) GetByID(id int) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"

	var u model.User
	err := repo.db.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("User tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &u, err
}

func (repo *UserRepository) Update(user *model.User) error {
	query := "UPDATE users SET name = $1, email = $2 where id = $3"
	result, err := repo.db.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (repo *UserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return err
}
