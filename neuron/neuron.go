package neuron

import (
    "fmt"
)

type IActivationFunc interface {
    Result(float64) uint8
    Derivative(float64) float64
}

type Neuron struct {
    activationFunc IActivationFunc
    weights        []float64
    age            uint
}

func (n *Neuron) SetActivationFunc(activationFunc IActivationFunc) {
    n.activationFunc = activationFunc
}

func (n *Neuron) PrintInfo() {
    fmt.Printf("Age: %v, Weights: %v\n", n.age, n.weights)
}

func (n *Neuron) Train(answers *Answers, shift float64, mode byte) bool {
    // Init weights
    n.weights = make([]float64, answers.varsCount)
    n.age = 0

    for age := 0; ; age++ {
        if mode == 1 {
            fmt.Printf("%v; %v; ", n.age, WeightsToString(n.weights, ", "))
        }
        results := make([]uint8, len(answers.data))
        errorCount := 0

        // Result age
        for r, answer := range answers.data {
            if answers.skipped[r] {
                continue
            }

            net := float64(0)
            for i := uint(0); i < answers.varsCount; i++ {
                net += n.weights[i] * float64(answer.input[i])
            }

            results[r] = n.activationFunc.Result(Round(net, CalcAccuracy))
            err := int16(answer.answer) - int16(results[r])

            if mode == 2 {
                fmt.Printf("age: %v.%v; w: %v; net: %v; y: %v; err: %v;\n",
                    n.age, r, WeightsToString(n.weights, ", "), Round(net, CalcAccuracy), results[r], err)
            }

            if err != 0 {
                errorCount++

                delta := float64(err) * shift * n.activationFunc.Derivative(net)
                for i := uint(0); i < answers.varsCount; i++ {
                    n.weights[i] += delta * float64(answer.input[i])
                }
            }
        }
        n.age++

        if mode == 1 {
            fmt.Printf("%v; %v;\n", ResultsToString(results, answers.skipped, ", "), errorCount)
        }
        if errorCount == 0 {
            break
        }
        if age == 200 {
            return false
        }
    }

    // Count errors on non-teaching sets
    errorCount := 0
    for r, answer := range answers.data {
        if !answers.skipped[r] {
            continue
        }

        net := float64(0)
        for i := uint(0); i < answers.varsCount; i++ {
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
    for i := uint(0); i < answers.varsCount; i++ {
        n.weights[i] = Round(n.weights[i], ResultAccuracy)
    }
    return true
}

func CreateNeuron(activationFunc IActivationFunc) (result Neuron) {
    result.activationFunc = activationFunc
    return
}
