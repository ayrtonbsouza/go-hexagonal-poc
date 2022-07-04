package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ayrtonbsouza/hexagonal-architecture-poc/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.IProductService
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
