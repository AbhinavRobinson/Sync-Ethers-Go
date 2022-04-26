package main

import (
	"github.com/rs/zerolog/log"
)

func initApp() {
	log.Info().Msg("Initializing App...")
	setupContracts()
}

func initApi(host string, port int) {
	log.Info().Msg("Initializing API...")
	setupApi(host, port)
}
