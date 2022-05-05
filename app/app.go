package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/architMahto/screecher-rest-api/app/clients"
	"github.com/architMahto/screecher-rest-api/middleware"
	"github.com/gorilla/mux"
)

func Run() {
	config := GetConfig()

	csvDbClient := clients.CsvDbClient{
		PathToDataDir: config.DataDirPath,
	}
	jsonDbClient := clients.JsonDbClient{
		PathToDataDir: config.DataDirPath,
	}

	router := mux.NewRouter()
	InitializeRoutes(router, &csvDbClient, &jsonDbClient)

	logger := log.New(os.Stdout, "", log.LstdFlags)
	loggingMiddleware := middleware.NewLoggingMiddleware(logger)

	router.Use(loggingMiddleware.Func())

	fmt.Println("Server running at http://localhost:" + config.AppPort)
	log.Fatal(http.ListenAndServe(":"+config.AppPort, router))
}
