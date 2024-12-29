package types

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type LoginUserReqBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int    `json:"number"`
}

type AuthUser struct {
	ID        int
	Number    int
	ExpiredAt time.Time
	jwt.RegisteredClaims
}

type SignupUserReqBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
