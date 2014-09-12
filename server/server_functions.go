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

	//
	// TODO: Check if it already exists
	//

	var redis_next_key, _ = redis.Get(NEXT_KEY_ID)
	var next_key int64
	if redis_next_key == nil {
		next_key = 1
		redis.Set(NEXT_KEY_ID, 1)
	} else {
		next_key = redis_next_key.Int64()
	}
	redis.Incr(NEXT_KEY_ID)

	log.Info("Storing url: %s with host %s on key %s", string(unshortedUrl), host, next_key)

	// TODO: Return id to hex
	// TODO: Create a data structure, populate that and return it as JSON

	w.Write([]byte("SHORT:" + unshortedUrl + " on " + host))
}
