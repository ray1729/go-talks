package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen-address", ":8084", "The address to listen on for HTTP requests.")

// START METRICS OMIT
	fooCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "foo_total",
		Help: "Number of times foo has been invoked",
	})

	fooErrCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "foo_error_total",
		Help: "Number of times foo has returned an error",
	})

	fooDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "foo_duration_seconds",
		Help: "Duration of foo invocation in seconds",
		Buckets: []float64{.001, .005, .01, .05, .1, .5, 1, 5},
	})
)
// END METRICS OMIT

// START FUNC OMIT
func foo() error {
	time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
	if rand.Float32() < .1 {
		return fmt.Errorf("foo error")
	}
	return nil
}
// END FUNC OMIT

// START WRAP OMIT
func Foo() error {
	fooCounter.Inc()
	start := time.Now()
	err := foo()
	if err != nil {
		fooErrCounter.Inc()
	}
	fooDuration.Observe(time.Now().Sub(start).Seconds())
	return err
}
// END WRAP OMIT

func main() {
	flag.Parse()
	go func() {
		for {
			Foo()
			time.Sleep(100*time.Millisecond)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
