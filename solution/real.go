package solution

import (
	"fmt"
	"github.com/2er0/GoGp/iface"
	"reflect"
)

type Real struct {
	value []int
	size  int
	score float32
}

func NewReal(size int) iface.Solution {
	return &Real{value: make([]int, size), size: size, score: -1}
}

func (r *Real) GetValues() interface{} {
	return r.value
}

func (r *Real) SetValues(d interface{}) {
	s := reflect.ValueOf(d)
	if s.Kind() != reflect.Slice {
		panic("SetValues() given a non-slice type")
	}

	if s.Len() != r.size {
		panic("SetValues() given slice has not the required length")
	}

	if s.Index(0).Kind() != reflect.Int {
		panic("SetValues() given a non int slice type")
	}

	var data []int
	for i := 0; i < s.Len(); i++ {
		data = append(data, int(s.Index(i).Int()))
	}

	r.value = data
}

func (r *Real) GetValue(i int) interface{} {
	return r.value[i]
}

func (r *Real) SetValue(i int, v interface{}) {
	s := reflect.ValueOf(v)
	if s.Kind() != reflect.Int {
		panic("SetValue() given a non int value")
	}

	if i > r.size {
		panic("SetValue() given index out of range")
	}

	r.value[i] = int(s.Int())
}

func (r *Real) GetScore() interface{} {
	return r.score
}

func (r *Real) SetScore(d interface{}) {
	s := reflect.ValueOf(d)
	if s.Kind() != reflect.Float32 && s.Kind() != reflect.Float64 && s.Kind() != reflect.Int {
		panic("SetScore() given a non-numeric type")
	}

	r.score = float32(s.Float())
}

func (r *Real) GetSize() int {
	return r.size
}

func (r *Real) Copy() iface.Solution {
	c := NewReal(r.GetSize())
	c.SetValues(r.GetValues())
	c.SetScore(r.GetScore())
	return c
}

func (r *Real) String() string {
	return fmt.Sprintf("%v -> %v", r.score, r.value)
}
