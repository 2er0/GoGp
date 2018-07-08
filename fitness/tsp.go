package fitness

import (
	"fmt"
	"github.com/2er0/GoGp/iface"
	"github.com/2er0/GoGp/utils"
	"math"
)

// A Tsp contains a distance matrix for a TSP problem
type Tsp struct {
	dist   [][]float32
	origin map[int]utils.Point
}

// NewTsp returns an instance of a new Tsp fitness evaluator
// it transform a list of points into a distance matrix
// on creation time to save time during evaluating of
// solution candidates
func NewTsp(dists map[int]utils.Point) iface.Fitness {
	startAt := 0
	if _, ok := dists[0]; !ok {
		startAt++
	}
	size := len(dists)
	dist := make([][]float32, size)
	for i := 0; i < size; i++ {
		dist[i] = make([]float32, size)
	}

	// build distance matrix
	for i := 0; i < size; i++ {
		dist[i][i] = 0
		a := dists[i + startAt]

		for j := i + 1; j < size; j++ {
			b := dists[j + startAt]
			dis := float32(math.Sqrt(math.Pow(b.X - a.X, 2) +
				math.Pow(b.Y - a.Y, 2)))

			dist[i][j] = dis
			dist[j][i] = dis
		}
	}

	return &Tsp{dist: dist, origin: dists}
}

// Calc returns the traveling distance of the given solution
func (t *Tsp) Calc(sol iface.Solution) interface{} {
	size := sol.GetSize()
	values := sol.GetValues().([]int)
	prev := values[0]
	var sum float32 = 0
	for i := 1; i < size; i++ {
		cur := values[i]
		if prev < 0 || prev >= size || cur < 0 || cur >= size {
			fmt.Println(prev, cur)
		}
		sum = sum + t.dist[prev][cur]
		prev = cur
	}
	return sum
}
