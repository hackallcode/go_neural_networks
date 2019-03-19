package neuron

import (
	"fmt"
)

type Neuron struct {
	activationFunc IActivationFunc
	errorFunc      IErrorFunc
	weights        []float64
	age            uint
	maxAge         uint
}

func (n *Neuron) PrintInfo() {
	fmt.Printf("Age: %v, Weights: [%v]\n", n.age, WeightsToString(n.weights, ", "))
}

func (n *Neuron) Train(learningSet *LearningSet, shift float64, mode byte) (isLearned bool) {
	// Init weights
	n.weights = make([]float64, learningSet.varsCount)
	n.age = 0

	for age := uint(0); ; age++ {
		if mode == 1 {
			fmt.Printf("%v; [%v]; ", n.age, WeightsToString(n.weights, ", "))
		}

		// Count age
		results := make([]float64, len(learningSet.data))
		for r, answer := range learningSet.data {
			if learningSet.skipped[r] {
				continue
			}

			net := float64(0)
			for i := uint(0); i < learningSet.varsCount; i++ {
				net += n.weights[i] * float64(answer.input[i])
			}

			results[r] = n.activationFunc.Result(Round(net, CalcAccuracy))
			err := float64(answer.answer) - float64(results[r])

			if mode == 2 {
				fmt.Printf("age: %v.%v; w: %v; net: %v; y: %v; err: %v;\n",
					n.age, r, WeightsToString(n.weights, ", "), Round(net, CalcAccuracy), results[r], err)
			}

			if Round(err, CalcAccuracy) != 0 {
				delta := err * shift * n.activationFunc.Derivative(net)
				for i := uint(0); i < learningSet.varsCount; i++ {
					n.weights[i] += delta * float64(answer.input[i])
				}
			}
		}
		n.age++

		// Count error
		errorSum := float64(0)
		for r, answer := range learningSet.data {
			if learningSet.skipped[r] {
				continue
			}

			net := float64(0)
			for i := uint(0); i < learningSet.varsCount; i++ {
				net += n.weights[i] * float64(answer.input[i])
			}

			result := n.activationFunc.Result(Round(net, CalcAccuracy))
			errorSum += n.errorFunc.Step(result, answer.answer)
		}
		errorSum = n.errorFunc.Result(errorSum)

		if mode == 1 {
			fmt.Printf("[%v]; %v\n", ResultsToString(results, learningSet.skipped, ", "), errorSum)
		}

		if Round(errorSum, CalcAccuracy) == 0 {
			isLearned = true
			break
		}
		if age == n.maxAge {
			isLearned = false
			break
		}
	}

	// Count errors on non-teaching sets
	errorSum := float64(0)
	for r, answer := range learningSet.data {
		if !learningSet.skipped[r] {
			continue
		}

		net := float64(0)
		for i := uint(0); i < learningSet.varsCount; i++ {
			net += n.weights[i] * float64(answer.input[i])
		}

		result := n.activationFunc.Result(Round(net, CalcAccuracy))
		errorSum += n.errorFunc.Step(result, answer.answer)
	}
	errorSum = n.errorFunc.Result(errorSum)

	// Round total values
	for i := uint(0); i < learningSet.varsCount; i++ {
		n.weights[i] = Round(n.weights[i], ResultAccuracy)
	}

	if mode == 1 {
		fmt.Printf("Checker error: %v; ", errorSum)
	}

	if Round(errorSum, CalcAccuracy) != 0 {
		return false
	}
	return isLearned
}

func CreateNeuron(activationFunc IActivationFunc, errorFunc IErrorFunc, maxAge uint) (result Neuron) {
	result.activationFunc = activationFunc
	result.errorFunc = errorFunc
	result.maxAge = maxAge
	return
}

func CopyNeuron(neuron Neuron) (result Neuron) {
	result.activationFunc = neuron.activationFunc
	result.errorFunc = neuron.errorFunc
	result.maxAge = neuron.maxAge
	return
}
