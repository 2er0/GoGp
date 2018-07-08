package solutioncheck

import (
	"github.com/2er0/GoGp/iface"
)

// A Tsp solutioncheck take a solution checks it
// for duplicates and missing stopps
type Tsp struct {
}

func NewTsp() iface.SolutionCheck {
	return &Tsp{}
}

func (t *Tsp) CheckAndFix(sol iface.Solution) interface{} {

	size := sol.GetSize()

	found := map[int]int{}
	double := map[int]bool{}

	for i := 0; i < size; i++ {
		found[i] = 0
		double[i] = false
	}

	for i := 0; i < size; i++ {
		city := sol.GetValue(i).(int)
		found[city] = found[city] + 1
	}

	var missing []int

	for i, v := range found {
		if v < 1 {
			missing = append(missing, i)
		}
	}

	for i := 0; i < size; i++ {

		item := sol.GetValue(i).(int)

		if found[item] > 1 && double[item] {

			missed := -1
			missed, missing = popFirst(missing)

			sol.SetValue(i, missed)
		}

		double[item] = true
	}

	misleng := len(missing)
	if misleng > 0 {
		for i := size - misleng; i < size; i++ {
			missed := -1
			missed, missing = popFirst(missing)

			sol.SetValue(i, missed)
		}
	}

	return sol.GetValues()
}

func popFirst(missing []int) (int, []int) {
	missed := missing[0]

	copy(missing, missing[1:])
	missing = missing[:len(missing)-1]
	return missed, missing
}
