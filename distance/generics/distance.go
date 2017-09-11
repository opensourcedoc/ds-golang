package distance

import (
	"errors"
	"github.com/ALTree/bigfloat"
	v "github.com/cwchentw/algo-golang/vector/generics"
	"math"
	"math/big"
	"reflect"
)

func Euclidean(p *v.Vector, q *v.Vector) (interface{}, error) {
	return Minkowski(p, q, 2)
}

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
			_num := new(big.Float).SetInt(num)

			result := big.NewFloat(1.0)
			result = bigfloat.Pow(result, big.NewFloat(1.0).Quo(big.NewFloat(1.0), _num))
			vec.SetAt(i, result)
		case reflect.TypeOf(big.NewFloat(0.0)).String():
			num := e.(*big.Float)
			num = num.Abs(num)

			result := big.NewFloat(1.0)
			result = bigfloat.Pow(result, big.NewFloat(1.0).Quo(big.NewFloat(1.0), num))
			vec.SetAt(i, result)
		default:
			return nil, errors.New("Unknown Type")
		}
	}

	out, err := vec.Reduce(add)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func Maximum(p *v.Vector, q *v.Vector) (interface{}, error) {
	return Chebyshev(p, q)
}

func Chebyshev(p *v.Vector, q *v.Vector) (interface{}, error) {
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

			vec.SetAt(i, float64(num))
		case "float64":
			num := e.(float64)
			if num < 0 {
				num = -num
			}

			vec.SetAt(i, num)
		case reflect.TypeOf(big.NewInt(0)).String():
			num := e.(*big.Int)
			num = num.Abs(num)
			vec.SetAt(i, num)
		case reflect.TypeOf(big.NewFloat(0.0)).String():
			num := e.(*big.Float)
			num = num.Abs(num)
			vec.SetAt(i, num)
		default:
			return nil, errors.New("Unknown Type")
		}
	}

	out, err := vec.Reduce(max)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func Manhattan(p *v.Vector, q *v.Vector) (interface{}, error) {
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

			vec.SetAt(i, float64(num))
		case "float64":
			num := e.(float64)
			if num < 0 {
				num = -num
			}

			vec.SetAt(i, num)
		case reflect.TypeOf(big.NewInt(0)).String():
			num := e.(*big.Int)
			num = num.Abs(num)
			vec.SetAt(i, num)
		case reflect.TypeOf(big.NewFloat(0.0)).String():
			num := e.(*big.Float)
			num = num.Abs(num)
			vec.SetAt(i, num)
		default:
			return nil, errors.New("Unknown Type")
		}
	}

	out, err := vec.Reduce(add)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func Canberra(p *v.Vector, q *v.Vector) (interface{}, error) {
	checkLength(p, q)

	pAbs, err := p.Map(abs)
	if err != nil {
		return nil, err
	}

	qAbs, err := q.Map(abs)
	if err != nil {
		return nil, err
	}

	s, err := p.Sub(q)
	if err != nil {
		return nil, err
	}

	sAbs, err := s.Map(abs)
	if err != nil {
		return nil, err
	}

	_len := s.Len()
	vec := v.WithSize(_len)
	for i := 0; i < _len; i++ {
		a := sAbs.GetAt(i)
		ts := reflect.TypeOf(a).String()

		b := pAbs.GetAt(i)
		c := qAbs.GetAt(i)
		switch ts {
		case "int":
			na := float64(a.(int))
			nb := float64(b.(int))
			nc := float64(c.(int))
			vec.SetAt(i, na/(nb+nc))
		case "float64":
			na := a.(float64)
			nb := b.(float64)
			nc := c.(float64)
			vec.SetAt(i, na/(nb+nc))
		case reflect.TypeOf(big.NewInt(0)).String():
			na := new(big.Float).SetInt(a.(*big.Int))
			nb := new(big.Float).SetInt(b.(*big.Int))
			nc := new(big.Float).SetInt(c.(*big.Int))
			vec.SetAt(i, na.Quo(na, nb.Add(nb, nc)))
		case reflect.TypeOf(big.NewFloat(0.0)).String():
			na := a.(*big.Float)
			nb := b.(*big.Float)
			nc := c.(*big.Float)
			vec.SetAt(i, na.Quo(na, nb.Add(nb, nc)))
		default:
			return nil, errors.New("Unknown Type")
		}
	}

	out, err := vec.Reduce(add)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func add(a interface{}, b interface{}) (interface{}, error) {
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
			return float64(na + nb), nil
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
}

func max(a interface{}, b interface{}) (interface{}, error) {
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
			if na > nb {
				return na, nil
			} else {
				return nb, nil
			}
		case "float64":
			na := float64(a.(int))
			nb := b.(float64)
			if na > nb {
				return na, nil
			} else {
				return nb, nil
			}
		default:
			return nil, errors.New("Unknown Type")
		}
	case "float64":
		switch tb {
		case "int":
			na := a.(float64)
			nb := float64(b.(int))
			if na > nb {
				return na, nil
			} else {
				return nb, nil
			}
		case "float64":
			na := a.(float64)
			nb := b.(float64)
			if na > nb {
				return na, nil
			} else {
				return nb, nil
			}
		default:
			return nil, errors.New("Unknown Type")
		}
	case reflect.TypeOf(big.NewInt(0)).String():
		na := a.(*big.Int)
		nb := b.(*big.Int)
		if na.Cmp(nb) > 0 {
			return na, nil
		} else {
			return nb, nil
		}
	case reflect.TypeOf(big.NewFloat(0.0)).String():
		na := a.(*big.Float)
		nb := b.(*big.Float)
		if na.Cmp(nb) > 0 {
			return na, nil
		} else {
			return nb, nil
		}
	default:
		return nil, errors.New("Unknown Type")
	}
}

func abs(a interface{}) (interface{}, error) {
	ta := reflect.TypeOf(a).String()

	switch ta {
	case "int":
		na := a.(int)
		if na > 0 {
			return na, nil
		} else {
			return -na, nil
		}
	case "float64":
		na := a.(float64)
		if na > 0 {
			return na, nil
		} else {
			return -na, nil
		}
	case reflect.TypeOf(big.NewInt(0)).String():
		na := a.(*big.Int)
		return na.Abs(na), nil
	case reflect.TypeOf(big.NewFloat(0.0)).String():
		na := a.(*big.Float)
		return na.Abs(na), nil
	default:
		return nil, errors.New("Unknown Type")
	}
}

func checkLength(p *v.Vector, q *v.Vector) {
	if p.Len() != q.Len() {
		panic("Unequal Length")
	}
}
