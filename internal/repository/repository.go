package repository

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

type Repository struct {
	Authorization
	Camera
	Situation
	IncidentType
	Notification
	UserSituation
}

func NewRepository() *Repository {
	return &Repository{}
}
