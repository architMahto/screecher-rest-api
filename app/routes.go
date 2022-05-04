package app

import (
	"net/http"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"github.com/architMahto/screecher-rest-api/app/handlers"
	"github.com/architMahto/screecher-rest-api/middleware"
	"github.com/gorilla/mux"
)

type ApiRouter struct {
	Router *mux.Router
}

func InitializeRoutes(router *mux.Router, fileDb *clients.FileDBClient) {
	apiRouter := ApiRouter{Router: router}
	userHandler := handlers.NewUserHandler(fileDb)
	screechHandler := handlers.NewScreechHandler(fileDb)

	// Home Route
	apiRouter.Get("/", handlers.Home)

	// User Routes
	apiRouter.Get("/api/users", userHandler.GetAllUsers)
	apiRouter.Get("/api/users/{user_id:[0-9]+}", userHandler.GetUserById)

	// Screech Routes
	apiRouter.Get(
		"/api/screeches",
		middleware.ValidateScreechQueryParams(screechHandler.GetAllScreeches),
	)
	apiRouter.Get("/api/screeches/{screech_id:[0-9]+}", screechHandler.GetScreechById)

	// 404 Not Found
	router.NotFoundHandler = http.HandlerFunc(handlers.HandleNotFound)
}

func (apiRouter *ApiRouter) Get(
	path string,
	fn func(res http.ResponseWriter, req *http.Request),
) {
	apiRouter.Router.HandleFunc(path, fn).Methods("GET")
}

func (apiRouter *ApiRouter) Post(
	path string,
	fn func(res http.ResponseWriter, req *http.Request),
) {
	apiRouter.Router.HandleFunc(path, fn).Methods("POST")
}

func (apiRouter *ApiRouter) Put(
	path string,
	fn func(res http.ResponseWriter, req *http.Request),
) {
	apiRouter.Router.HandleFunc(path, fn).Methods("PUT")
}

func (apiRouter *ApiRouter) Delete(
	path string,
	fn func(res http.ResponseWriter, req *http.Request),
) {
	apiRouter.Router.HandleFunc(path, fn).Methods("DELETE")
}
