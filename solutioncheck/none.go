package solutioncheck

import (
	"github.com/2er0/GoGp/iface"
)

// A None solutioncheck does nothing with a solution candidate,
// for a function where every number is valid, is nothing to change
type None struct {
}

func NewNone() iface.SolutionCheck {
	return &None{}
}

func (n *None) CheckAndFix(sol iface.Solution) interface{} {
	return sol.GetValues()
}
