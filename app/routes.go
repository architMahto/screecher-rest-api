package app

import (
	"net/http"

	"github.com/architMahto/screecher-rest-api/app/handlers"
	"github.com/gorilla/mux"
)

type ApiRouter struct {
	Router *mux.Router
}

func InitializeRoutes(router *mux.Router) {
	apiRouter := ApiRouter{Router: router}

	// Home Route
	apiRouter.Get("/", handlers.Home)

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
