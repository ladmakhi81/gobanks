package routers

import (
	"github.com/gorilla/mux"
	"github.com/ladmakhi81/gobanks/handlers"
	"github.com/ladmakhi81/gobanks/middlewares"
	"github.com/ladmakhi81/gobanks/utils"
)

type AuthRoute struct {
	Router     *mux.Router
	Handlers   handlers.AuthHandler
	Middleware middlewares.Middlewares
}

func (authRoute AuthRoute) Setup() {
	// sign-in account
	authRoute.Router.HandleFunc(
		"/auth/sign-in",
		utils.ApiHandler(authRoute.Handlers.Login),
	)

	// sign-up account
	authRoute.Router.HandleFunc(
		"/auth/sign-up",
		utils.ApiHandler(authRoute.Handlers.Signup),
	)

	// sign-out account
	authRoute.Router.HandleFunc(
		"/auth/sign-out",
		authRoute.Middleware.CheckAuth(
			utils.ApiHandler(authRoute.Handlers.Logout),
		),
	).Methods("DELETE")

	// get profile account
	authRoute.Router.HandleFunc(
		"/auth/profile",
		authRoute.Middleware.CheckAuth(
			utils.ApiHandler(authRoute.Handlers.ProfileAccount),
		),
	).Methods("GET")
}
