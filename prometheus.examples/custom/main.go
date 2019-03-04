// PART1 OMIT
package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")

	fooCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "foo_total",
		Help: "Number of times foo has been invoked",
	})
)
// PART2 OMIT

func init() {
	prometheus.MustRegister(fooCounter)
}

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
// END OMIT
