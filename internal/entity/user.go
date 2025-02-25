package entity

import (
	"time"

	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int
	Username  string
	Password  string
	Roles     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) ToUserResponseDTO() *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		Id:        u.Id,
		Username:  u.Username,
		Roles: u.Roles,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) HashPassword() errs.MessageErr {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)

	if err != nil {
		return errs.NewInternalServerError()
	}

	u.Password = string(b)

	return nil
}

func (u *User) Compare(password string) errs.MessageErr {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(password),
	); err != nil {
		return errs.NewUnauthorizedError("invalid password")
	}

	return nil
}

func (u *User) NewClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":       u.Id,
		"username": u.Username,
		"expr":     time.Now().Add(24 * time.Hour),
	}
}
