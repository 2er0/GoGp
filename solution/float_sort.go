package solution

import "github.com/2er0/GoGp/iface"

// A FloatComp compares the float values
// like a float score or else
type FloatComp struct {
	s1 interface{}
	s2 interface{}
}

func NewFloatCompMin() iface.Comp {
	return &FloatComp{nil, nil}
}

func (fs *FloatComp) Less() bool {
	return fs.Comp(fs.s1, fs.s2)
}

func (fs *FloatComp) Comp(i, j interface{}) bool {
	return i.(float32) < j.(float32)
}
