package domain

type Notification struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	SituationID int    `json:"situation_id"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	SentAt      string `json:"sent_at"`
}
