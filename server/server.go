package server

import (
	"code.google.com/p/gorilla/mux"
	"fmt"
	"net/http"
)

func Serve() {
	fmt.Println("Kickstarting server")

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	server.ListenAndServe()
}
