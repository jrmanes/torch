package handlers

import (
	"net/http"

	"github.com/jrmanes/torch/config"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Router(r *mux.Router, cfg config.MutualPeersConfig) *mux.Router {
	r.Use(LogRequest)

	s := r.PathPrefix("/api/v1").Subrouter()

	// get config
	s.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		GetConfig(w, r, cfg)
	}).Methods("GET")

	// get nodes
	s.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		List(w, r, cfg)
	}).Methods("GET")
	// get node details by node name
	s.HandleFunc("/noId/{nodeName}", func(w http.ResponseWriter, r *http.Request) {
		GetNoId(w, r, cfg)
	}).Methods("GET")

	// generate
	s.HandleFunc("/gen", func(w http.ResponseWriter, r *http.Request) {
		Gen(w, r, cfg)
	}).Methods("POST")
	s.HandleFunc("/genAll", func(w http.ResponseWriter, r *http.Request) {
		GenAll(w, r, cfg)
	}).Methods("POST")

	// metrics
	s.Handle("/metrics", promhttp.Handler())

	return r
}
