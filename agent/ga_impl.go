package agent

import (
	"fmt"
	"github.com/2er0/GoGp/iface"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func (g *BaseGA) Init() {

	g.xo.SetUp(g.check, g.gen)
	g.mut.SetUp(g.check, g.gen)
	g.mutUpd.SetLambda(g.lambda)
	g.sel.SetUp(g.mu, g.popComp)

	var pop []iface.Solution
	for i := 0; i < g.mu*3; i++ {
		sol := g.gen.New(g.entitySize)
		sol.SetValues(g.check.CheckAndFix(sol))
		sol.SetScore(g.fit.Calc(sol))
		pop = append(pop, sol)
	}

	sort.Slice(pop, func(i, j int) bool {
		return g.popComp.Comp(pop[i].GetScore(), pop[j].GetScore())
	})

	g.pop = g.sel.Selection(pop[:1], pop[1:])
}

// Run executes the agent
func (g *BaseGA) Run(status chan interface{}, intClose bool) {
	if intClose {
		// if run one only - send close signal
		defer close(status)
	}
	// take time and save it
	defer g.endrun(time.Now(), status)

	status <- fmt.Sprintf("%v Starting execution", g)
	for i := 0; i < g.genCount; i++ {

		if g.sigma < 0.0000001 {
			status <- fmt.Sprintf("%s Generation: %6d | Sigma to low", g, i+1)
			break
		}

		wg := &sync.WaitGroup{}
		wg.Add(g.lambda)

		runInfo := make(chan resValue)

		for j := 0; j < g.lambda; j++ {
			// if both are the same then the children will not succeed over the
			// parents - this case is currently not important
			item1 := rand.Intn(g.mu)
			item2 := rand.Intn(g.mu)

			go func(wg *sync.WaitGroup, item1, item2 int, runInfo chan resValue) {

				defer wg.Done()

				par1 := g.pop[item1]
				par2 := g.pop[item2]

				// create child by crossover
				sol := g.xo.XO(par1, par2)

				// mutate child
				if rand.Float32() < 0.1 {
					index := rand.Intn(sol.GetSize())
					status <- fmt.Sprintf("%s Mutate ", g)
					// mutate one element randomly
					sol.SetValue(index, g.mut.MutOne(sol.GetValue(index), g.sigma))
					sol.SetValues(g.check.CheckAndFix(sol))
				}

				sol.SetScore(g.fit.Calc(sol))

				res := resValue{}

				// check if child is better then parents
				if g.popComp.Comp(par1.GetScore(), sol.GetScore()) ||
					g.popComp.Comp(par2.GetScore(), sol.GetScore()) {
					res.suc = false
				} else {
					res.suc = true
				}

				res.sol = sol

				runInfo <- res

			}(wg, item1, item2, runInfo)
		}

		go monitorWorker(wg, runInfo)

		var childes []iface.Solution
		var offSprings []iface.Solution

		sucCount := 0

		// collect new children and apply offspring selection or not
		for res := range runInfo {
			if res.suc {
				sucCount++
			}
			if !g.offSpring {
				childes = append(childes, res.sol)
			} else {
				// run with offspring selection
				if res.suc {
					offSprings = append(offSprings, res.sol)
				} else {
					childes = append(childes, res.sol)
				}
			}
		}

		for i := 0; i < int(g.mu/3); i++ {
			sol := g.gen.New(g.entitySize)
			sol.SetValues(g.check.CheckAndFix(sol))
			sol.SetScore(g.fit.Calc(sol))
			childes = append(childes, sol)
		}

		var newPop []iface.Solution

		if !g.offSpring {
			newPop = g.sel.Selection(g.pop, childes)
		} else {

			// create new population by taking offspring selection into account
			sort.Slice(offSprings, func(i, j int) bool {
				return g.popComp.Comp(offSprings[i].GetScore(), offSprings[j].GetScore())
			})

			if len(offSprings) > g.mu {
				newPop = offSprings[0:g.mu]
			} else {

				sort.Slice(childes, func(i, j int) bool {
					return g.popComp.Comp(childes[i].GetScore(), childes[j].GetScore())
				})

				less := g.mu - len(offSprings)

				newPop = append(offSprings, childes[0:less]...)
			}
		}

		g.pop = newPop

		g.sigma = g.mutUpd.Update(sucCount, g.sigma)

		status <- fmt.Sprintf("%s Generation: %6d | best: %v | sigma: %v\n",
			g, i+1, g.pop[0], g.sigma)
	}
}
