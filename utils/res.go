package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ladmakhi81/gobanks/types"
)

func JsonRes(w http.ResponseWriter, status int, value any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{"data": value})
}

func UnAuthJsonErr(w http.ResponseWriter) {
	JsonRes(w, http.StatusUnauthorized, map[string]any{"message": "user unauthorized"})
}

func ApiHandler(fn types.ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			JsonRes(w, http.StatusBadRequest, err.Error())
			return
		}
	}
}
