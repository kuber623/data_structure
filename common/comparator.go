package common

import (
	"reflect"
	"time"
)

const (
	greater = 1
	lesser  = -1
	equal   = 0
)

type Comparator func(a, b interface{}) int

func NumberComparator(a, b interface{}) int {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	if va.Kind() != vb.Kind() {
		panic("number comparator: value type not match")
	}
	switch va.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		inta, intb := va.Int(), vb.Int()
		switch {
		case inta > intb:
			return greater
		case inta < intb:
			return lesser
		default:
			return equal
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uinta, uintb := va.Uint(), vb.Uint()
		switch {
		case uinta > uintb:
			return greater
		case uinta < uintb:
			return lesser
		default:
			return equal
		}
	case reflect.Float32, reflect.Float64:
		floata, floatb := va.Float(), vb.Float()
		switch {
		case floata > floatb:
			return greater
		case floata < floatb:
			return lesser
		default:
			return equal
		}
	default:
		panic("number comparator: incomparable value type")
	}
}

func StringComparator(a, b interface{}) int {
	stra := a.(string)
	strb := b.(string)
	min := len(strb)
	if len(stra) < len(strb) {
		min = len(stra)
	}
	diff := uint8(0)
	for i := 0; i < min && diff == 0; i++ {
		diff = stra[i] - strb[i]
	}
	if diff == 0 {
		switch {
		case len(stra) > len(strb):
			return greater
		case len(stra) < len(strb):
			return lesser
		default:
			return equal
		}
	}
	if diff > 0 {
		return greater
	} else {
		return lesser
	}
}

func TimeComparator(a, b interface{}) int {
	timea := a.(time.Time)
	timeb := b.(time.Time)
	switch {
	case timea.After(timeb):
		return 1
	case timea.Before(timeb):
		return -1
	default:
		return 0
	}
}
