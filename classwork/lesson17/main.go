package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*

const maxSize = 10_000_000

func appendSlice() {
	for {
		var s []int
		for i := 0; i < maxSize; i++ {
			s = append(s, i)
		}
	}
}
func main() {
	go appendSlice()
	if err := http.ListenAndServe(":8080", nil); err != nil { //   	/debug/pprof/
		log.Println(err)
		return
	}
} */

/*
 go build -o app.exe main.go
 go run .\main.go
 go tool pprof -http=":9090" .\app.exe -seconds=30 http://localhost:8080/debug/pprof/profile
*/

var httpTotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "total_requests",
		Help: "Total HTTP requests",
	},
	[]string{"path"},
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	httpTotalRequests.WithLabelValues(r.URL.Path).Inc()
	fmt.Fprintln(w, "INDEX PAGE!")
}
func main() {
	prometheus.MustRegister(httpTotalRequests)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}

// go get github.com/prometheus/client_golang/prometheus
// go get github.com/prometheus/client_golang/prometheus/promhttp

// .\prometheus.exe
