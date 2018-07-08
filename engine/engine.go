package engine

import (
	"github.com/2er0/GoGp/iface"
	"log"
	"os"
	"sync"
)

// An Engine represents the basic management system
// which takes care of handling and running tasks with agents
// and there specific problem
type Engine struct {
	todos []iface.Task
	wg    sync.WaitGroup
	dones []iface.Task
	log   *log.Logger
}

// New returns a new Engine with the basic functionality
func New() *Engine {
	return &Engine{todos: []iface.Task{}, dones: []iface.Task{}, log: log.New(os.Stdout, "", log.Ldate|log.Ltime)}
}

// Execute first task in queue
func (e *Engine) Run() {
	// check if there is something to execute
	if len(e.todos) < 1 {
		return
	}
	// get first element and execute in gorun
	toRun := e.todos[0]
	// reuse allocated size of slice
	copy(e.todos, e.todos[1:])
	e.todos = e.todos[:len(e.todos)-1]

	e.wg.Add(1)

	status := make(chan interface{})

	// execute an goroutine
	go toRun.Run(status, true)

	// wait for goroutine to finish
	for s := range status {
		e.log.Print(s)
	}

	e.wg.Done()

	// add finished task to done
	e.dones = append(e.dones, toRun)
	e.log.Print(toRun.Result())
}

// Run every not executed task
func (e *Engine) RunAll() {
	// check if there is something to execute
	if len(e.todos) < 1 {
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(e.todos))
	status := make(chan interface{})

	// run every given task in the todos queue in parallel
	for i := 0; i < len(e.todos); i++ {
		go func(wg *sync.WaitGroup, item iface.Task, st chan interface{}) {
			defer wg.Done()
			item.Run(st, false)
			st <- item.Result()
		}(wg, e.todos[i], status)
	}

	// wait for all task to complete
	go monitorWorker(wg, status)

	// log every thing from the running task to the
	// engines log
	for s := range status {
		e.log.Print(s)
	}

	e.dones = append(e.dones, e.todos...)
	e.todos = e.todos[:0]
}

// Add and Init new task in the engine
func (e *Engine) Add(task iface.Task) {
	task.Init()
	e.todos = append(e.todos, task)
}

func monitorWorker(wg *sync.WaitGroup, status chan interface{}) {
	wg.Wait()
	close(status)
}

func (e *Engine) GetAll() {
	for i := range e.dones {
		e.log.Println(e.dones[i].Result())
	}
}