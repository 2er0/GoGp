package test

import (
	"github.com/2er0/GoGp/agent"
	"github.com/2er0/GoGp/engine"
	"testing"
)

func TestBase(t *testing.T) {
	var e = *engine.New()

	a1 := agent.NewTest()
	a2 := agent.NewTest()
	e.Add(a1)
	e.Add(a2)
	e.Run()
	e.Run()
}
