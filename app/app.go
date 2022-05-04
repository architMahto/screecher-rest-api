package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"github.com/architMahto/screecher-rest-api/middlewares"
	"github.com/gorilla/mux"
)

func Run() {
	config := GetConfig()

	fileDbClient := clients.FileDBClient{
		PathToDataDir: config.DataDirPath,
	}

	router := mux.NewRouter()
	InitializeRoutes(router, &fileDbClient)

	logger := log.New(os.Stdout, "", log.LstdFlags)
	loggingMiddleware := middlewares.NewLoggingMiddleware(logger)

	router.Use(loggingMiddleware.Func())

	fmt.Println("Server running at http://localhost:" + config.AppPort)
	log.Fatal(http.ListenAndServe(":"+config.AppPort, router))
}
