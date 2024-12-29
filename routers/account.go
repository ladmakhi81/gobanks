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
	// DELETE ACCOUNT BY ID
	accRoute.Router.HandleFunc(
		"/api/v1/accounts/{id}",
		accRoute.Middleware.CheckAuth(
			utils.ApiHandler(accRoute.Handlers.DeleteAccountHandler),
		),
	).Methods("DELETE")

	// GET ACCOUNTS
	accRoute.Router.HandleFunc(
		"/api/v1/accounts",
		accRoute.Middleware.CheckAuth(
			utils.ApiHandler(accRoute.Handlers.GetAccountsHandler),
		),
	).Methods("GET")

	// GET ACCOUNT BY ID
	accRoute.Router.HandleFunc(
		"/api/v1/accounts/{id}",
		accRoute.Middleware.CheckAuth(
			utils.ApiHandler(accRoute.Handlers.GetAccountByIdHandler),
		),
	).Methods("GET")

	// TRANSFER MONEY
	accRoute.Router.HandleFunc(
		"/api/v1/accounts/transfer",
		accRoute.Middleware.CheckAuth(
			utils.ApiHandler(accRoute.Handlers.TransferMoneyHandler),
		),
	).Methods("POST")
}
