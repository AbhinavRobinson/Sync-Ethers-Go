package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "Jan 02 15:04"})

	host := "localhost"
	port := 8080

	// Setups
	initApp()
	initApi(host, port)
}
