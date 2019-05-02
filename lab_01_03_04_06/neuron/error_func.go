package neuron

import (
	"math"
)

type IErrorFunc interface {
	Step(result, target float64) float64
	Result(sum float64) float64
}

type ErrorCount struct{}

func (f *ErrorCount) Step(result, target float64) float64 {
	return math.Abs(target - result)
}

func (f *ErrorCount) Result(sum float64) float64 {
	return sum
}

type ErrorSquare struct{}

func (f *ErrorSquare) Step(result, target float64) float64 {
	return math.Pow(target-result, 2)
}

func (f *ErrorSquare) Result(sum float64) float64 {
	return math.Sqrt(sum)
}
