package generator

import (
	"github.com/2er0/GoGp/iface"
	"github.com/2er0/GoGp/solution"
	"math/rand"
)

// A Real creates new real solutions
type Real struct {
}

func NewRealGen() iface.Generator {
	return &Real{}
}

func (rg *Real) New(size int) iface.Solution {
	nr := solution.NewReal(size)
	var data []int
	for i := 0; i < size; i++ {
		value := int(rand.NormFloat64() * float64(100))

		data = append(data, value)
	}

	nr.SetValues(data)

	return nr
}

func (rg *Real) NewClean(size int) iface.Solution {
	return solution.NewReal(size)
}
