package agent

import (
	"fmt"
	"github.com/2er0/GoGp/iface"
	"github.com/2er0/GoGp/utils"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func (e *BaseES) Init() {

	// init the given pieces
	// - Mutator
	// - MutChanger
	// - Selection
	e.mut.SetUp(e.check, e.gen)
	e.mutUpd.SetLambda(e.lambda)
	e.sel.SetUp(e.mu, e.popComp)

	// generate a multiple size of parents
	var pop []iface.Solution
	for i := 0; i < e.mu*3; i++ {
		sol := e.gen.New(e.entitySize)
		sol.SetValues(e.check.CheckAndFix(sol))
		sol.SetScore(e.fit.Calc(sol))
		pop = append(pop, sol)
	}

	// sort parent candidates
	sort.Slice(pop, func(i, j int) bool {
		return e.popComp.Comp(pop[i].GetScore(), pop[j].GetScore())
	})

	// select best parents
	e.pop = e.sel.Selection(pop[:1], pop[1:])
}

// endrun saves the taken time for execution
func (e *BaseES) endrun(start time.Time, status chan interface{}) {
	e.duration = utils.TimeTrack(start, fmt.Sprintf("%v", e), e.genCount, status)
}

// Run executes the agent
func (e *BaseES) Run(status chan interface{}, intClose bool) {
	if intClose {
		// if run one only - send close signal
		defer close(status)
	}
	// take time and save it
	defer e.endrun(time.Now(), status)

	status <- fmt.Sprintf("%v Starting execution", e)

	// run n generations
	for i := 0; i < e.genCount; i++ {

		if e.sigma < 0.0000001 {
			status <- fmt.Sprintf("%s Generation: %6d | Sigma to low", e, i+1)
			break
		}

		// create wait group for every mutation
		wg := &sync.WaitGroup{}
		wg.Add(e.lambda)

		// create communication channel for
		// created childes
		runInfo := make(chan resValue)

		// create lambda new children by mutation of random
		// selected parents
		for j := 0; j < e.lambda; j++ {
			item := rand.Intn(e.mu)
			// run each mutation in a goroutine
			go func(wg *sync.WaitGroup, item int, runInfo chan resValue) {
				defer wg.Done()
				par := e.pop[item]
				// make mutation and make validation of it
				sol := e.mut.Mut(par, e.sigma)

				// calculate score
				sol.SetScore(e.fit.Calc(sol))

				res := resValue{}

				// check if it is better then it's parent
				if e.popComp.Comp(par.GetScore(), sol.GetScore()) {
					res.suc = false
				} else {
					res.suc = true
				}

				res.sol = sol

				runInfo <- res

			}(wg, item, runInfo)
		}

		// run sync in own goroutine
		go monitorWorker(wg, runInfo)

		var childes []iface.Solution
		sucCount := 0

		// collect every child and wait for closing of channel
		for res := range runInfo {
			childes = append(childes, res.sol)
			if res.suc {
				sucCount++
			}
		}

		// select new population
		newPop := e.sel.Selection(e.pop, childes)

		e.pop = newPop

		// update sigma
		e.sigma = e.mutUpd.Update(sucCount, e.sigma)

		status <- fmt.Sprintf("%s Generation: %6d | best: %v | sigma: %v\n",
			e, i+1, e.pop[0], e.sigma)
	}
}

// Result returns the last status from the agent and it's process
func (e *BaseES) Result() string {
	return fmt.Sprintf("%s Gens: %6d in %s | Best: %v", e, e.genCount, e.duration, e.pop[0])
}

// String returns the std information about the agent
func (e *BaseES) String() string {
	return fmt.Sprintf("%s:%s |", e.name, e.id)
}
