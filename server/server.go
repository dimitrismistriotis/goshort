package server

import (
	"code.google.com/p/gorilla/mux"
	"fmt"
	"github.com/op/go-logging"
	"net/http"
	"os"
)

func Serve(log *logging.Logger, port string) {
	log.Debug(fmt.Sprintf("Starting server on port %s", port))

	r := mux.NewRouter()

	// For IAMALIVE testing:
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	_, err := redis_instance()

	if err != nil {
		log.Critical("Cannot access Redis server")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	server.ListenAndServe()
}
