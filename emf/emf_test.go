package emf_test

import (
	"github.com/glassechidna/go-emf/emf"
	"github.com/glassechidna/go-emf/emf/unit"
	"testing"
)

func TestEmit(t *testing.T) {
	emf.Emit(emf.MSI{
		"TableName":  emf.Dimension("mytable"),
		"ItemCount":  emf.Metric(7, unit.Count),
		"TableSize":  emf.Metric(55, unit.Bytes),
		"OtherValue": 55,
	})
}
