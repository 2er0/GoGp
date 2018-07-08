package agent

import (
	"fmt"
	"github.com/2er0/GoGp/utils"
	"log"
	"math/rand"
	"os"
	"time"
)

// A Test represents a test agent for testing the engine
type Test struct {
	*Base
	timeout  int
	execDone bool
}

// NewTest returns a new TestAgent which waits a random time
// and comes back after that time
func NewTest() *Test {
	uuid := utils.PseudoUuid()
	name := "Test"
	return &Test{
		Base:
		//&Base{id: string(uuid), name: name, log: log.New(os.Stdout, "", log.Ldate|log.Ltime)},
		NewBase(string(uuid), name, log.New(os.Stdout, "", utils.BaseLogInfo())),
		timeout:  rand.Intn(100),
		execDone: false}
}

// Init inits the TestAgent
func (a *Test) Init() {
	return
}

// String returns the std information about the agent
func (a *Test) String() string {
	return fmt.Sprintf("%s:%s |", a.name, a.id)
}

// Run executes the agent
func (a *Test) Run(status chan interface{}, intClose bool) {
	// close channel on func exit
	if intClose {
		defer close(status)
	}

	a.log.Printf("%s Start running for %d\n", a, a.timeout)

	time.Sleep(time.Duration(a.timeout) * time.Millisecond)

	a.execDone = true
	status <- fmt.Sprintf("%s status: %t\n", a, a.execDone)
	a.log.Printf("%s End Running\n", a)
}

// Result returns the last status from the agent and it's process
func (a *Test) Result() string {
	a.log.Printf("%s done: %t , runtime: %d\n", a, a.execDone, a.timeout)
	return a.String()
}
