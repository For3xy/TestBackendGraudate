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

func (p *AuthPostgres) LoginUser(email, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", usersTable)
	err := p.db.Get(&user, query, email, password)

	user.LastLogin = time.Now()
	_, errUpdate := p.db.Exec("UPDATE users SET last_login = $1 WHERE id = $2", user.LastLogin, user.ID)
	if errUpdate != nil {
		return user, errUpdate
	}

	return user, err
}
