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

	redis, err := redis_instance()

	if err != nil {
		log.Critical("Cannot access Redis server")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	r := mux.NewRouter()

	// For IAMALIVE testing:
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	// Unsure about this pattern, have also seen the same with having log and
	// redis in package variables.
	r.HandleFunc("/short/{url:(.*$)}", func(w http.ResponseWriter, r *http.Request) {
		shortUrl(w, r, log, redis, port)
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	server.ListenAndServe()
}
