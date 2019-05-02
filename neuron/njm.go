package neuron

import (
	"fmt"
)

type NJM struct {
	activationFunc IActivationFunc
	errorFunc      IErrorFunc
	weights        []float64
	hidden         [][]float64
	output         [][]float64
	age            uint
	maxAge         uint
	Epsilon        float64
	n, j, m        uint
}

func CreateNJM(n, j, m uint, w float64, activationFunc IActivationFunc, errorFunc IErrorFunc, maxAge uint) (result NJM) {
	result.activationFunc = activationFunc
	result.errorFunc = errorFunc
	result.maxAge = maxAge
	result.hidden = make([][]float64, j)
	for i := range result.hidden {
		result.hidden[i] = make([]float64, n+1)
		for nn := uint(0); nn <= n; nn++ {
			result.hidden[i][nn] = w
		}
	}
	result.output = make([][]float64, m)
	for i := range result.output {
		result.output[i] = make([]float64, j+1)
		for jj := uint(0); jj <= j; jj++ {
			result.output[i][jj] = w
		}
	}
	result.n = n
	result.j = j
	result.m = m
	return
}

func (l *NJM) Train(x, t []float64, shift float64, mode byte) (isLearned bool) {
	for l.age = 0; l.age < l.maxAge; l.age++ {
		if mode >= 2 {
			fmt.Printf("age: %v; ", l.age + 1)
			for j := uint(0); j < l.j; j++ {
				fmt.Printf("h: [%v]; ", WeightsToString(l.hidden[j], ", "))
			}
			for m := uint(0); m < l.m; m++ {
				fmt.Printf("o: [%v]; ", WeightsToString(l.output[m], ", "))
			}
		}

		hiddenNets := make([]float64, l.j)
		hiddenResults := make([]float64, l.j)
		hiddenErrors := make([]float64, l.j)
		for j := uint(0); j < l.j; j++ {
			hiddenNets[j] = float64(l.hidden[j][0])
			for n := uint(0); n < l.n; n++ {
				hiddenNets[j] += l.hidden[j][n+1] * float64(x[n])
			}
			hiddenResults[j] = l.activationFunc.Result(Round(hiddenNets[j], CalcAccuracy))
		}

		nets := make([]float64, l.m)
		results := make([]float64, l.m)
		errors := make([]float64, l.m)
		for m := uint(0); m < l.m; m++ {
			nets[m] = float64(l.output[m][0])
			for j := uint(0); j < l.j; j++ {
				nets[m] += l.output[m][j+1] * float64(hiddenResults[j])
			}
			results[m] = l.activationFunc.Result(Round(nets[m], CalcAccuracy))
		}

		errorSum := float64(0)
		for m := uint(0); m < l.m; m++ {
			errors[m] = (float64(t[m]) - float64(results[m])) * l.activationFunc.Derivative(nets[m])
			errorSum += l.errorFunc.Step(results[m], t[m])
		}
		errorSum = Round(l.errorFunc.Result(errorSum), CalcAccuracy)

		for j := uint(0); j < l.j; j++ {
			sum := float64(0)
			for m := uint(0); m < l.m; m++ {
				sum += l.output[m][j+1] * errors[m]
			}
			hiddenErrors[j] = sum * l.activationFunc.Derivative(hiddenNets[j])
		}

		for j := uint(0); j < l.j; j++ {
			l.hidden[j][0] += shift * hiddenErrors[j]
			for n := uint(0); n < l.n; n++ {
				l.hidden[j][n+1] += shift * x[n] * hiddenErrors[j]
			}
		}

		for m := uint(0); m < l.m; m++ {
			l.output[m][0] += shift * errors[m]
			for j := uint(0); j < l.j; j++ {
				l.output[m][j+1] += shift * hiddenResults[j] * errors[m]
			}
		}

		if mode >= 1 {
			fmt.Printf("y: [%v]; Epsilon: %v;\n", AllResultsToString(results, ", "), errorSum)
		}

		if Round(errorSum, ResultAccuracy) == 0 {
			if mode >= 1 {
				for j := uint(0); j < l.j; j++ {
					fmt.Printf("h: [%v]; ", WeightsToString(l.hidden[j], ", "))
				}
				for m := uint(0); m < l.m; m++ {
					fmt.Printf("o: [%v]; ", WeightsToString(l.output[m], ", "))
				}
			}
			return true
		}
	}
	if mode >= 1 {
		for j := uint(0); j < l.j; j++ {
			fmt.Printf("h: [%v]; ", WeightsToString(l.hidden[j], ", "))
		}
		for m := uint(0); m < l.m; m++ {
			fmt.Printf("o: [%v]; ", WeightsToString(l.output[m], ", "))
		}
	}
	return false
}
