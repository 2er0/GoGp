package mutator

import (
	"github.com/2er0/GoGp/iface"
	"math/rand"
)

type Real struct {
	check iface.SolutionCheck
	gen   iface.Generator
}

func NewReal() iface.Mutator {
	return &Real{check: nil, gen: nil}
}

func (r *Real) Mut(sol iface.Solution, sigma float32) iface.Solution {
	if !r.Check() {
		panic("No solution checker given")
	}

	res := r.gen.NewClean(sol.GetSize())

	size := sol.GetSize()

	for i := 0; i < size; i++ {
		res.SetValue(i, r.mutIt(sol.GetValue(i).(int), sigma))
	}

	res.SetValues(r.check.CheckAndFix(res))

	return res
}

func (r *Real) SetUp(check iface.SolutionCheck, gen iface.Generator) {
	r.check = check
	r.gen = gen
}

func (r *Real) Check() bool {
	return r.check != nil
}

func (r *Real) mutIt(i int, sigma float32) int {
	return int(float32(rand.NormFloat64())*sigma) + i
}

func (r *Real) MutOne(i interface{}, sigma float32) interface{} {
	return r.mutIt(i.(int), sigma)
}
