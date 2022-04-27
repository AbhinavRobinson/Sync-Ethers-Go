package main

import (
	"github.com/rs/zerolog/log"
)

func initApp() {
	log.Debug().Msg("Initializing App...")
}

func initApi(host string, port int) {
	log.Debug().Msg("Initializing API...")
	setupApi(host, port)
}
