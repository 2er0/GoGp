package crossover

import (
	"github.com/2er0/GoGp/iface"
	"math/rand"
	"reflect"
)

// A Simple cut crossover
type Simple struct {
	check iface.SolutionCheck
	gen   iface.Generator
}

// NewSimple returns an instance of a new Simple crossover
func NewSimple() iface.Crossover {
	return &Simple{check: nil, gen: nil}
}

// XO executes the crossover on two parent solutions
// by a random cut simple cut and returns a new childSo
func (s *Simple) XO(s1, s2 iface.Solution) iface.Solution {
	if !s.Check() {
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

	res := s.gen.NewClean(s1.GetSize())

	cut := rand.Intn(size + 1)

	for i := 0; i < size; i++ {
		if i < cut {
			res.SetValue(i, s1.GetValue(i))
		} else {
			res.SetValue(i, s2.GetValue(i))
		}
	}

	res.SetValues(s.check.CheckAndFix(res))

	return res
}

func (s *Simple) SetUp(c iface.SolutionCheck, gen iface.Generator) {
	s.check = c
	s.gen = gen
}

func (s *Simple) Check() bool {
	return s.check != nil
}
