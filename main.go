package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ladmakhi81/gobanks/database"
	"github.com/ladmakhi81/gobanks/handlers"
	"github.com/ladmakhi81/gobanks/middlewares"
	"github.com/ladmakhi81/gobanks/repositories"
	"github.com/ladmakhi81/gobanks/routers"
	"github.com/ladmakhi81/gobanks/utils"
)

func main() {
	router := mux.NewRouter()
	const port = ":8080"

	db := database.NewDatabaseServer()
	db.Setup()

	// repositories
	accountRepo := repositories.AccountRepository{
		DatabaseServer: db,
	}
	sessionRepo := repositories.SessionRepository{
		DatabaseServer: db,
	}

	// middlewares
	middleware := middlewares.Middlewares{
		TokenUtil: utils.TokenUtil{
			AccountRepo: accountRepo,
			SessionRepo: sessionRepo,
		},
	}

	// handlers
	accountHandlers := handlers.AccountHandler{
		Repo: accountRepo,
	}
	authHandlers := handlers.AuthHandler{
		SessionRepo: sessionRepo,
		AccountRepo: accountRepo,
	}

	// routers
	accountRouter := routers.AccountRoute{
		Router:     router,
		Handlers:   accountHandlers,
		Middleware: middleware,
	}
	authRouter := routers.AuthRoute{
		Router:     router,
		Handlers:   authHandlers,
		Middleware: middleware,
	}

	accountRouter.Setup()
	authRouter.Setup()

	log.Println("the server running", port)

	err := http.ListenAndServe(port, router)

	log.Fatalln(err)
}
