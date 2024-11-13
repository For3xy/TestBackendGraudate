package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testbackendGraudate/domain"
	"time"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (p *AuthPostgres) CreateUser(user domain.User) (int, error) {
	var id int
	user.CreatedAt = time.Now()
	query := fmt.Sprintf("INSERT INTO %s (username, role, email, phone_number, password, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", usersTable)
	row := p.db.QueryRow(query, user.Username, user.Role, user.Email, user.Phone, user.Password, user.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
