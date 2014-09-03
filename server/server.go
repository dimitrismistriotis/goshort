package server

import (
	"code.google.com/p/gorilla/mux"
	"github.com/op/go-logging"
	"net/http"
)

// TODO: Add host to signature
// TODO: Add port to signature
func Serve(log *logging.Logger) {
	// TODO: Log insteas of fmt
	log.Debug("Starting server")

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	server := &http.Server{
		// TODO: Format host and port
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	server.ListenAndServe()
}
