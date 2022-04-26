package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	port := 8080
	log.Debug().Msgf("Starting Server on Port: %d", port)

	http.HandleFunc("/", routes)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
	if err != nil {
		log.Err(err)
		return
	}
}

func routes(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("To Be Implemented"))
	if err != nil {
		log.Err(err)
		return
	}
}
