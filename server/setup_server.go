package server

import (
	redis "github.com/garyburd/redigo/redis"
	redisurl "github.com/soveran/redisurl"
	"os"
	"time"
)

// Gets redis instance from configuration parameters or environment's REDISTOGO_URL.
func redis_instance() (redis_con redis.Conn, err error) {
	if os.Getenv("REDISTOGO_URL") != "" {
		redis_con, err = redisurl.ConnectToURL(os.Getenv("REDISTOGO_URL"))
	} else {
		redis_con, err = redis.Dial("tcp", ":6379")
	}

	if err != nil {
		return
	}

	// Try to set a "dummy" value, panic if redis is not accessible.
	err = redis_con.Send("Set", "INIT", time.Now().UTC())

	return
}
