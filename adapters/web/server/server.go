package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Codeflix-FullCycle/hexagonal-architecture/adapters/web/handler"
	"github.com/Codeflix-FullCycle/hexagonal-architecture/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer(Service application.ProductServiceInterface) *WebServer {
	return &WebServer{
		Service: Service,
	}
}

func (w *WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}
