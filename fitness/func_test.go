package fitness

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"regexp"
	"testing"
)

func TestFoo(t *testing.T) {
	boothFunc := "(x1 + 2 * x2 - 7) ** 2 + (2 * x1 + x2 - 5) ** 2"
	exp, err := govaluate.NewEvaluableExpression(boothFunc)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if len(exp.Vars()) != 4 {
		t.Fail()
	}

	parameters := make(map[string]interface{}, 2)
	parameters["x1"] = 1
	parameters["x2"] = 2

	_, err = exp.Evaluate(parameters)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestFoo2(t *testing.T) {
	eq := "(x1 + 2 * x2 - 7) ** 2 + (2 * x1 + x2 - 5) ** 2"
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

	if len(counter) != 2 {
		t.Fail()
	}
}
