package domain

type UserSituation struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	SituationID int    `json:"situation_id"`
	AssignedAt  string `json:"assigned_at"`
}
