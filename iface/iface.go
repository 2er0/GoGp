package iface

// Definition to run a agent which contains everything to
// optimize the given problem
type Task interface {
	Init()
	Run(status chan interface{}, intClose bool)
	Result() string
	String() string
}

// Description for a solution candidate
type Solution interface {
	GetValues() interface{}
	SetValues(interface{})
	GetValue(i int) interface{}
	SetValue(i int, v interface{})
	GetScore() interface{}
	SetScore(interface{})
	GetSize() int
	Copy() Solution
	String() string
}

// Compare interface to sort solution candidates
type Comp interface {
	Less() bool
	Comp(i, j interface{}) bool
}

// Creating a new solution for the given problem set
type Generator interface {
	New(size int) Solution
	NewClean(size int) Solution
}

// Calculating the fitness for a given solution candidate
type Fitness interface {
	Calc(sol Solution) interface{}
}

// Mutating by a solution or by singe value
type Mutator interface {
	Mut(sol Solution, sigma float32) Solution
	MutOne(i interface{}, sigma float32) interface{}
	SetUp(check SolutionCheck, gen Generator)
	Check() bool
}

// Mutation distance changer for adapting the sigma in the mutation
type MutChanger interface {
	SetLambda(lam int)
	Update(suc int, sig float32) float32
}

// crossover for given parent solution and producing one child
type Crossover interface {
	XO(s1, s2 Solution) Solution
	SetUp(check SolutionCheck, gen Generator)
	Check() bool
}

// domain specific solution fixer, for correcting
// crossover or mutation error
type SolutionCheck interface {
	CheckAndFix(sol Solution) interface{}
}

// Selection for the best candidates
type Selection interface {
	SetUp(mu int, co Comp)
	Selection(parent, child []Solution) []Solution
}
