package agent

import (
	"github.com/2er0/GoGp/iface"
	"sync"
)

// A resValue represents the temporal communication
// struct for running in parallel
type resValue struct {
	sol iface.Solution
	suc bool
}

// monitorWorker waits for all goroutines to finish
// and close the communication channel after that
// basic synchronisation
func monitorWorker(wg *sync.WaitGroup, runInfo chan resValue) {
	wg.Wait()
	close(runInfo)
}
