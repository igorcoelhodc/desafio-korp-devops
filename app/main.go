package main

import (
	"encoding/json"
	"log"
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
	upMetric.Set(1)
}

type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestsTotal.WithLabelValues("/projeto-korp").Inc()

	w.Header().Set("Content-Type", "application/json")

	resp := Response{
		Nome:	"Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "erro ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("serving on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}