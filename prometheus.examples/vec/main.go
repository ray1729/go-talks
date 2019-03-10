package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// START OMIT
var (
	addr = flag.String("listen-address", ":8083", "The address to listen on for HTTP requests.")

	fooCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "foo_total",
			Help: "Number of times foo has been invoked",
		},
		[]string{"name"},
	)
)

func foo() {
	names := []string{"maggie", "milly", "molly", "may"}
	n := names[rand.Intn(len(names))]
	fooCounter.With(prometheus.Labels{"name": n}).Inc()
}
// END OMIT

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
