package internal_jwt

import (
	"errors"
	"strings"

	"github.com/adityatresnobudi/go-restapi-gin/pkg/constants"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
	"github.com/golang-jwt/jwt"
)

type internalJwtImpl struct {
}

type InternalJwt interface {
	GenerateToken(jwtClaim jwt.MapClaims, secretKey string) string
	ValidateBearerToken(bearerToken string, secretKey string) (jwt.MapClaims, errs.MessageErr)
}

func NewInternalJwt() InternalJwt {
	return &internalJwtImpl{}
}

func (ij *internalJwtImpl) signToken(claims jwt.MapClaims, secretKey string) string {
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func (ij *internalJwtImpl) GenerateToken(jwtClaim jwt.MapClaims, secretKey string) string {
	return ij.signToken(jwtClaim, secretKey)
}

func (ij *internalJwtImpl) parseToken(stringToken string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(constants.INVALID_TOKEN_ERROR_MESSAGE)
		}

		return []byte(secretKey), nil
	})

	if err != nil {

		var vErr *jwt.ValidationError

		if errors.As(err, &vErr) {
			if vErr.Errors == jwt.ValidationErrorExpired {
				return nil, errors.New("expired token")
			}
		}

		return nil, errors.New(constants.INVALID_TOKEN_ERROR_MESSAGE)
	}

	return token, nil
}

func (ij *internalJwtImpl) ValidateBearerToken(
	bearerToken string,
	secretKey string,
) (jwt.MapClaims, errs.MessageErr) {

	if bearer := strings.HasPrefix(bearerToken, constants.BEARER); !bearer {
		return nil, errs.NewUnauthenticatedError(constants.INVALID_TOKEN_ERROR_MESSAGE)
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return nil, errs.NewUnauthenticatedError(constants.INVALID_TOKEN_ERROR_MESSAGE)
	}

	token, err := ij.parseToken(splitToken[1], secretKey)

	if err != nil {
		return nil, errs.NewUnauthenticatedError(constants.INVALID_TOKEN_ERROR_MESSAGE)
	}

	var mapClaims jwt.MapClaims

	if v, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {

		return nil, errs.NewUnauthenticatedError(constants.INVALID_TOKEN_ERROR_MESSAGE)
	} else {
		mapClaims = v

	}

	return mapClaims, nil
}
