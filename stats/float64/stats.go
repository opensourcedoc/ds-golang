package stats

import (
	"github.com/cwchentw/algo-golang/vector/float64"
	"strconv"
)

func Sum(v vector.IVector) float64 {
	sum := 0.0

	for i := 0; i < v.Len(); i++ {
		sum += v.GetAt(i)
	}

	return sum
}

func Mean(v vector.IVector) float64 {
	return Sum(v) / float64(v.Len())
}

func Median(v vector.IVector) float64 {
	if v.Len() == 0 {
		panic("No valid vector data")
	} else if v.Len() == 1 {
		return v.GetAt(0)
	}

	sorted := v.Sort()
	_len := sorted.Len()

	if _len%2 == 0 {
		i := _len / 2
		return (sorted.GetAt(i-1) + sorted.GetAt(i)) / float64(2)
	} else {
		i := _len / 2
		return sorted.GetAt(i)
	}
}

func Mode(v vector.IVector) vector.IVector {
	if v.Len() == 0 {
		panic("No valid vector data")
	} else if v.Len() == 1 {
		return v
	}

	m := make(map[string]int)

	max := -1
	for i := 0; i < v.Len(); i++ {
		f := strconv.FormatFloat(v.GetAt(i), 'E', -1, 64)
		_, ok := m[f]
		if !ok {
			m[f] = 0
		} else {
			m[f]++
		}

		if m[f] > max {
			max = m[f]
		}
	}

	arr := make([]float64, 0)
	for k, _ := range m {
		if m[k] >= max {
			f, _ := strconv.ParseFloat(k, 64)
			arr = append(arr, f)
		}
	}

	return vector.FromArray(arr).Sort()
}
