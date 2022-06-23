package node

import (
	"fmt"
	"strconv"
	"unsafe"
)

type DataType int

const (
	DataTypeInt DataType = iota + 1
	DataTypeFloat
	DataTypeVariable
)

type Env map[string]any

type Closures func(Env) Number

type Number struct {
	datatype DataType
	value    [8]byte
}

func NewInt(v int) Number {
	n := Number{}
	n.SetInt(v)
	return n
}
func NewFloat(v float64) Number {
	n := Number{}
	n.SetFloat(v)
	return n
}

func (n Number) String() string {
	if n.datatype == DataTypeInt {
		return strconv.Itoa(n.GetInt())
	} else {
		return strconv.FormatFloat(n.GetFloat(), 'f', 10, 64)
	}
}

func (n *Number) SetInt(v int) {
	n.datatype = DataTypeInt
	*(*int64)(unsafe.Pointer(&n.value[0])) = int64(v)
}

func (n Number) GetInt() int {

	if n.datatype == DataTypeInt {
		return int(*(*int64)(unsafe.Pointer(&n.value[0])))
	} else {
		vf := *(*float64)(unsafe.Pointer(&n.value[0]))
		return int(vf)
	}
}
func (n Number) GetFloat() float64 {
	if n.datatype == DataTypeFloat {
		return *(*float64)(unsafe.Pointer(&n.value[0]))
	} else {
		vf := *(*int64)(unsafe.Pointer(&n.value[0]))
		return float64(vf)
	}
}
func (n *Number) SetFloat(v float64) {
	n.datatype = DataTypeFloat
	*(*float64)(unsafe.Pointer(&n.value[0])) = v
}

func (n *Number) CastToInt() int {
	if n.datatype == DataTypeInt {
		return int(*(*int64)(unsafe.Pointer(&n.value[0])))
	}
	n.datatype = DataTypeInt
	v := *(*float64)(unsafe.Pointer(&n.value[0]))
	*(*int64)(unsafe.Pointer(&n.value[0])) = int64(v)
	return int(v)
}

func (n *Number) CastToFloat() float64 {
	if n.datatype == DataTypeFloat {
		return *(*float64)(unsafe.Pointer(&n.value[0]))
	}
	n.datatype = DataTypeFloat
	v := float64(*(*int64)(unsafe.Pointer(&n.value[0])))
	*(*float64)(unsafe.Pointer(&n.value[0])) = v

	return v
}

func (n Number) Add(r Number) Number {
	if n.datatype == DataTypeFloat || r.datatype == DataTypeFloat {
		v := n.CastToFloat() + r.CastToFloat()
		return NewFloat(v)
	} else {
		return NewInt(n.GetInt() + r.GetInt())
	}
}
func (n Number) minus(r Number) Number {
	if n.datatype == DataTypeFloat || r.datatype == DataTypeFloat {
		v := n.CastToFloat() - r.CastToFloat()
		return NewFloat(v)
	} else {
		return NewInt(n.GetInt() - r.GetInt())
	}
}

func (n Number) multiplication(r Number) Number {
	if n.datatype == DataTypeFloat || r.datatype == DataTypeFloat {
		v := n.CastToFloat() * r.CastToFloat()
		return NewFloat(v)
	} else {
		return NewInt(n.GetInt() * r.GetInt())
	}
}
func (n Number) division(r Number) Number {
	if n.datatype == DataTypeFloat || r.datatype == DataTypeFloat {
		v := n.CastToFloat() / r.CastToFloat()
		return NewFloat(v)
	} else {
		fmt.Println("/", n.GetInt(), r.GetInt())
		return NewInt(n.GetInt() / r.GetInt())
	}
}
