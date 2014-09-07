package server

import (
	godis "github.com/simonz05/godis/redis"
	"net/url"
	"os"
	"time"
)

// Gets redis instance from configuration parameters or environment's REDISTOGO_URL.
func redis_instance() (redis *godis.Client, err error) {
	var host, password string

	if os.Getenv("REDISTOGO_URL") != "" {
		host, password = parseRedistogoUrl()
	} else {
		host = "tcp:localhost:6379"
		password = ""
	}

	db := 0

	redis = godis.New(host, db, password)

	// Try to set a "dummy" value, panic if redis is not accessible.
	err = redis.Set("INIT", time.Now().UTC())

	return
}

// Great artists steal:
// https://gist.github.com/TheDudeWithTheThing/6468746
//
func parseRedistogoUrl() (string, string) {
	redisUrl := os.Getenv("REDISTOGO_URL")
	redisInfo, _ := url.Parse(redisUrl)
	server := redisInfo.Host
	password := ""
	if redisInfo.User != nil {
		password, _ = redisInfo.User.Password()
	}
	return server, password
}
