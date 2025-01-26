package models

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
