package distance

import (
	"errors"
	"github.com/ALTree/bigfloat"
	v "github.com/cwchentw/algo-golang/vector"
	"math"
	"math/big"
	"reflect"
)

func Minkowski(p *v.Vector, q *v.Vector, n int) (interface{}, error) {
	if n <= 0 {
		panic("Exp should be larger than zero")
	}

	checkLength(p, q)

	s, err := p.Sub(q)
	if err != nil {
		return nil, err
	}

	_len := s.Len()
	vec := v.WithSize(_len)
	for i := 0; i < _len; i++ {
		e := s.GetAt(i)
		ts := reflect.TypeOf(e).String()

		switch ts {
		case "int":
			num := e.(int)
			if num < 0 {
				num = -num
			}

			vec.SetAt(i, math.Pow(float64(num), 1.0/float64(n)))
		case "float64":
			num := e.(float64)
			if num < 0 {
				num = -num
			}

			vec.SetAt(i, math.Pow(num, 1/float64(n)))
		case reflect.TypeOf(big.NewInt(0)).String():
			num := e.(*big.Int)
			num = num.Abs(num)

			result := big.NewInt(1)
			for j := 0; j < n; j++ {
				result = result.Mul(result, num)
			}
			vec.SetAt(i, result)
		case reflect.TypeOf(big.NewFloat(0.0)).String():
			num := e.(*big.Float)
			num = num.Abs(num)

			result := big.NewFloat(1.0)
			for j := 0; j < n; j++ {
				result = bigfloat.Pow(result, big.NewFloat(1.0).Quo(big.NewFloat(1.0), num))
			}
			vec.SetAt(i, result)
		default:
			return nil, errors.New("Unknown Type")
		}
	}

	out, err := vec.ReduceBy(func(a interface{}, b interface{}) (interface{}, error) {
		ta := reflect.TypeOf(a).String()
		tb := reflect.TypeOf(b).String()

		if !(ta == "float64" || ta == "int") &&
			!(tb == "float64" || tb == "int") {
			if ta != tb {
				return nil, errors.New("Unequal Type")
			}
		}

		switch ta {
		case "int":
			switch tb {
			case "int":
				na := a.(int)
				nb := b.(int)
				return na + nb, nil
			case "float64":
				na := a.(int)
				nb := b.(float64)
				return float64(na) + nb, nil
			default:
				return nil, errors.New("Unknown Type")
			}
		case "float64":
			switch tb {
			case "int":
				na := a.(float64)
				nb := b.(int)
				return na + float64(nb), nil
			case "float64":
				na := a.(float64)
				nb := b.(float64)
				return na + nb, nil
			default:
				return nil, errors.New("Unknown Type")
			}
		case reflect.TypeOf(big.NewInt(0)).String():
			na := a.(*big.Int)
			nb := b.(*big.Int)
			return na.Add(na, nb), nil
		case reflect.TypeOf(big.NewFloat(0.0)).String():
			na := a.(*big.Float)
			nb := b.(*big.Float)
			return na.Add(na, nb), nil
		default:
			return nil, errors.New("Unknown Type")
		}
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}

func checkLength(p *v.Vector, q *v.Vector) {
	if p.Len() != q.Len() {
		panic("Unequal Length")
	}
}
