package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen-address", ":8085", "The address to listen on for HTTP requests.")

	// START METRICS OMIT
	inFlightGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "myhttp",
		Name: "in_flight_requests",
		Help: "A guage of requests currently being served by the wrapped handler",
	})

	counter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "myhttp",
		Name: "requests_total",
		Help: "A counter for requests to the wrapped handler",
	},
		[]string{"code", "method"})

	duration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "myhttp",
		Name: "request_duration_seconds",
		Help: "A histogram of latencies for requests",
	},
		[]string{"handler","code","method"})
	// END METRICS OMIT
)

func main() {
	flag.Parse()

	// START HANDLERS OMIT
	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	goodbyeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Goodbye"))
	})

	helloChain := promhttp.InstrumentHandlerInFlight(
		inFlightGauge,
		promhttp.InstrumentHandlerDuration(
			duration.MustCurryWith(prometheus.Labels{"handler": "hello"}),
			promhttp.InstrumentHandlerCounter(counter, helloHandler)))

	goodbyeChain := promhttp.InstrumentHandlerInFlight(
		inFlightGauge,
		promhttp.InstrumentHandlerDuration(
			duration.MustCurryWith(prometheus.Labels{"handler": "goodbye"}),
			promhttp.InstrumentHandlerCounter(counter, goodbyeHandler)))
	// END HANDLERS OMIT

	// START ROUTES OMIT
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/hello", helloChain)
	http.Handle("/goodbye", goodbyeChain)
	// END ROUTES OMIT

	log.Fatal(http.ListenAndServe(*addr, nil))
}
