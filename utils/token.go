package utils

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/ladmakhi81/gobanks/entities"
	"github.com/ladmakhi81/gobanks/repositories"
	"github.com/ladmakhi81/gobanks/types"
)

type TokenUtil struct {
	AccountRepo repositories.AccountRepository
	SessionRepo repositories.SessionRepository
}

func (TokenUtil) GenerateJwtToken(account *entities.Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":        account.ID,
		"Number":    account.Number,
		"ExpiredAt": time.Now().Add(time.Hour * 1),
	})
	tokenString, err := token.SignedString([]byte("xxx"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (tokenUtil TokenUtil) ValidateToken(token string) (*types.AuthUser, error) {
	decodedToken, err := jwt.ParseWithClaims(token, &types.AuthUser{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("xxx"), nil
	})
	if err != nil {

		return nil, err
	}
	if !decodedToken.Valid {
		return nil, errors.New("permission defined")
	}
	claims := decodedToken.Claims.(*types.AuthUser)

	session, err := tokenUtil.SessionRepo.GetSessionByToken(decodedToken.Raw)

	if err != nil || session == nil {
		return nil, errors.New("permission defined")
	}

	return claims, err
}
