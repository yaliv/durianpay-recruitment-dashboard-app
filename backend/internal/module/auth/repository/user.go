package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
)

type UserRepository interface {
	GetUserByEmail(email string) (*entity.User, error)
}

type User struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *User {
	return &User{db: db}
}

func (r *User) GetUserByEmail(email string) (*entity.User, error) {
	var u entity.User

	q := `SELECT id, email, password_hash, role FROM users WHERE email = ?`
	if err := r.db.Get(&u, q, email); err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrorNotFound("user not found")
		}
		return nil, entity.WrapError(err, entity.ErrorCodeInternal, "db error")
	}

	return &u, nil
}
