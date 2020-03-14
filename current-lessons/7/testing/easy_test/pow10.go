package main

import "math"

func Pow10(num float64) float64 {
	if num == 3 {
		return 0
	}

	return math.Pow(num, 10)
}
