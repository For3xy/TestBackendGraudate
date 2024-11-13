package domain

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	LastLogin time.Time `json:"last_login" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}
