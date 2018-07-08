package fitness

import (
	"github.com/2er0/GoGp/iface"
	"reflect"
)

// TestFunc contains a simple x^2 + 1 test function
type TestFunc struct {
	count int
}

// NewTestFunc returns an instance of TestFunc
func NewTestFunc(count int) iface.Fitness {
	return &TestFunc{count: count}
}

func (f *TestFunc) Calc(sol iface.Solution) interface{} {
	if f.count != sol.GetSize() {
		panic("Sizes does not match")
	}

	var value float32
	val := reflect.ValueOf(sol.GetValues())
	for i := 0; i < val.Len(); i++ {
		value = float32(val.Index(i).Float())
	}

	res := value*value + 1

	return res
}

func (f *TestFunc) SetData(d interface{}) {

}
