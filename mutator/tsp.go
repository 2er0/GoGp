package mutator

import (
	"github.com/2er0/GoGp/iface"
	"math"
	"math/rand"
)

// A Tsp Mutator is special
// it needs the count of stop
// to be in the upper bound
type Tsp struct {
	check iface.SolutionCheck
	gen   iface.Generator
	size  int
}

func NewTsp(size int) iface.Mutator {
	return &Tsp{check: nil, gen: nil, size: size}
}

func (t *Tsp) Mut(sol iface.Solution, sigma float32) iface.Solution {
	if !t.Check() {
		panic("No solution checker given")
	}

	res := t.gen.NewClean(sol.GetSize())

	size := sol.GetSize()

	for i := 0; i < size; i++ {
		res.SetValue(i, t.mutIt(sol.GetValue(i).(int), sigma))
	}

	res.SetValues(t.check.CheckAndFix(res))

	return res
}

func (t *Tsp) SetUp(check iface.SolutionCheck, gen iface.Generator) {
	t.check = check
	t.gen = gen
}

func (t *Tsp) Check() bool {
	return t.check != nil
}

// mutIt mutates one value in the bound from 0 to stop count
func (t *Tsp) mutIt(i int, sigma float32) int {
	return int(math.Abs(float64(int(float32(rand.NormFloat64())*sigma)+i))) % t.size
}

func (t *Tsp) MutOne(i interface{}, sigma float32) interface{} {
	return t.mutIt(i.(int), sigma)
}
