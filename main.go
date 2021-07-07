package main

import (
	"fmt"
	"github.com/go-http-utils/logger"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/echo/"):]
	log.Printf("received a friendly greeting at `%s`", title)
	fmt.Fprintf(w, "%s\n", title)
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	log.Printf("counter is now %d", counter)
	fmt.Fprintf(w, "%s\n", strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hi", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Simply Marvelous, Darling!\n"))
	})

	mux.HandleFunc("/more", incrementCounter)
	mux.HandleFunc("/echo/", echoString)

	log.Print("serving on port 80")

	log.Fatal(http.ListenAndServe(":80", logger.Handler(mux, os.Stdout, logger.DevLoggerType)))
}
