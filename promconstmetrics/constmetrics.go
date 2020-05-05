// Facade for hiding complexity of recording const metrics with explicit timestamp for Prometheus
package promconstmetrics

import (
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type labelValueMetric struct {
	timestamp   time.Time
	value       float64
	labelValues []string // empty, if Ref is not a vector
}

type Ref struct {
	desc               *prometheus.Desc
	valueType          prometheus.ValueType
	variableLabelCount int                 // if 0, len(metrics) == 1 and its labelValues is empty
	metrics            []*labelValueMetric // one for each unique labelValues (= variable labels) combo
	idxForLabelValues  map[string]int      // index to "metrics", keyed by hash of labelValues
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

var _ prometheus.Collector = (*Collector)(nil)

// registers a metric, or a metric vector if "variableLabels" given
func (c *Collector) Register(
	name string,
	help string,
	constLabels prometheus.Labels,
	variableLabels ...string,
) *Ref {
	c.mu.Lock()
	defer c.mu.Unlock()

	ref := &Ref{
		desc: prometheus.NewDesc(
			name,
			help,
			variableLabels,
			constLabels),
		variableLabelCount: len(variableLabels),
		valueType:          prometheus.GaugeValue,
		metrics:            []*labelValueMetric{},
		idxForLabelValues:  map[string]int{},
	}

	c.refs = append(c.refs, ref)

	return ref
}

func (c *Collector) Observe(ref *Ref, value float64, ts time.Time, labelValues ...string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(labelValues) != ref.variableLabelCount {
		panic("invalid variableLabelCount")
	}

	labelValuesKey := ""
	if len(labelValues) > 0 {
		labelValuesKey = strings.Join(labelValues, "|")
	}

	idx, found := ref.idxForLabelValues[labelValuesKey]
	if !found {
		ref.metrics = append(ref.metrics, &labelValueMetric{
			labelValues: labelValues,
		})

		idx = len(ref.metrics) - 1
		ref.idxForLabelValues[labelValuesKey] = idx
	}

	lvm := ref.metrics[idx]
	lvm.timestamp = ts
	lvm.value = value
}

// for prometheus.Collector
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	// "unchecked collector"
}

// for prometheus.Collector
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, ref := range c.refs {
		for _, metric := range ref.metrics {
			ch <- prometheus.NewMetricWithTimestamp(metric.timestamp, prometheus.MustNewConstMetric(
				ref.desc,
				ref.valueType,
				metric.value,
				metric.labelValues...))
		}
	}
}
