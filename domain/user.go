package domain

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LastLogin string `json:"last_login"`
	CreatedAt string `json:"created_at"`
}
