package server

import (
	"encoding/json"
	"fmt"
	"github.com/op/go-logging"
	godis "github.com/simonz05/godis/redis"
	"net/http"
	"os"
)

const (
	NEXT_KEY_ID = "next_key_id"
	URL_PREFIX  = "URL_"
)

type GoShortResponse struct {
	Hostname string
	Route    string
	FullUrl  string
	Message  string
	Error    string
}

//
// Example: curl http://localhost:8080/short/ --data-urlencode "u=http://www.mistriotis.com" -v
//
func shortUrl(w http.ResponseWriter, r *http.Request, log *logging.Logger, redis *godis.Client, port string) {
	log.Debug("Call to shortUrl")

	host := "localhost"
	if os.Getenv("HOSTNAME") != "" {
		host = os.Getenv("HOSTNAME")
	}

	if port != "" || port != "80" {
		host += ":" + port
	}

	unshortedUrl := r.FormValue("u")

	//
	// TODO: Check if unshortedUrl is valid.
	//

	//
	// TODO: If not valid return an erroneous JSON response.
	//

	//
	// TODO: Check if it already exists.
	//

	//
	// TODO: If it already exists return the existent one in a JSON response.
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
	log.Debug("Key: %d to hex %s", next_key, fmt.Sprintf("%x", next_key))

	redis.Set(URL_PREFIX+fmt.Sprintf("%d", next_key), unshortedUrl)

	response := GoShortResponse{
		Hostname:   host,
		Route:      "/" + fmt.Sprintf("%x", next_key),
		ShortedUrl: host + "/" + fmt.Sprintf("%x", next_key),
		Message:    "Url: " + unshortedUrl + " shorted to " + host + "/" + fmt.Sprintf("%x", next_key),
		Error:      "",
	}

	var res, _ = json.Marshal(response)
	w.Write(res)
}
