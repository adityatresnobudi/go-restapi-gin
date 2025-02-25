package service

import (
	"context"

	"github.com/adityatresnobudi/go-restapi-gin/config"
	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/user_repo"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/internal_jwt"
)

type UserService interface {
	Create(ctx context.Context, payload dto.RegisterRequestDTO) (*dto.CommonBaseResponseDTO, errs.MessageErr)
	Login(ctx context.Context, payload dto.LoginRequestDTO) (*dto.LoginResponseDTO, errs.MessageErr)
	GetById(ctx context.Context, id int) (*entity.User, errs.MessageErr)
}

type userServiceIMPL struct {
	userRepo    user_repo.Repository
	internalJWT internal_jwt.InternalJwt
	cfg         config.Config
}

func NewUserService(
	userRepo user_repo.Repository,
	internalJWT internal_jwt.InternalJwt,
	cfg config.Config,
) UserService {
	return &userServiceIMPL{
		userRepo:    userRepo,
		internalJWT: internalJWT,
		cfg:         cfg,
	}
}

func (u *userServiceIMPL) Create(ctx context.Context, payload dto.RegisterRequestDTO) (*dto.CommonBaseResponseDTO, errs.MessageErr) {
	user := entity.User{
		Username: payload.Username,
		Password: payload.Password,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	result := dto.CommonBaseResponseDTO{Message: "New user successfully registered"}

	return &result, nil
}
func (u *userServiceIMPL) Login(ctx context.Context, payload dto.LoginRequestDTO) (*dto.LoginResponseDTO, errs.MessageErr) {

	user, err := u.userRepo.GetByUsername(ctx, payload.Username)

	if err != nil {
		return nil, err
	}

	if err = user.Compare(payload.Password); err != nil {
		return nil, err
	}

	claim := user.NewClaim()

	token := u.internalJWT.GenerateToken(claim, u.cfg.Jwt.SecretKey)

	result := dto.LoginResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "OK"},
		Data: dto.LoginResponseDataDTO{
			Token: token,
		},
	}

	return &result, nil
}

func (u *userServiceIMPL) GetById(ctx context.Context, id int) (*entity.User, errs.MessageErr) {
	return u.userRepo.GetById(ctx, id)
}
