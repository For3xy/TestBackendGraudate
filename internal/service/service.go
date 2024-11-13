package service

import (
	"testbackendGraudate/domain"
	"testbackendGraudate/internal/repository"
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

type Service struct {
	Authorization
	Camera
	Situation
	IncidentType
	Notification
	UserSituation
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}
