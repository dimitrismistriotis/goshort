package server

import (
	"github.com/op/go-logging"
	godis "github.com/simonz05/godis/redis"
	"net/http"
	"os"
)

const NEXT_KEY_ID = "next_key_id"

//
// Example: curl http://localhost:8080/short/ --data-urlencode "u=http://www.mistriotis.com" -v
// TODO: Put on readme
//
func shortUrl(w http.ResponseWriter, r *http.Request, log *logging.Logger, redis *godis.Client) {
	log.Debug("Call to shortUrl")
	host := "localhost"
	if os.Getenv("HOSTNAME") != "" {
		host = os.Getenv("HOSTNAME")
	}

	unshortedUrl := r.FormValue("u")

	//
	// TODO: Check if unshortedUrl is valid.
	//
	log.Info("Storing url: %s with host %s", string(unshortedUrl), host)

	//
	// TODO: Check if it already exists
	//

	// TODO: Get NEXT_KEY_ID
	// TODO: Inesert unshorten into NEXT_KEY_ID
	// TODO: increase NEXT_KEY_ID
	// TODO: Return id to hex
	// TODO: Create a data structure, populate that and return it as JSON

	w.Write([]byte("SHORT:" + unshortedUrl + " on " + host))
}
