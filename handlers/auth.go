package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ladmakhi81/gobanks/entities"
	"github.com/ladmakhi81/gobanks/repositories"
	"github.com/ladmakhi81/gobanks/types"
	"github.com/ladmakhi81/gobanks/utils"
)

type AuthHandler struct {
	SessionRepo repositories.SessionRepository
	AccountRepo repositories.AccountRepository
	TokenUtil   utils.TokenUtil
}

func (authHandler AuthHandler) Signup(w http.ResponseWriter, r *http.Request) error {
	reqBody := new(types.SignupUserReqBody)
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		return err
	}
	account := &entities.Account{
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
	}
	if err := authHandler.AccountRepo.CreateAccount(account); err != nil {
		return err
	}
	token, err := authHandler.TokenUtil.GenerateJwtToken(account)
	if err != nil {
		return err
	}
	session := entities.NewSession(account.ID, token)
	if err := authHandler.SessionRepo.CreateSession(session); err != nil {
		return err
	}
	utils.JsonRes(w, http.StatusOK, map[string]any{"accessToken": token, "accountID": account.ID})
	return nil
}

func (authHandler AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	reqBody := new(types.LoginUserReqBody)
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		return err
	}
	defer r.Body.Close()
	account, err := authHandler.AccountRepo.GetAccountByNumber(reqBody.Number)
	if err != nil {
		return err
	}
	if !(account.FirstName == reqBody.FirstName && account.LastName == reqBody.LastName) {
		return errors.New("user not found")
	}
	token, err := authHandler.TokenUtil.GenerateJwtToken(account)
	if err != nil {
		return err
	}
	session := entities.NewSession(account.ID, token)
	if err := authHandler.SessionRepo.CreateSession(session); err != nil {
		return err
	}
	utils.JsonRes(w, http.StatusOK, map[string]any{"accessToken": token, "accountID": account.ID})
	return nil
}

func (authHandler AuthHandler) Logout(w http.ResponseWriter, r *http.Request) error {
	authUserId := r.Context().Value("Auth").(*types.AuthUser).ID
	err := authHandler.SessionRepo.DeleteSessionByUserId(authUserId)
	if err != nil {
		return err
	}
	utils.JsonRes(w, http.StatusOK, map[string]string{"message": "logout successfully"})
	return nil
}

func (authHandler AuthHandler) ProfileAccount(w http.ResponseWriter, r *http.Request) error {
	authUserId := r.Context().Value("Auth").(*types.AuthUser).ID
	account, err := authHandler.AccountRepo.GetAccountByID(authUserId)
	if err != nil {
		return err
	}
	utils.JsonRes(w, http.StatusOK, account)
	return nil
}
