package service

import (
	"strings"

	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
)

func (a *accountServiceIMPL) createValidator(payload dto.CreateAccountRequestDTO) errs.MessageErr {
	errArr := make([]errs.MessageErr, 0)

	if strings.TrimSpace(payload.AccountNumber) == "" {
		errArr = append(errArr, errs.NewBadRequest("account number cannot be empty"))
	}

	if payload.Balance < 1 {
		errArr = append(errArr, errs.NewBadRequest("amount cannot be less than 1"))
	}

	if len(errArr) > 0 {
		msgArr := make([]string, 0)
		for _, value := range errArr {
			msgArr = append(msgArr, value.Error())
		}
		errMsg := strings.Join(msgArr, ", ")
		err := errs.ErrorData{
			ErrCode:       errArr[0].Code(),
			ErrStatusCode: errArr[0].StatusCode(),
			ErrMessage:    errMsg,
		}
		return &err
	}

	return nil
}

func (a *accountServiceIMPL) updateValidator(payload dto.UpdateAccountRequestDTO) errs.MessageErr {
	errArr := make([]errs.MessageErr, 0)

	if payload.Balance < 1 {
		errArr = append(errArr, errs.NewBadRequest("amount cannot be less than 1"))
	}

	if len(errArr) > 0 {
		msgArr := make([]string, 0)
		for _, value := range errArr {
			msgArr = append(msgArr, value.Error())
		}
		errMsg := strings.Join(msgArr, ", ")
		err := errs.ErrorData{
			ErrCode:       errArr[0].Code(),
			ErrStatusCode: errArr[0].StatusCode(),
			ErrMessage:    errMsg,
		}
		return &err
	}

	return nil
}
