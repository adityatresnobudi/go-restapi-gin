package dto

import "time"

type UserResponseDTO struct {
	Id        int       `json:"id" example:"1"`
	Username  string    `json:"username" example:"adit"`
	Roles         string    `json:"roles" example:"user"`
	CreatedAt time.Time `json:"created_at" example:"2025-02-22T15:11:19.25616+07:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-02-22T15:11:19.25616+07:00"`
}

type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseDataDTO struct {
	Token string `json:"token"`
}

type LoginResponseDTO struct {
	CommonBaseResponseDTO
	Data LoginResponseDataDTO `json:"data"`
}
