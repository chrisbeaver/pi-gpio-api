package server

import (
	// "encoding/json"
    // "log"
	"net/http"
	// "io"
	"github.com/stianeikeland/go-rpio"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	 
	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	// io.WriteString(w, `{"alive": true}`)
	pin := rpio.Pin(10)

	pin.Output()       // Output mode
	pin.High()         // Set pin High
	pin.Low()          // Set pin Low
	pin.Toggle()       // Toggle pin (Low -> High -> Low)

	pin.Input()        // Input mode
	// io.WriteString(w, pin.Read())  // Read state from pin (High / Low)
}



