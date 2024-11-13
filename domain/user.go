package domain

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Phone     string    `json:"phone_number" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
}
