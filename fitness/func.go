package fitness

import (
	"fmt"
	"github.com/2er0/GoGp/iface"
	"github.com/Knetic/govaluate"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// A MathFunc Fitness evaluator
type Func struct {
	exp    *govaluate.EvaluableExpression
	count  int
	params []string
}

// NewFunc returns an instance of a new MathFunc fitness function.
// The math function interpretation is form: https://github.com/Knetic/govaluate
func NewFunc(eq string) iface.Fitness {

	re, err := regexp.Compile(`x\d`)
	if err != nil {
		panic("Regex for function not valid")
	}

	if !re.MatchString(eq) {
		panic("Function not valid! Variables must be described as x1, x2, ...")
	}

	matches := re.FindAllStringIndex(eq, -1)
	if matches == nil {
		panic("Function not valid! Variables not found!")
	}

	counter := make(map[string]int)
	for i, _ := range matches {
		s := matches[i][0]
		e := matches[i][1]
		counter[eq[s:e]]++
	}

	exp, err := govaluate.NewEvaluableExpression(eq)
	if err != nil {
		panic("Expression not valid")
	}

	return &Func{exp: exp, count: len(counter), params: exp.Vars()}
}

// Calc returns the fitness of the given solution by a math function
func (f *Func) Calc(sol iface.Solution) interface{} {
	if f.count != sol.GetSize() {
		panic("Sizes does not match")
	}

	parameters := make(map[string]interface{}, f.count)
	val := reflect.ValueOf(sol.GetValues())
	for i := 0; i < val.Len(); i++ {
		parameters[strings.Join([]string{"x", strconv.Itoa(i + 1)}, "")] = val.Index(i).Interface()
	}

	res, err := f.exp.Evaluate(parameters)
	if err != nil {
		panic(fmt.Sprintf("Function evaluation was not successful\n%v", err))
	}

	return res
}
