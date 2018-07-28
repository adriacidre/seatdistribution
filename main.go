package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/adriacidre/seatdistribution/airplane"
)

var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", "9000", "Port to listen on")
	plane    *airplane.Airplane
)

func main() {
	plane = airplane.New()
	mux := http.NewServeMux()

	// assign a seat
	mux.HandleFunc("/sections", AddSectionHandler)
	mux.HandleFunc("/sections/assign", AssignHandler)
	mux.HandleFunc("/sections/seat", GetSeatHandler)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}
