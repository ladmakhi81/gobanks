package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ladmakhi81/gobanks/contracts"
	"github.com/ladmakhi81/gobanks/types"
	"github.com/ladmakhi81/gobanks/utils"
)

type AccountHandler struct {
	Repo contracts.AccountRepository
}

func (accHandler AccountHandler) DeleteAccountHandler(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id, parseErr := strconv.Atoi(params["id"])
	if parseErr != nil {
		return parseErr
	}
	err := accHandler.Repo.DeleteAccount(id)
	if err != nil {
		return err
	}
	utils.JsonRes(w, http.StatusOK, map[string]string{"message": "delete successfully ..."})
	return nil
}

func (accHandler AccountHandler) GetAccountByIdHandler(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id, parseErr := strconv.Atoi(params["id"])
	if parseErr != nil {
		return parseErr
	}
	acc, err := accHandler.Repo.GetAccountByID(id)
	if err != nil {
		return err
	}
	utils.JsonRes(w, http.StatusOK, acc)
	return nil
}

func (accHandler AccountHandler) GetAccountsHandler(w http.ResponseWriter, r *http.Request) error {
	accounts, err := accHandler.Repo.GetAccounts()
	if err != nil {
		return err
	}
	utils.JsonRes(w, http.StatusOK, accounts)
	return nil
}

func (accHandler AccountHandler) TransferMoneyHandler(w http.ResponseWriter, r *http.Request) error {
	reqBody := new(types.TransferMoneyReqBody)
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		return err
	}
	defer r.Body.Close()
	receiverAccount, receiverErr := accHandler.Repo.GetAccountByNumber(reqBody.ToAccount)
	if receiverErr != nil {
		return receiverErr
	}
	authUserId := r.Context().Value("Auth").(*types.AuthUser).ID
	senderAccount, senderErr := accHandler.Repo.GetAccountByID(authUserId)
	if senderErr != nil {
		return senderErr
	}
	if senderAccount.Balance < reqBody.Amount {
		return errors.New("not enough credit")
	}
	if err := accHandler.Repo.WithDrawCredit(senderAccount.Number, reqBody.Amount); err != nil {
		return err
	}
	if err := accHandler.Repo.DepositCredit(receiverAccount.Number, reqBody.Amount); err != nil {
		return err
	}
	utils.JsonRes(w, http.StatusOK, map[string]string{"message": "transfer successfully ..."})
	return nil
}
