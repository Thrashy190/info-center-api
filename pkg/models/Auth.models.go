package models

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
