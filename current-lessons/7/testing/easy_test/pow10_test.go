package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	Num      float64
	Expected float64
}

func TestPow10(t *testing.T) {
	asrt := assert.New(t)
	for i, c := range createPow10Cases() {
		res := Pow10(c.Num)
		asrt.Equal(c.Expected, res, "[%d] Expected: %v; Actual: %v", i, c.Expected, res)
	}
}

func createPow10Cases() []Case {
	return []Case{
		{
			Num:      0,
			Expected: 0,
		},
		{
			Num:      1,
			Expected: 1,
		},
		{
			Num:      2,
			Expected: 1024,
		},
		{
			Num:      3,
			Expected: 59049,
		},
	}
}
