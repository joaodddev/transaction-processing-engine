package metrics

import "github.com/prometheus/client_golang/prometheus"

var TransactionsProcessed = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "transactions_processed_total",
		Help: "Total number of processed transactions",
	},
)

var HttpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	[]string{"method", "path"},
)

func Register() {

	prometheus.MustRegister(
		TransactionsProcessed,
		HttpRequestsTotal,
	)
}
