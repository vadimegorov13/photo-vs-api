package models

type RegisterUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
