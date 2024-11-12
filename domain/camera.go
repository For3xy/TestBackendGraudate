package domain

type Camera struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
}
