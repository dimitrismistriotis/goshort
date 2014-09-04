package server

import (
	"code.google.com/p/gorilla/mux"
	"fmt"
	"github.com/op/go-logging"
	"net/http"
)

func Serve(log *logging.Logger, port string) {
	log.Debug(fmt.Sprintf("Starting server on port %s", port))

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	server.ListenAndServe()
}
