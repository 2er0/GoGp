package crossover

import (
	"github.com/2er0/GoGp/iface"
	"math/rand"
	"reflect"
)

// A Pattern crossover
type Pattern struct {
	check iface.SolutionCheck
	gen   iface.Generator
}

// NewPattern returns an instance of a new Pattern crossover
func NewPattern() iface.Crossover {
	return &Pattern{check: nil, gen: nil}
}

// XO executes the crossover on two parent solutions with
// a random pattern and returns a new child
func (p *Pattern) XO(s1, s2 iface.Solution) iface.Solution {
	if !p.Check() {
		panic("No solution checker given")
	}

	size := s1.GetSize()
	v1 := reflect.ValueOf(s1.GetValues())
	v2 := reflect.ValueOf(s2.GetValues())

	if v1.Kind() != reflect.Slice || v2.Kind() != reflect.Slice {
		panic("XO() given solution don't have slices")
	}

	if v1.Len() != s2.GetSize() || v2.Len() != s1.GetSize() {
		panic("XO() given solutions doesn't have the same size")
	}

	if v1.Index(0).Kind() != v2.Index(0).Kind() {
		panic("XO() given slices are not matching")
	}

	res := p.gen.NewClean(s1.GetSize())

	// create pattern by random permutation
	pattern := rand.Perm(size)

	// set child data
	for i := 0; i < size; i++ {
		if pattern[i]%2 > 0 {
			res.SetValue(i, s1.GetValue(i))
		} else {
			res.SetValue(i, s2.GetValue(i))
		}
	}

	res.SetValues(p.check.CheckAndFix(res))

	return res
}

// SetUp loads everything to be executable
func (p *Pattern) SetUp(c iface.SolutionCheck, gen iface.Generator) {
	p.check = c
	p.gen = gen
}

func (p *Pattern) Check() bool {
	return p.check != nil
}
