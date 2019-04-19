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
	Epsilon        float64
}

func (n *Neuron) PrintInfo() {
	fmt.Printf("age: %v, Epsilon: %v, Weights: [%v]\n", n.age, n.Epsilon, WeightsToString(n.weights, ", "))
}

func (n *Neuron) Train(learningSet *LearningSet, shift float64, mode byte) (isLearned bool) {
	// Init weights
	n.weights = make([]float64, learningSet.varsCount)

	for n.age = 0; n.age < n.maxAge; {
		if mode == 1 {
			fmt.Printf("%v; [%v]; ", n.age, WeightsToString(n.weights, ", "))
		}

		// Count age
		results := make([]float64, len(learningSet.data))
		errorSum := float64(0)
		for r, answer := range learningSet.data {
			if learningSet.isTest[r] {
				continue
			}

			net := float64(0)
			for i := uint(0); i < learningSet.varsCount; i++ {
				net += n.weights[i] * float64(answer.input[i])
			}

			results[r] = n.activationFunc.Result(Round(net, CalcAccuracy))
			err := float64(answer.answer) - float64(results[r])
			errorSum += n.errorFunc.Step(results[r], answer.answer)

			if mode == 2 {
				fmt.Printf("age: %v.%v; w: %v; net: %v; y: %v; Epsilon: %v;\n",
					n.age, r, WeightsToString(n.weights, ", "), Round(net, CalcAccuracy), results[r], Round(err, CalcAccuracy))
			}

			if Round(err, CalcAccuracy) != 0 {
				delta := err * shift * n.activationFunc.Derivative(net)
				for i := uint(0); i < learningSet.varsCount; i++ {
					n.weights[i] += delta * float64(answer.input[i])
				}
			}
		}
		errorSum = Round(n.errorFunc.Result(errorSum), CalcAccuracy)

		if mode == 1 {
			fmt.Printf("[%v]; %v\n", ResultsToString(results, learningSet.isTest, ", "), errorSum)
		}

		n.age++
		if errorSum == 0 {
			isLearned = true
			break
		}
	}

	// Count errors on non-teaching sets
	n.Epsilon = float64(0)
	for r, answer := range learningSet.data {
		if !learningSet.isTest[r] {
			continue
		}

		net := float64(0)
		for i := uint(0); i < learningSet.varsCount; i++ {
			net += n.weights[i] * float64(answer.input[i])
		}

		result := n.activationFunc.Result(Round(net, CalcAccuracy))
		n.Epsilon += n.errorFunc.Step(result, answer.answer)
	}
	n.Epsilon = Round(n.errorFunc.Result(n.Epsilon), CalcAccuracy)

	// Round total values
	for i := uint(0); i < learningSet.varsCount; i++ {
		n.weights[i] = Round(n.weights[i], ResultAccuracy)
	}

	if n.Epsilon != 0 {
		return false
	}
	return isLearned
}

func (n *Neuron) PrintNet(learningSet *LearningSet) {
	for r, answer := range learningSet.data {
		net := float64(0)
		for i := uint(0); i < learningSet.varsCount; i++ {
			net += n.weights[i] * float64(answer.input[i])
		}
		fmt.Print(n.activationFunc.Result(Round(net, ResultAccuracy)))
		if r < len(learningSet.data)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
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
