package user_pg

import (
	"context"
	"database/sql"
	"log"

	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/user_repo"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
)

type userPG struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) user_repo.Repository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) Create(ctx context.Context, user entity.User) errs.MessageErr {

	if _, err := u.db.ExecContext(ctx, CREATE_USER_QUERY, user.Username, user.Password); err != nil {
		log.Printf("db create user: %s\n", err.Error())
		return errs.NewInternalServerError()
	}

	return nil
}
func (u *userPG) GetByUsername(ctx context.Context, username string) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	if err := u.db.QueryRowContext(ctx, GET_BY_USERNAME_QUERY, username).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Roles); err != nil {
		log.Printf("db get by username: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user was not found")
		}

		return nil, errs.NewInternalServerError()
	}

	return &user, nil
}
func (u *userPG) GetById(ctx context.Context, id int) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	if err := u.db.QueryRowContext(ctx, GET_BY_ID_QUERY, id).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Roles); err != nil {
		log.Printf("db get by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user was not found")
		}

		return nil, errs.NewInternalServerError()
	}

	return &user, nil
}
