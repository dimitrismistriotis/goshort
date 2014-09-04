package main

import (
	"github.com/dimitrismistriotis/goshort/server"
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("short_server")

// Function main is designed to be called from a heroku style Procfile, example:
// PORT=8888 go run main.go
//
// Environment variables:
// Port: Server port to listen to.
func main() {
	var log = logging.MustGetLogger("goshort")
	var format = logging.MustStringFormatter("%{color} %{level}%{id:03x}%{color:reset} %{message}")
	logging.SetFormatter(format)

	log.Debug("Starting from main")
	var port = "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	server.Serve(log, port)

	os.Exit(0)
}
