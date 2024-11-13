package repository

import (
	"github.com/jmoiron/sqlx"
	"testbackendGraudate/domain"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
}

type Camera interface {
}

type Situation interface {
}

type IncidentType interface {
}

type Notification interface {
}

type UserSituation interface {
}

type Repository struct {
	Authorization
	Camera
	Situation
	IncidentType
	Notification
	UserSituation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
