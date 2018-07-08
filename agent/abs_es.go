package agent

import (
	"github.com/2er0/GoGp/iface"
	"github.com/2er0/GoGp/utils"
	"log"
	"os"
)

// A BaseES represents the base struct for an ES agent
// it contains every thing it needs to run the ES algorithm
// but does not know about the problem domain
// and is base on the Base struct
type BaseES struct {
	*Base
	mu         int
	entitySize int
	genCount   int
	lambda     int
	pop        []iface.Solution
	popComp    iface.Comp
	gen        iface.Generator
	fit        iface.Fitness
	sigma      float32
	mut        iface.Mutator
	mutUpd     iface.MutChanger
	sel        iface.Selection
	check      iface.SolutionCheck
}

// NewES returns a new instance of a BaseES with the build in Base struct
func NewES(mu, lambda, genCount, entitySize int,
	f iface.Fitness,
	g iface.Generator,
	m iface.Mutator,
	s iface.Selection,
	mc iface.MutChanger,
	c iface.SolutionCheck,
	comp iface.Comp) iface.Task {
	uuid := utils.PseudoUuid()
	name := "ES"

	newES := newES(mu, lambda, genCount, entitySize, f, g, m, s, mc, c, comp, uuid, name)

	return newES
}

// newES returns the concrete instance of an created BaseES
func newES(mu, lambda, genCount, entitySize int,
	f iface.Fitness,
	g iface.Generator,
	m iface.Mutator,
	s iface.Selection,
	mc iface.MutChanger,
	c iface.SolutionCheck,
	comp iface.Comp,
	uuid string,
	name string) *BaseES {

	newES := &BaseES{
		Base:       NewBase(uuid, name, log.New(os.Stdout, "", utils.BaseLogInfo())),
		mu:         mu,
		entitySize: entitySize,
		genCount:   genCount,
		lambda:     lambda,
		pop:        nil,
		popComp:    comp,
		gen:        g,
		fit:        f,
		sigma:      1,
		mut:        m,
		mutUpd:     mc,
		check:      c,
		sel:        s,
	}

	return newES
}
