package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type Specification struct {
	HttpPort string // to set: export GOAPP_HTTPPORT="8080"
}

var s Specification

func main() {
	err := setup()
	if err != nil {
		log.Fatal("Failed setup: ", err)
	}

	log.Print("Listening on port " + s.HttpPort)
	http.HandleFunc("/", handler) // redirect all urls to the handler function
	http.ListenAndServe("0.0.0.0:"+s.HttpPort, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world on port "+s.HttpPort+" \n")
}

func setup() error {
	log.Print("starting up")

	err := envconfig.Process("goapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	if s.HttpPort == "" {
		log.Print("No HttpPort url provided, defaulting to 9999")
		s.HttpPort = "9999"
	}
	return nil
}
