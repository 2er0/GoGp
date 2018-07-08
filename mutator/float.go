package mutator

import (
	"github.com/2er0/GoGp/iface"
	"math/rand"
)

// A Float mutator mutates all values by floats
type Float struct {
	check iface.SolutionCheck
	gen   iface.Generator
}

func NewFloat() iface.Mutator {
	return &Float{check: nil, gen: nil}
}

func (f *Float) Mut(sol iface.Solution, sigma float32) iface.Solution {
	if !f.Check() {
		panic("No solution checker given")
	}

	res := f.gen.NewClean(sol.GetSize())

	size := sol.GetSize()

	for i := 0; i < size; i++ {
		res.SetValue(i, f.mutIt(sol.GetValue(i).(float32), sigma))
	}

	res.SetValues(f.check.CheckAndFix(res))

	return res
}

func (f *Float) SetUp(check iface.SolutionCheck, gen iface.Generator) {
	f.check = check
	f.gen = gen
}

func (f *Float) Check() bool {
	return f.check != nil
}

func (f *Float) mutIt(i float32, sigma float32) float32 {
	return float32(rand.NormFloat64())*sigma + i
}

func (f *Float) MutOne(i interface{}, sigma float32) interface{} {
	return f.mutIt(i.(float32), sigma)
}
