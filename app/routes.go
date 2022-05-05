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

func InitializeRoutes(
	router *mux.Router,
	csvDb *clients.CsvDbClient,
	jsonDb *clients.JsonDbClient,
) {
	apiRouter := ApiRouter{Router: router}
	userHandler := handlers.NewUserHandler(csvDb)
	screechHandler := handlers.NewScreechHandler(csvDb)
	authHandler := handlers.NewAuthHandler(csvDb, jsonDb)

	// Home Route
	apiRouter.Get("/", handlers.Home)

	// Auth Routers
	apiRouter.Post(
		"/api/auth/signup",
		middleware.ValidateUserCreateReqBody(userHandler.CreateUser),
	)
	apiRouter.Post(
		"/api/auth/signin",
		middleware.ValidateUserSignInReqBody(authHandler.SignIn),
	)

	// User Routes
	apiRouter.Get("/api/users", middleware.DoesAuthHeaderExist(userHandler.GetAllUsers))
	apiRouter.Get("/api/users/{user_id:[0-9]+}",
		middleware.DoesAuthHeaderExist(userHandler.GetUserById),
	)
	apiRouter.Put(
		"/api/users/{user_id:[0-9]+}",
		middleware.DoesAuthHeaderExist(
			middleware.ValidateUserUpdateReqBody(userHandler.UpdateUser),
		),
	)

	// Screech Routes
	apiRouter.Get(
		"/api/screeches",
		middleware.ValidateScreechQueryParams(screechHandler.GetScreeches),
	)
	apiRouter.Get("/api/screeches/{screech_id:[0-9]+}", screechHandler.GetScreechById)
	apiRouter.Post(
		"/api/screeches",
		middleware.DoesAuthHeaderExist(
			middleware.ValidateScreechBody(screechHandler.CreateScreech),
		),
	)
	apiRouter.Put(
		"/api/screeches/{screech_id:[0-9]+}",
		middleware.DoesAuthHeaderExist(
			middleware.ValidateScreechBody(screechHandler.UpdateScreech),
		),
	)

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
