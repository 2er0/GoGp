package agent

import (
	"github.com/2er0/GoGp/iface"
	"github.com/2er0/GoGp/utils"
)

// A BaseGA represents the base struct for an GA agent
// it contains every thing it needs to run the GA algorithm
// but does not know about the problem domain
// and is base on the BaseES struct
type BaseGA struct {
	*BaseES
	xo        iface.Crossover
	offSpring bool
}

// NewGA returns a new instance of a BaseGA with build in BaseES
func NewGA(mu, lambda, genCount, entitySize int, ofs bool,
	f iface.Fitness,
	g iface.Generator,
	xo iface.Crossover,
	m iface.Mutator,
	s iface.Selection,
	mc iface.MutChanger,
	c iface.SolutionCheck,
	comp iface.Comp) iface.Task {
	uuid := utils.PseudoUuid()
	name := "GA"

	newGA := &BaseGA{
		BaseES:    newES(mu, lambda, genCount, entitySize, f, g, m, s, mc, c, comp, uuid, name),
		xo:        xo,
		offSpring: ofs,
	}

	return newGA
}
