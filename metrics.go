package main

import (
	"strings"

	"github.com/Masterminds/log-go"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	sampleMetric *prometheus.GaugeVec
)

func initCollectors() {
	defaultLabels := []string{"instance"}
	log.Debugf("Default labels: %v", strings.Join(defaultLabels, ","))

	sampleMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sample_metric",
			Help: "An example gauge metric",
		},
		defaultLabels,
	)
	prometheus.MustRegister(sampleMetric)
}
