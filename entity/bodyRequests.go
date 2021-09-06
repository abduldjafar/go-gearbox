package entity

type UserRegisterRequest struct {
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
