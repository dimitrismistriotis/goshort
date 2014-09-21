package server

import (
	"github.com/op/go-logging"
	"os"

	redis "github.com/garyburd/redigo/redis"
	redisurl "github.com/soveran/redisurl"
)

// Gets Redis instance from configuration parameters or environment's REDISTOGO_URL.
func redis_instance(log *logging.Logger) (redis_con redis.Conn, err error) {
	if os.Getenv("REDISTOGO_URL") != "" {
		log.Debug("Instantiating with REDISTOGO_URL.")
		redis_con, err = redisurl.ConnectToURL(os.Getenv("REDISTOGO_URL"))
	} else {
		log.Debug("Instantiating with local-host.")
		redis_con, err = redis.Dial("tcp", ":6379")
	}

	if err != nil {
		log.Critical("Error on instantiation.")
		return
	}

	// Try to set a "dummy" value, panic if Redis is not accessible.
	// err = redis_con.Send("SET", "INIT", time.Now().UTC())

	return
}
