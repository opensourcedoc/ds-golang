package distance

import (
	vec "github.com/cwchentw/algo-golang/vector"
	"math"
	"testing"
)

func TestEuclidean(t *testing.T) {
	v1 := vec.New(0, 1, 2)
	v2 := vec.New(3, 4, 5)

	d, _ := Euclidean(v1, v2)
	if !(math.Abs(d.(float64)-5.196152) < 1.0/1000000) {
		t.Log(d.(float64))
		t.Error("Wrong distance value")
	}
}

func TestMinkowski(t *testing.T) {
	v1 := vec.New(0, 1, 2)
	v2 := vec.New(3, 4, 5)

	d, _ := Minkowski(v1, v2, 3)
	if !(math.Abs(d.(float64)-4.326749) < 1.0/1000000) {
		t.Log(d.(float64))
		t.Error("Wrong distance value")
	}
}
