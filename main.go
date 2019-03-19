package main

import (
    "./neuron"
    "fmt"
    "math"
    "strings"
)

const Shift = 0.3

func RK(in []uint8) uint8 {
    if len(in) != 2 {
        return 0
    }
    if in[0] != in[1] {
        return 1
    } else {
        return 0
    }
}

func Answer0(in []uint8) uint8 {
    if len(in) != 4 {
        return 0
    }
    if !(in[0] == 1 && in[1] == 1) && in[2] == 1 && in[3] == 1 {
        return 1
    } else {
        return 0
    }
}

func Answer11(in []uint8) uint8 {
    if len(in) != 4 {
        return 0
    }
    if (in[0] == 1 || in[1] == 1) && in[2] == 1 && in[3] == 1 {
        return 1
    } else {
        return 0
    }
}

type ActivationStep struct{}

func (a *ActivationStep) Result(net float64) uint8 {
    if net < 0 {
        return 0
    } else {
        return 1
    }
}

func (a *ActivationStep) Derivative(net float64) float64 {
    return 1
}

type ActivationExp struct{}

func (a *ActivationExp) Count(net float64) float64 {
    return 1 / (1 + math.Exp(-net))
}

func (a *ActivationExp) Result(net float64) uint8 {
    if a.Count(net) < 0.5 {
        return 0
    } else {
        return 1
    }
}

func (a *ActivationExp) Derivative(net float64) float64 {
    return a.Count(net) * (1 - a.Count(net))
}

func calc(name string, activationFunc neuron.IActivationFunc, answers neuron.Answers, mode byte) {
    fmt.Println("\n--------------------")
    fmt.Println(strings.ToUpper(name))
    fmt.Println("--------------------")

    fmt.Println("\nLearning:")
    myNeuron := neuron.CreateNeuron(activationFunc)
    myNeuron.Train(&answers, Shift, mode)
    myNeuron.PrintInfo()

    fmt.Println("\nMinimal learning set:")
    shortestAnswers := myNeuron.FindMinAnswers(&answers, Shift)
    shortestAnswers.PrintInfo()

    fmt.Println("\nTraining on this set:")
    myNeuron.Train(&shortestAnswers, Shift, mode)
    myNeuron.PrintInfo()
}

func main() {
    answers0 := neuron.CreateAnswers(Answer0, 4, 2, 1)
    fmt.Println("\nLearning set:")
    answers0.PrintInfo()
    //calc("step", &ActivationStep{}, answers0, 2)
    calc("exponential", &ActivationExp{}, answers0, 1)

    //answers11 := neuron.CreateAnswers(Answer11, 4, 2, 1)
    //fmt.Println("\nLearning set:")
    //answers11.PrintInfo()
    //calc("step", &ActivationStep{}, answers11, 1)
    //calc("exponential", &ActivationExp{}, answers11, 1)

    //rk := neuron.CreateAnswers(RK, 2, 2, 1)
    //calc("step", &ActivationStep{}, rk, 1)
}
