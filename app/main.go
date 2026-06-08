package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Volume total de requisições recebidas.",
		},
		[]string{"path"},
	)
	upMetric = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "up",
			Help: "Disponibilidade do serviço (1 = online, 0 = offline).",
		},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(upMetric)
	upMetrics.Set(1)
}

type Response struct {
	Nome	string ˋjson:"nome"ˋ
	Horario	string ˋjson:"horario"ˋ
}

func desafioKorpHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestTotal.WithLabelValues("/desafio-korp").Inc()

	w.Header().Set("Content-Type", "application/json")

	resp := Response{
		Nome:	"Desafio Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/desafio-korp", desafioKorpHandler)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}