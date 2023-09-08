package main

import (
	"flightcheck/internal/health"
	"flightcheck/internal/iranair"
)

var quit = make(chan struct{})

func main() {
	go health.Run()

	go iranair.Run()
	// go atta.Run()
	// go airtour.Run()
	// go meraj.Run()

	<-quit
}
