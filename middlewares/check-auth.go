package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/ladmakhi81/gobanks/utils"
)

func (middleware Middlewares) CheckAuth(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationTokens := strings.Split(r.Header.Get("Authorization"), " ")

		hasBearer := strings.ToLower(authorizationTokens[0]) == "bearer"
		if !hasBearer {
			utils.UnAuthJsonErr(w)

			return
		}
		token := authorizationTokens[1]
		authUser, err := middleware.TokenUtil.ValidateToken(token)

		if err != nil {
			utils.UnAuthJsonErr(w)

			return
		}
		newContext := context.WithValue(r.Context(), "Auth", authUser)
		fn(w, r.WithContext(newContext))
	}
}
