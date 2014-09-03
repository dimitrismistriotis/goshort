package main

import (
	"github.com/dimitrismistriotis/goshort/server"
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("short_server")

func main() {
	var log = logging.MustGetLogger("goshort")
	var format = logging.MustStringFormatter("%{color} %{level}%{id:03x}%{color:reset} %{message}")
	logging.SetFormatter(format)

	log.Debug("Starting from main")
	server.Serve(log)

	os.Exit(0)
}
