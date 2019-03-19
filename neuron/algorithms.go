package neuron

import (
	"math"
	"strconv"
	"sync"
)

/////////////////////
//     General     //
/////////////////////

const CalcAccuracy = 1000000000
const ResultAccuracy = 1000

// 10 - 1 digit after point
// 100 - 2 digits after point
// ...
func Round(x float64, accuracy float64) float64 {
	return math.Round(x*accuracy) / accuracy
}

func WeightsToString(weights []float64, separator string) (result string) {
	for _, v := range weights {
		result += strconv.FormatFloat(Round(v, ResultAccuracy), 'f', -1, 64) + separator
	}
	if len(result) >= len(separator) {
		result = result[:len(result)-len(separator)]
	}
	return
}

func ResultsToString(results []float64, skipped []bool, separator string) (result string) {
	for i, v := range results {
		if !skipped[i] {
			result += strconv.FormatFloat(Round(v, ResultAccuracy), 'f', -1, 64) + separator
		}
	}
	if len(result) >= len(separator) {
		result = result[:len(result)-len(separator)]
	}
	return
}

////////////////////////////
// Find min answer matrix //
////////////////////////////

type MinAnswerParams struct {
	// Input params
	inSkipped []bool
	inNeuron  Neuron
	inShift   float64

	// Output params
	outCount       *uint
	outNeuron      *Neuron
	outLearningSet *LearningSet

	// Sync params
	mutex  *sync.Mutex
	waiter *sync.WaitGroup
}

func FindMinAnswersIteration(learningSet *LearningSet, begin int, count uint, params MinAnswerParams) {
	defer params.waiter.Done()

	for i := begin; i < len(learningSet.skipped); i++ {
		// If must be skipped OR already exist
		if params.inSkipped[i] || !learningSet.skipped[i] {
			continue
		}

		params.waiter.Add(2)
		go func(i int) {
			defer params.waiter.Done()

			// Init neuron
			neuron := CopyNeuron(params.inNeuron)

			// ... and answers set
			currentAnswers := *learningSet
			currentAnswers.skipped = make([]bool, len(learningSet.skipped))
			copy(currentAnswers.skipped, learningSet.skipped)
			currentAnswers.skipped[i] = false

			// If correct answers write result
			if neuron.Train(&currentAnswers, params.inShift, 0) {
				params.mutex.Lock()
				// If more compact OR same compact answers, but younger neuron
				if (count < *params.outCount) || (count == *params.outCount && neuron.age < params.outNeuron.age) {
					*params.outCount = count
					*params.outNeuron = neuron
					*params.outLearningSet = currentAnswers
				}
				params.mutex.Unlock()
			}

			FindMinAnswersIteration(&currentAnswers, i, count+1, params)
		}(i)
	}
}

func (n *Neuron) FindMinAnswers(learningSet *LearningSet, shift float64) LearningSet {
	beginAnswers := *learningSet
	beginAnswers.skipped = make([]bool, len(learningSet.skipped))
	for i := range learningSet.skipped {
		beginAnswers.skipped[i] = true
	}

	outCount := uint(len(learningSet.skipped))
	resultLearningSet := *learningSet
	params := MinAnswerParams{
		inSkipped: learningSet.skipped,
		inNeuron:  *n,
		inShift:   shift,

		outCount:       &outCount,
		outNeuron:      n,
		outLearningSet: &resultLearningSet,

		mutex:  &sync.Mutex{},
		waiter: &sync.WaitGroup{},
	}

	params.waiter.Add(1)
	FindMinAnswersIteration(&beginAnswers, 0, 0, params)
	params.waiter.Wait()
	return *params.outLearningSet
}
