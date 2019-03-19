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

func calc(name string, learningSet neuron.LearningSet, neuron neuron.Neuron, mode byte) {
	fmt.Println("\n--------------------")
	fmt.Println(strings.ToUpper(name))
	fmt.Println("--------------------")

	fmt.Println("\nLearning:")
	neuron.Train(&learningSet, Shift, mode)
	neuron.PrintInfo()

	//fmt.Println("\nMinimal learning set:")
	//shortestAnswers := neuron.FindMinAnswers(&learningSet, Shift)
	//shortestAnswers.PrintInfo()
	//
	//fmt.Println("\nTraining on this set:")
	//neuron.Train(&shortestAnswers, Shift, mode)
	//neuron.PrintInfo()
}

func main() {
	//learningSet_1_0 := neuron.CreateBoolLearningSet(AnswerFunc_1_0, 4, 2, 1)
	//fmt.Println("\nLearning set:")
	//learningSet_1_0.PrintInfo()
	//calc("step", learningSet_1_0, neuron.CreateNeuron(&neuron.ActivationStep{}, &neuron.ErrorCount{}, 200), 1)
	//calc("exponential", learningSet_1_0, neuron.CreateNeuron(&neuron.ActivationExp{}, &neuron.ErrorCount{}, 200), 1)

	//learningSet_1_11 := neuron.CreateBoolLearningSet(AnswerFunc_1_11, 4, 2, 1)
	//fmt.Println("\nLearning set:")
	//learningSet_1_11.PrintInfo()
	//calc("step", learningSet_1_11, neuron.CreateNeuron(&neuron.ActivationStep{}, &neuron.ErrorCount{}, 200), 1)
	//calc("exponential", learningSet_1_11, neuron.CreateNeuron(&neuron.ActivationExp{}, &neuron.ErrorCount{}, 200), 1)

	//rk_1 := neuron.CreateBoolLearningSet(RK_1, 2, 2, 1)
	//calc("step", rk_1, neuron.CreateNeuron(&neuron.ActivationStep{}, &neuron.ErrorCount{}, 200), 1)

	learningSet_3_0 := neuron.CreateFloatLearningSet(AnswerFunc_3_0, 6, -2, 4, 20)
	fmt.Println("\nLearning set:")
	learningSet_3_0.PrintInfo()
	calc("step", learningSet_3_0, neuron.CreateNeuron(&neuron.ActivationLinear{}, &neuron.ErrorSquare{}, 4000), 1)
}
