package routers

import (
	"github.com/gorilla/mux"
	"github.com/ladmakhi81/gobanks/handlers"
	"github.com/ladmakhi81/gobanks/middlewares"
	"github.com/ladmakhi81/gobanks/utils"
)

type AccountRoute struct {
	Router     *mux.Router
	Handlers   handlers.AccountHandler
	Middleware middlewares.Middlewares
}

func (accRoute AccountRoute) Setup() {
	accountApiRouter := accRoute.Router.PathPrefix("/accounts").Subrouter()

	// DELETE ACCOUNT BY ID
	accountApiRouter.HandleFunc(
		"/{id}",
		accRoute.Middleware.CheckAuth(
			utils.ApiHandler(accRoute.Handlers.DeleteAccountHandler),
		),
	).Methods("DELETE")

	// GET ACCOUNTS
	accountApiRouter.HandleFunc(
		"",
		accRoute.Middleware.CheckAuth(
			utils.ApiHandler(accRoute.Handlers.GetAccountsHandler),
		),
	).Methods("GET")

	// GET ACCOUNT BY ID
	accountApiRouter.HandleFunc(
		"/{id}",
		accRoute.Middleware.CheckAuth(
			utils.ApiHandler(accRoute.Handlers.GetAccountByIdHandler),
		),
	).Methods("GET")

	// TRANSFER MONEY
	accountApiRouter.HandleFunc(
		"/transfer",
		accRoute.Middleware.CheckAuth(
			utils.ApiHandler(accRoute.Handlers.TransferMoneyHandler),
		),
	).Methods("POST")
}
