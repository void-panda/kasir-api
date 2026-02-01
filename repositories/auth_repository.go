package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (repo *AuthRepository) CreateUser(u *model.User) error {
	query := "INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id"
	return repo.db.QueryRow(query, u.Name, u.Email, u.Password).Scan(&u.ID)
}

func (repo *AuthRepository) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	query := "SELECT id, name, email, password FROM users WHERE email = $1"

	err := repo.db.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User with the given Email is not found")
		}
		return nil, err
	}

	return &u, nil
}
