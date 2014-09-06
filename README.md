# GO short - Yet another URL shortener in go

It seems that everybody's favourite golang learning exercise is to write a URL shortener. This project intents to have the following differences from similar projects:

* Coded with a "'''main.go'''" file that uses functionality in a library.
* Designed to be deployed on heroku.
* Configurable to shorten ULRs from specific sites.
* Attempts to retrieve configuration values from environment using heroku's conventions and then from a configuration file.

## Build

Run
```bash
go get .
```
in order to retrieve all external dependencies, which should be:

* Gorilla Mux
* godis - a Redis client for Go

## Deploy

TODO
