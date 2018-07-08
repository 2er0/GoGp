package selection

import (
	"github.com/2er0/GoGp/iface"
	"sort"
)

// A Without selection take only the best children
// without the parents
type Without struct {
	mu   int
	comp iface.Comp
}

func NewWithOut() iface.Selection {
	return &Without{mu: 0, comp: nil}
}

func (w *Without) SetUp(mu int, co iface.Comp) {
	w.mu = mu
	w.comp = co
}

func (w *Without) Selection(parents, child []iface.Solution) []iface.Solution {

	sort.Slice(child, func(i, j int) bool {
		return w.comp.Comp(child[i].GetScore(), child[j].GetScore())
	})

	return child[0:w.mu]
}

// A With selection merges the parents and childrens
// and selects the best of all
type With struct {
	mu   int
	comp iface.Comp
}

func NewWith() iface.Selection {
	return &With{mu: 0, comp: nil}
}

func (w *With) SetUp(mu int, co iface.Comp) {
	w.mu = mu
	w.comp = co
}

func (w *With) Selection(parents, child []iface.Solution) []iface.Solution {

	pAndC := append(parents, child...)

	sort.Slice(pAndC[:], func(i, j int) bool {
		return w.comp.Comp(pAndC[i].GetScore(), pAndC[j].GetScore())
	})

	return pAndC[0:w.mu]
}

// A WithoutElit selection takes the best parent
// and selects the best children
type WithoutElit struct {
	mu   int
	comp iface.Comp
}

func NewWithoutElit() iface.Selection {
	return &WithoutElit{mu: 0, comp: nil}
}

func (w *WithoutElit) SetUp(mu int, co iface.Comp) {
	w.mu = mu
	w.comp = co
}

func (w *WithoutElit) Selection(parents, child []iface.Solution) []iface.Solution {

	eAndC := append(child, parents[0])

	sort.Slice(eAndC[:], func(i, j int) bool {
		return w.comp.Comp(eAndC[i].GetScore(), eAndC[j].GetScore())
	})

	return eAndC[0:w.mu]
}
