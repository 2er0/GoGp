package generator

import (
	"github.com/2er0/GoGp/iface"
	"github.com/2er0/GoGp/solution"
	"math/rand"
)

// A Float creates new float solutions
type Float struct {
}

func NewFloatGen() iface.Generator {
	return &Float{}
}

// New returns a new solution with random values
func (fg *Float) New(size int) iface.Solution {
	nr := solution.NewFloat(size)
	var data []float32
	for i := 0; i < size; i++ {
		data = append(data, float32(rand.NormFloat64())*float32(100))
	}

	nr.SetValues(data)

	return nr
}

// NewClean returns a new solution with no random values
func (fg *Float) NewClean(size int) iface.Solution {
	return solution.NewFloat(size)
}
