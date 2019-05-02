package main

import (
	"fmt"
	"math"
	"strings"

	"./neuron"
)

const Shift = 0.3

func answerFuncL01V00(in []float64) float64 {
	if len(in) != 4 {
		return 0
	}
	if !(in[0] == 1 && in[1] == 1) && in[2] == 1 && in[3] == 1 {
		return 1
	} else {
		return 0
	}
}

func answerFuncL01V04(in []float64) float64 {
	if len(in) != 4 {
		return 0
	}
	if (in[0] == 0 || in[2] == 1) && in[1] == 1 || in[1] == 1 && in[3] == 1 {
		return 1
	} else {
		return 0
	}
}

func answerFuncL01V11(in []float64) float64 {
	if len(in) != 4 {
		return 0
	}
	if (in[0] == 1 || in[1] == 1) && in[2] == 1 && in[3] == 1 {
		return 1
	} else {
		return 0
	}
}

func rk1(in []float64) float64 {
	if len(in) != 2 {
		return 0
	}
	if in[0] != in[1] {
		return 1
	} else {
		return 0
	}
}

func answerFuncL03V00(in []float64) float64 {
	if len(in) != 1 {
		return 0
	}
	return 0.5*math.Sin(0.5*in[0]) - 0.5
}

func answerFuncL03V11(in []float64) float64 {
	if len(in) != 1 {
		return 0
	}
	return math.Tan(in[0])
}

func lab01(name string, learningSet neuron.LearningSet, neuron neuron.Neuron, mode byte) {
	fmt.Println("\n--------------------")
	fmt.Println(strings.ToUpper(name))
	fmt.Println("--------------------")

	fmt.Println("\nLearning:")
	neuron.Train(&learningSet, Shift, mode)
	neuron.PrintInfo()

	fmt.Println("\nMinimal learning set:")
	shortestAnswers := neuron.FindMinAnswers(&learningSet, Shift)
	shortestAnswers.PrintInfo()

	fmt.Println("\nTraining on this set:")
	neuron.Train(&shortestAnswers, Shift, mode)
	neuron.PrintInfo()
}

func lab03(name string, learningSet neuron.LearningSet, neuron neuron.Neuron, mode byte) {
	fmt.Println("\n--------------------")
	fmt.Println(strings.ToUpper(name))
	fmt.Println("--------------------")

	fmt.Println("\nLearning:")
	neuron.Train(&learningSet, 1, mode)
	neuron.PrintInfo()

	fmt.Println("\nCorrect out:")
	learningSet.PrintResults()
	fmt.Println("\nResult out:")
	neuron.PrintNet(&learningSet)
	fmt.Println()
}

func lab03Charts(answerFunc neuron.AnswerFunc, width uint, begin float64, end float64, points uint) {
	fmt.Println("Epsilon (Age)")
	learningSet := neuron.CreateFloatLearningSet(answerFunc, width, begin, end, points)
	for maxAge := uint(0); maxAge < 500; maxAge += 10 {
		n := neuron.CreateNeuron(&neuron.ActivationLinear{}, &neuron.ErrorSquare{}, maxAge)
		n.Train(&learningSet, 1, 0)
		fmt.Printf("%v,%v\n", maxAge, n.Epsilon)
	}

	fmt.Println("\nEpsilon (Shift)")
	learningSet = neuron.CreateFloatLearningSet(answerFunc, width, begin, end, points)
	for shift := 0.01; shift <= 1; shift = neuron.Round(shift+.01, neuron.CalcAccuracy) {
		n := neuron.CreateNeuron(&neuron.ActivationLinear{}, &neuron.ErrorSquare{}, 200)
		n.Train(&learningSet, shift, 0)
		fmt.Printf("%v,%v\n", shift, n.Epsilon)
	}

	fmt.Println("\nEpsilon (Width)")
	for width := uint(1); width < 20; width++ {
		learningSet = neuron.CreateFloatLearningSet(answerFunc, width, begin, end, points)
		n := neuron.CreateNeuron(&neuron.ActivationLinear{}, &neuron.ErrorSquare{}, 500)
		n.Train(&learningSet, 0.1, 0)
		fmt.Printf("%v,%v\n", width, n.Epsilon)
	}
}

func lab06(name string, x, t []float64, njm neuron.NJM, mode byte) {
	fmt.Println("\n--------------------")
	fmt.Println(strings.ToUpper(name))
	fmt.Println("--------------------")

	njm.Train(x, t, 1, mode)
}

func lab04(name string, learningSet neuron.LearningSet, neuron neuron.Neuron, mode byte) {
	fmt.Println("\n--------------------")
	fmt.Println(strings.ToUpper(name))
	fmt.Println("--------------------")

	fmt.Println("\nLearning:")
	neuron.Train(&learningSet, Shift, mode)
	neuron.PrintInfo()

	fmt.Println("\nMinimal learning set:")
	shortestAnswers := neuron.FindMinAnswers(&learningSet, Shift)
	shortestAnswers.PrintInfo()

	fmt.Println("\nTraining on this set:")
	neuron.Train(&shortestAnswers, Shift, mode)
	neuron.PrintInfo()
}

