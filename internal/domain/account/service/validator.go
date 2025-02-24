package service

import (
	"strings"

	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errors"
)

func (a *accountServiceIMPL) createValidator(payload dto.CreateAccountRequestDTO) errors.MessageErr {
	errArr := make([]errors.MessageErr, 0)

	if strings.TrimSpace(payload.AccountNumber) == "" {
		errArr = append(errArr, errors.NewBadRequest("account number cannot be empty"))
	}

	if payload.Balance < 1 {
		errArr = append(errArr, errors.NewBadRequest("amount cannot be less than 1"))
	}

	if len(errArr) > 0 {
		msgArr := make([]string, 0)
		for _, value := range errArr {
			msgArr = append(msgArr, value.Error())
		}
		errMsg := strings.Join(msgArr, ", ")
		err := errors.ErrorData{
			ErrCode:       errArr[0].Code(),
			ErrStatusCode: errArr[0].StatusCode(),
			ErrMessage:    errMsg,
		}
		return &err
	}

	return nil
}

func (a *accountServiceIMPL) updateValidator(payload dto.UpdateAccountRequestDTO) errors.MessageErr {
	errArr := make([]errors.MessageErr, 0)

	if payload.Balance < 1 {
		errArr = append(errArr, errors.NewBadRequest("amount cannot be less than 1"))
	}

	if len(errArr) > 0 {
		msgArr := make([]string, 0)
		for _, value := range errArr {
			msgArr = append(msgArr, value.Error())
		}
		errMsg := strings.Join(msgArr, ", ")
		err := errors.ErrorData{
			ErrCode:       errArr[0].Code(),
			ErrStatusCode: errArr[0].StatusCode(),
			ErrMessage:    errMsg,
		}
		return &err
	}

	return nil
}
