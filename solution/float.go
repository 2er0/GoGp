package solution

import (
	"fmt"
	"github.com/2er0/GoGp/iface"
	"reflect"
)

type Float struct {
	value []float32
	size  int
	score float32
}

func NewFloat(size int) iface.Solution {
	return &Float{value: make([]float32, size), size: size, score: -1}
}

func (f *Float) GetValues() interface{} {
	return f.value
}

func (f *Float) SetValues(d interface{}) {
	s := reflect.ValueOf(d)
	if s.Kind() != reflect.Slice {
		panic("SetValue() given a non-slice type")
	}

	if s.Len() != f.size {
		panic("SetValue() given slice has not the required length")
	}

	if s.Index(0).Kind() != reflect.Float32 {
		panic("SetValue() given a non float32 slice type")
	}

	var data []float32
	for i := 0; i < s.Len(); i++ {
		data = append(data, float32(s.Index(i).Float()))
	}

	f.value = data
}

func (f *Float) GetValue(i int) interface{} {
	return f.value[i]
}

func (f *Float) SetValue(i int, v interface{}) {
	s := reflect.ValueOf(v)
	if s.Kind() != reflect.Float32 && s.Kind() != reflect.Float64 {
		panic("SetValue() given a non float value")
	}

	if i > f.size {
		panic("SetValue() given index out of range")
	}

	f.value[i] = float32(s.Float())
}

func (f *Float) GetScore() interface{} {
	return f.score
}

func (f *Float) SetScore(d interface{}) {
	s := reflect.ValueOf(d)
	if s.Kind() != reflect.Float32 && s.Kind() != reflect.Float64 {
		panic("SetScore() given a non-float type")
	}

	f.score = float32(s.Float())
}

func (f *Float) GetSize() int {
	return f.size
}

func (f *Float) Copy() iface.Solution {
	c := NewFloat(f.size)
	c.SetValues(f.GetValues())
	c.SetScore(f.GetScore())
	return c
}

func (f *Float) String() string {
	return fmt.Sprintf("%v -> %v", f.score, f.value)
}
