package main

import (
	f "fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	l "github.com/rs/zerolog/log"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l.Logger = l.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	port := 8080
	l.Debug().Msgf("Starting Server on Port: %d", port)

	http.HandleFunc("/", routes)
	err := http.ListenAndServe(f.Sprintf("localhost:%d", port), nil)
	if err != nil {
		l.Err(err)
		return
	}
}

func routes(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello, World"))
	if err != nil {
		l.Err(err)
		return
	}
}
