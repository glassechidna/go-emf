package emf

import (
	"encoding/json"
	"fmt"
	"github.com/glassechidna/go-emf/emf/unit"
	"io"
	"os"
	"time"
)

var Namespace = "goemf"
var LogGroupName = ""
var Writer io.Writer = os.Stdout

type MSI map[string]interface{}

func Emit(m MSI) {
	metrics := []cwMetricDefinition{}
	dimensions := []string{}

	raw := map[string]interface{}{}

	for key, value := range m {
		switch value := value.(type) {
		case *metric:
			raw[key] = value.Value
			metrics = append(metrics, cwMetricDefinition{Name: key, Unit: string(value.Unit)})
		case *dimension:
			raw[key] = value.Value
			dimensions = append(dimensions, key)
		default:
			raw[key] = value
		}
	}

	raw["_aws"] = metadata{
		Timestamp:    int(time.Now().UnixNano() / 1e6),
		LogGroupName: LogGroupName,
		CloudWatchMetrics: []cwMetricDirective{
			{
				Namespace:  Namespace,
				Dimensions: [][]string{dimensions},
				Metrics:    metrics,
			},
		},
	}

	j, _ := json.Marshal(raw)
	fmt.Fprintln(Writer, string(j))
}

type metric struct {
	Value float64
	Unit  unit.Unit
}

func Metric(value float64, unit unit.Unit) interface{} {
	return &metric{Value: value, Unit: unit}
}

type dimension struct {
	Value string
}

func Dimension(value string) interface{} {
	return &dimension{Value: value}
}

type metadata struct {
	Timestamp         int
	LogGroupName      string `json:",omitempty"`
	CloudWatchMetrics []cwMetricDirective
}

type cwMetricDirective struct {
	Namespace  string
	Dimensions [][]string
	Metrics    []cwMetricDefinition
}

type cwMetricDefinition struct {
	Name string
	Unit string
}
