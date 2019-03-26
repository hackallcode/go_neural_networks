package neuron

import (
	"math"
)

type IActivationFunc interface {
	Result(float64) float64
	Derivative(float64) float64
}

type ActivationStep struct{}

func (a *ActivationStep) Result(net float64) float64 {
	if net < 0 {
		return 0
	} else {
		return 1
	}
}

func (a *ActivationStep) Derivative(net float64) float64 {
	return 1
}

type ActivationModule struct{}

func (a *ActivationModule) Result(net float64) float64 {
	if 0.5*(net/(1+math.Abs(net))+1) < 0.5 {
		return 0
	} else {
		return 1
	}
}

func (a *ActivationModule) Derivative(net float64) float64 {
	return 0.5 * math.Pow(math.Abs(net)+1, 2)
}

type ActivationExp struct{}

func (a *ActivationExp) Count(net float64) float64 {
	return 1 / (1 + math.Exp(-net))
}

func (a *ActivationExp) Result(net float64) float64 {
	if a.Count(net) < 0.5 {
		return 0
	} else {
		return 1
	}
}

func (a *ActivationExp) Derivative(net float64) float64 {
	return a.Count(net) * (1 - a.Count(net))
}

type ActivationLinear struct{}

func (a *ActivationLinear) Result(net float64) float64 {
	return net
}

func (a *ActivationLinear) Derivative(net float64) float64 {
	return 1
}