func main() {
	// learningSetL01V00 := neuron.CreateBoolLearningSet(answerFuncL01V00, 4, 2, 1)
	// fmt.Println("\nLearning set:")
	// learningSetL01V00.PrintInfo()
	// lab01("step", learningSetL01V00, neuron.CreateNeuron(&neuron.ActivationStep{}, &neuron.ErrorCount{}, 200), 1)
	// lab01("exponential", learningSetL01V00, neuron.CreateNeuron(&neuron.ActivationExp{}, &neuron.ErrorCount{}, 200), 1)

	// learningSetL01V04 := neuron.CreateBoolLearningSet(answerFuncL01V04, 4, 2, 1)
	// fmt.Println("\nLearning set:")
	// learningSetL01V04.PrintInfo()
	// lab01("step", learningSetL01V04, neuron.CreateNeuron(&neuron.ActivationStep{}, &neuron.ErrorCount{}, 200), 1)
	// lab01("module", learningSetL01V04, neuron.CreateNeuron(&neuron.ActivationModule{}, &neuron.ErrorCount{}, 200), 1)

	// learningSetL01V11 := neuron.CreateBoolLearningSet(answerFuncL01V11, 4, 2, 1)
	// fmt.Println("\nLearning set:")
	// learningSetL01V11.PrintInfo()
	// lab01("step", learningSetL01V11, neuron.CreateNeuron(&neuron.ActivationStep{}, &neuron.ErrorCount{}, 200), 1)
	// lab01("exponential", learningSetL01V11, neuron.CreateNeuron(&neuron.ActivationExp{}, &neuron.ErrorCount{}, 200), 1)

	// learningSetRk1 := neuron.CreateBoolLearningSet(rk1, 2, 2, 1)
	// fmt.Println("\nLearning set:")
	// learningSetRk1.PrintInfo()
	// lab01("step", learningSetRk1, neuron.CreateNeuron(&neuron.ActivationStep{}, &neuron.ErrorCount{}, 200), 1)

	// learningSetL03V00 := neuron.CreateFloatLearningSet(answerFuncL03V00, 7, -2, 4, 20)
	// fmt.Println("\nLearning set:")
	// learningSetL03V00.PrintInfo()
	// lab03("step", learningSetL03V00, neuron.CreateNeuron(&neuron.ActivationLinear{}, &neuron.ErrorSquare{}, 100), 1)
	// lab03("step", learningSetL03V00, neuron.CreateNeuron(&neuron.ActivationLinear{}, &neuron.ErrorSquare{}, 300), 1)
	// lab03Charts(answerFuncL03V00, 7, -2, 4, 20)

	// learningSetL03V11 := neuron.CreateFloatLearningSet(answerFuncL03V11, 4, 2, 3, 20)
	// fmt.Println("\nLearning set:")
	// learningSetL03V11.PrintInfo()
	// lab03("step", learningSetL03V11, neuron.CreateNeuron(&neuron.ActivationLinear{}, &neuron.ErrorSquare{}, 2000), 1)
	// lab03Charts(answerFuncL03V11)

	// lab06("exponential", []float64{0.3, -0.1, 0.9}, []float64{0.1, -0.6, 0.2, 0.7},
	// neuron.CreateNJM(3, 3, 4, 0.5, &neuron.ActivationExpExp{}, &neuron.ErrorSquare{}, 1000), 2)

	// lab06("exponential", []float64{-.3}, []float64{-.3, .1, .1},
	// neuron.CreateNJM(1, 1, 3, 0.5, &neuron.ActivationExpExp{}, &neuron.ErrorSquare{}, 1000), 1)

    // lab06("exponential", []float64{-.1}, []float64{-.1, .2, .2},
    // neuron.CreateNJM(1, 1, 3, 0.5, &neuron.ActivationExpExp{}, &neuron.ErrorSquare{}, 1000), 1)

    // learningSetL04V00 := neuron.CreateRBFBoolLearningSet(answerFuncL01V00, 4, [][]float64{
	// 	{0, 0, 1, 1},
	// 	{0, 1, 1, 1},
	// 	{1, 0, 1, 1},
	// })
    // // learningSetL04V00.DisableByArray([]int{0, 2, 3, 4, 5, 6, 8, 9, 12, 13, 15})
	// fmt.Println("\nLearning set:")
	// learningSetL04V00.PrintInfo()
	// lab04("step", learningSetL04V00, neuron.CreateNeuron(&neuron.ActivationStep{}, &neuron.ErrorCount{}, 200), 1)
	// lab04("exponential", learningSetL04V00, neuron.CreateNeuron(&neuron.ActivationExp{}, &neuron.ErrorCount{}, 200), 1)

    // learningSetL04V11 := neuron.CreateRBFBoolLearningSet(answerFuncL01V11, 4, [][]float64{
    //     {0, 1, 1, 1},
    //     {1, 0, 1, 1},
    //     {1, 1, 1, 1},
    // })
    // fmt.Println("\nLearning set:")
    // learningSetL04V11.PrintInfo()
    // lab04("step", learningSetL04V11, neuron.CreateNeuron(&neuron.ActivationStep{}, &neuron.ErrorCount{}, 200), 1)
    // lab04("exponential", learningSetL04V11, neuron.CreateNeuron(&neuron.ActivationExp{}, &neuron.ErrorCount{}, 200), 1)
}
