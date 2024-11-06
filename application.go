package golwf

import (
	"fmt"
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

func (application *Application) Run(logging bool) {
	application.mux.HandleFunc("/", application.router.ServeHTTP)

	fmt.Printf("Starting server on http://0.0.0.0:%d", application.port)

	err := http.ListenAndServe(":"+strconv.Itoa(application.port), application.mux)

	log.Fatal(err)
}
