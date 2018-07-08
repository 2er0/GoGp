package generator

import (
	"github.com/2er0/GoGp/iface"
	"github.com/2er0/GoGp/solution"
	"math/rand"
)

// A Tsp creates new tsp solutions
type Tsp struct {
}

func NewTsp() iface.Generator {
	return &Tsp{}
}

func (t *Tsp) New(size int) iface.Solution {
	nr := solution.NewReal(size)
	nr.SetValues(rand.Perm(size))

	return nr
}

func (t *Tsp) NewClean(size int) iface.Solution {
	nr := solution.NewReal(size)
	for i := 0; i < size; i++ {
		nr.SetValue(i, i)
	}
	return nr
}
