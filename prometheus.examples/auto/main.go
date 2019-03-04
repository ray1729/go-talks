// START OMIT
package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto" // HL
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen-address", ":8082", "The address to listen on for HTTP requests.")

	fooCounter = promauto.NewCounter(prometheus.CounterOpts{ // HL
		Name: "foo_total",
		Help: "Number of times foo has been invoked",
	})
)
// END OMIT

func foo() {
	fooCounter.Inc()
}

func main() {
	flag.Parse()
	go func() {
		for {
			foo()
			time.Sleep(time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
