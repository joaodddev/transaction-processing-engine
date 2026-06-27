package middleware

import (
	"log"
	nethttp "net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/joaodddev/transaction-processing-engine/internal/observability/metrics"
)

func Logging(next nethttp.Handler) nethttp.Handler {

	return nethttp.HandlerFunc(func(
		w nethttp.ResponseWriter,
		r *nethttp.Request,
	) {

		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Printf(
			"%s %s %s",
			r.Method,
			r.URL.Path,
			duration,
		)

		metrics.HttpRequestsTotal.With(
			prometheus.Labels{
				"method": r.Method,
				"path":   r.URL.Path,
			},
		).Inc()
	})
}
