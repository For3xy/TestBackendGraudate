package service

import "testbackendGraudate/internal/repository"

type Authorization interface {
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
	return &Service{}
}
