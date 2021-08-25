package my_prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sync"
)

var once sync.Once

func Server() {
	once.Do(func() {
		// create a new mux server
		server := http.NewServeMux()
		// register a new handler for the /metrics endpoint
		server.Handle("/metrics", promhttp.Handler())
		// start an http server using the mux server
		log.Fatalln(http.ListenAndServe(":9002", server))
	})

}
