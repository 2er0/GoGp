package agent

import (
	"log"
	"time"
)

// A Base represents the base struct for an agent
// it stores only the basic information
type Base struct {
	id       string
	name     string
	log      *log.Logger
	duration time.Duration
}

// NewBase returns a new instance of a Base
func NewBase(id string, name string, log *log.Logger) *Base {
	return &Base{id: id, name: name, log: log, duration: 0}
}
