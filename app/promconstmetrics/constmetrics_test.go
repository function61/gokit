package promconstmetrics

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/function61/gokit/testing/assert"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
)

var (
	t0 = time.Date(2020, 1, 22, 12, 0, 0, 0, time.UTC)
)

func TestConstMetrics(t *testing.T) {
	c := NewCollector()

	allCollectors := prometheus.NewRegistry()
	assert.Ok(t, allCollectors.Register(c))

	stars := c.Register("stars", "Stars in GitHub", prometheus.Labels{
		"repo": "varasto",
		"org":  "function61",
	})

	c.Observe(stars, 3, t0)

	expositionOutput := &bytes.Buffer{}

	assert.Ok(t, gatherToTextExport(allCollectors, expositionOutput))
	assert.Equal(t, expositionOutput.String(), `# HELP stars Stars in GitHub
# TYPE stars gauge
stars{org="function61",repo="varasto"} 3 1579694400000
`)

	c.Observe(stars, 11, t0.Add(2*time.Second))

	expositionOutput.Reset()

	assert.Ok(t, gatherToTextExport(allCollectors, expositionOutput))
	assert.Equal(t, expositionOutput.String(), `# HELP stars Stars in GitHub
# TYPE stars gauge
stars{org="function61",repo="varasto"} 11 1579694402000
`)
}

func TestVariableLabels(t *testing.T) {
	c := NewCollector()

	allCollectors := prometheus.NewRegistry()
	assert.Ok(t, allCollectors.Register(c))

	stars := c.Register("stars", "Stars in GitHub", prometheus.Labels{
		"org": "function61",
	}, "repo")

	c.Observe(stars, 11, t0, "varasto")
	c.Observe(stars, 5, t0, "gokit")

	expositionOutput := &bytes.Buffer{}

	assert.Ok(t, gatherToTextExport(allCollectors, expositionOutput))
	assert.Equal(t, expositionOutput.String(), `# HELP stars Stars in GitHub
# TYPE stars gauge
stars{org="function61",repo="gokit"} 5 1579694400000
stars{org="function61",repo="varasto"} 11 1579694400000
`)
}

func gatherToTextExport(registry *prometheus.Registry, output io.Writer) error {
	metricFamilies, err := registry.Gather()
	if err != nil {
		return err
	}

	wireEncoder := expfmt.NewEncoder(output, expfmt.FmtText)

	for _, metricFamily := range metricFamilies {
		if err := wireEncoder.Encode(metricFamily); err != nil {
			return err
		}
	}

	return nil
}
