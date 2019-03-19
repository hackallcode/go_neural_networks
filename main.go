package main

import (
    "./neuron"
    "fmt"
    "math"
    "strings"
)

const Shift = 0.3

func AnswerFunc_1_0(in []float64) float64 {
    if len(in) != 4 {
        return 0
    }
    if !(in[0] == 1 && in[1] == 1) && in[2] == 1 && in[3] == 1 {
        return 1
    } else {
        return 0
    }
}

func AnswerFunc_1_11(in []float64) float64 {
    if len(in) != 4 {
        return 0
    }
    if (in[0] == 1 || in[1] == 1) && in[2] == 1 && in[3] == 1 {
        return 1
    } else {
        return 0
    }
}

func RK_1(in []float64) float64 {
    if len(in) != 2 {
        return 0
    }
    if in[0] != in[1] {
        return 1
    } else {
        return 0
    }
}

func AnswerFunc_3_0(in []float64) float64 {
    if len(in) != 1 {
        return 0
    }
    return 0.5*math.Sin(0.5*in[0]) - 0.5
}

func calc(name string, activationFunc neuron.IActivationFunc, learningSet neuron.LearningSet, maxAge uint, mode byte) {
    fmt.Println("\n--------------------")
    fmt.Println(strings.ToUpper(name))
    fmt.Println("--------------------")

    fmt.Println("\nLearning:")
    myNeuron := neuron.CreateNeuron(activationFunc, maxAge)
    myNeuron.Train(&learningSet, Shift, mode)
    myNeuron.PrintInfo()

    fmt.Println("\nMinimal learning set:")
    shortestAnswers := myNeuron.FindMinAnswers(&learningSet, Shift)
    shortestAnswers.PrintInfo()

    fmt.Println("\nTraining on this set:")
    myNeuron.Train(&shortestAnswers, Shift, mode)
    myNeuron.PrintInfo()
}

func main() {
    //learningSet_1_0 := neuron.CreateBoolLearningSet(AnswerFunc_1_0, 4, 2, 1)
    //fmt.Println("\nLearning set:")
    //learningSet_1_0.PrintInfo()
    //calc("step", &neuron.ActivationStep{}, learningSet_1_0, 200, 1)
    //calc("exponential", &neuron.ActivationExp{}, learningSet_1_0, 200, 1)

    //learningSet_1_11 := neuron.CreateBoolLearningSet(AnswerFunc_1_11, 4, 2, 1)
    //fmt.Println("\nLearning set:")
    //learningSet_1_11.PrintInfo()
    //calc("step", &neuron.ActivationStep{}, learningSet_1_11, 200, 1)
    //calc("exponential", &neuron.ActivationExp{}, learningSet_1_11, 200, 1)

    //rk_1 := neuron.CreateBoolLearningSet(RK_1, 2, 2, 1)
    //calc("step", &neuron.ActivationStep{}, rk_1, 200, 1)

    learningSet_3_0 := neuron.CreateFloatLearningSet(AnswerFunc_3_0, 6, -2, 4, 20)
    fmt.Println("\nLearning set:")
    learningSet_3_0.PrintInfo()
    calc("step", &neuron.ActivationLinear{}, learningSet_3_0, 2000, 1)
}
