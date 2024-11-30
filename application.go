package golwf

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
)

type Application struct {
	mux    *http.ServeMux
	router *Router
	port   int
}

func InitApplication(port int) *Application {
	return &Application{mux: http.NewServeMux(), router: NewRouter(), port: port}
}

func (application *Application) GetRouter() *Router {
	return application.router
}

func (application *Application) Run(logging bool) {
	application.mux.HandleFunc("/", application.router.ServeHTTP)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH", "HEAD", "CONNECT", "TRACE"},
	}).Handler(application.mux)

	fmt.Printf("Starting server on http://0.0.0.0:%d\n", application.port)

	err := http.ListenAndServe(":"+strconv.Itoa(application.port), handler)

	log.Fatal(err)
}
