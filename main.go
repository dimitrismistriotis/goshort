package main

import (
	"fmt"
	"github.com/dimitrismistriotis/goshort/server"
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("short_server")

func main() {
	var format = logging.MustStringFormatter("%{level} %{message}")
	logging.SetFormatter(format)

	server.Serve()
	log.Info("Starting from main")

	os.Exit(0)
}
