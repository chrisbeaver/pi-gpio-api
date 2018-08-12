package server

import (
	"encoding/json"
	"strconv"
    // "log"
	"net/http"
	// "time"
	"os"
	"fmt"
	// "io"
	// "sort"
	"github.com/stianeikeland/go-rpio"
)

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin = rpio.Pin(10)
)
type GPIOPin struct {
	Name  string `json:"name"`
	State int    `json:"state"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	 
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()
	
	pins := make(map[int]GPIOPin)
	for v := 0; v < 28; v++ {
		pin := rpio.Pin(v)
		pins[v] = GPIOPin{Name: "GPIO" + strconv.Itoa(v), State: int(pin.Read())}
	}
	jsonString, _ := json.Marshal(pins)
	fmt.Fprintf(w, string(jsonString))
	
}

func HighHandler(w http.ResponseWriter, r *http.Request) { 
	for i, _ := range r.URL.Query() {
		fmt.Fprintf(w, i+"\n")
	}
	// fmt.Printf("%+v\n", yourProject)
}

func LowHandler(w http.ResponseWriter, r *http.Request) { }



