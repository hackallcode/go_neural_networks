package neuron

import (
    "fmt"
)

type Neuron struct {
    activationFunc IActivationFunc
    weights        []float64
    age            uint
    maxAge         uint
}

func (n *Neuron) SetActivationFunc(activationFunc IActivationFunc) {
    n.activationFunc = activationFunc
}

func (n *Neuron) PrintInfo() {
    fmt.Printf("Age: %v, Weights: [%v]\n", n.age, WeightsToString(n.weights, ", "))
}

func (n *Neuron) Train(learningSet *LearningSet, shift float64, mode byte) bool {
    // Init weights
    n.weights = make([]float64, learningSet.varsCount)
    n.age = 0

    for age := uint(0); ; age++ {
        if mode == 1 {
            fmt.Printf("%v; [%v]; ", n.age, WeightsToString(n.weights, ", "))
        }
        results := make([]float64, len(learningSet.data))
        errorSum := uint(0)

        // Result age
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
                errorSum++

                delta := err * shift * n.activationFunc.Derivative(net)
                for i := uint(0); i < learningSet.varsCount; i++ {
                    n.weights[i] += delta * float64(answer.input[i])
                }
            }
        }
        n.age++

        if mode == 1 {
            fmt.Printf("[%v]; %v;\n", ResultsToString(results, learningSet.skipped, ", "), errorSum)
        }
        if errorSum == 0 {
            break
        }
        if age == n.maxAge {
            return false
        }
    }

    // Count errors on non-teaching sets
    errorCount := 0
    for r, answer := range learningSet.data {
        if !learningSet.skipped[r] {
            continue
        }

        net := float64(0)
        for i := uint(0); i < learningSet.varsCount; i++ {
            net += n.weights[i] * float64(answer.input[i])
        }
        result := n.activationFunc.Result(net)

        err := int16(answer.answer) - int16(result)
        if err != 0 {
            errorCount++
        }
    }
    if errorCount > 0 {
        return false
    }

    // Round total values
    for i := uint(0); i < learningSet.varsCount; i++ {
        n.weights[i] = Round(n.weights[i], ResultAccuracy)
    }
    return true
}

func CreateNeuron(activationFunc IActivationFunc, maxAge uint) (result Neuron) {
    result.activationFunc = activationFunc
    result.maxAge = maxAge
    return
}
