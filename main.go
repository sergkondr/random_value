package random_value

import (
	"math/rand"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type SomeMetric struct {
	Min int `toml:"minimal"`
	Max int `toml:"maximal"`
}

func (_ *SomeMetric) Description() string {
	return "This plugin returns random number between min and max"
}

func (_ *SomeMetric) SampleConfig() string {
	var sampleConfig = `
[[inputs.random_value]]
  minimal = 0
  maximal = 100
`
	return sampleConfig
}

func (m *SomeMetric) Gather(acc telegraf.Accumulator) error {
	now := time.Now()

	acc.AddFields("random_value", map[string]interface{}{
		"value":   getValue(m.Min, m.Max),
		"minimal": m.Min,
		"maximal": m.Max,
	}, nil, now)
	return nil
}

func getValue(min, max int) int {
	return rand.Intn(max-min) + min
}

func init() {
	inputs.Add("random_value", func() telegraf.Input {
		return &SomeMetric{}
	})
}
