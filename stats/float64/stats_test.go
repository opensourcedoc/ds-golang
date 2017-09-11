package stats

import (
	"github.com/cwchentw/algo-golang/vector/float64"
	"testing"
)

func TestVectorMedianOddValues(t *testing.T) {
	v := vector.New(4, 2, 3, 5, 1)
	median := Median(v)

	if !(median == 3) {
		t.Error("Wrong value")
	}
}

func TestVectorMedianEvenValues(t *testing.T) {
	v := vector.New(4, 2, 6, 3, 5, 1)
	median := Median(v)

	if !(median == 3.5) {
		t.Log(median)
		t.Error("Wrong value")
	}
}

func TestVectorMode(t *testing.T) {
	t.Parallel()

	v := vector.New(4, 2, 4, 3, 3, 1, 2, 4, 5, 6, 3)

	mode := Mode(v)

	if !(mode.GetAt(0) == 3) {
		t.Error("Wrong value")
	}

	if !(mode.GetAt(1) == 4) {
		t.Log(mode.GetAt(1))
		t.Error("Wrong value")
	}
}
