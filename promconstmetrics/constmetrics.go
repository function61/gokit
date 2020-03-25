// Facade for hiding complexity of recording const metrics with explicit timestamp for Prometheus
package promconstmetrics

import (
	"sort"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Ref struct {
	desc         *prometheus.Desc
	labelValues  []string
	latestMetric prometheus.Metric // is nil until first Observe() call
}

type Collector struct {
	refs []*Ref
	mu   sync.Mutex
}

func NewCollector() *Collector {
	return &Collector{
		refs: []*Ref{},
	}
}

func (c *Collector) Register(name string, help string, labels prometheus.Labels) *Ref {
	c.mu.Lock()
	defer c.mu.Unlock()

	labelKeys, labelValues := splitLabelsAndValues(labels)

	ref := &Ref{
		desc:        prometheus.NewDesc(name, help, labelKeys, nil),
		labelValues: labelValues,
	}

	c.refs = append(c.refs, ref)

	return ref
}

func (c *Collector) Observe(ref *Ref, value float64, ts time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	metrics, err := prometheus.NewConstMetric(
		ref.desc,
		prometheus.GaugeValue,
		value,
		ref.labelValues...)
	if err != nil {
		panic(err)
	}

	ref.latestMetric = prometheus.NewMetricWithTimestamp(ts, metrics)
}

// contract of prometheus.Collector
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	// unchecked collector
}

// contract of prometheus.Collector
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, ref := range c.refs {
		// first Observe() not called => no collection
		if ref.latestMetric == nil {
			continue
		}

		ch <- ref.latestMetric
	}
}

// {"name": "Joonas", "occupation": "Captain"} =>
//   ["name", "occupation"]
//   ["Joonas", "Captain"]
func splitLabelsAndValues(labels prometheus.Labels) ([]string, []string) {
	labelKeys := []string{}
	for key := range labels {
		labelKeys = append(labelKeys, key)
	}

	sort.Strings(labelKeys)

	labelValues := []string{}
	for _, key := range labelKeys {
		labelValues = append(labelValues, labels[key])
	}

	return labelKeys, labelValues
}
