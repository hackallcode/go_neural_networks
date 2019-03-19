package neuron

import (
    "context"
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
    for _, v := range weights{
        result += strconv.FormatFloat(Round(v, ResultAccuracy), 'f', -1, 64) + separator
    }
    result = result[:len(result) - len(separator)]
    return
}

func ResultsToString(results []uint8, skipped []bool, separator string) (result string) {
    for i, v := range results {
        if !skipped[i] {
            result += strconv.Itoa(int(v)) + separator
        }
    }
    result = result[:len(result) - len(separator)]
    return
}

////////////////////////////
// Find min answer matrix //
////////////////////////////

type MinAnswerParams struct {
    // Input params
    inSkipped        []bool
    inActivationFunc IActivationFunc
    inShift          float64

    // Output params
    wasFound   *bool
    outNeuron  *Neuron
    outAnswers *Answers

    // Sync params
    mutex       *sync.Mutex
    calcWaiter  *sync.WaitGroup
    readyWaiter *sync.WaitGroup
    ctx         *context.Context
}

func FindMinAnswersIteration(answers *Answers, begin int, params MinAnswerParams) {
    for i := begin; i < len(answers.skipped); i++ {
        // If must be skipped OR already exist
        if params.inSkipped[i] || !answers.skipped[i] {
            continue
        }

        params.calcWaiter.Add(1)
        params.readyWaiter.Add(1)
        func(i int) {
            // Init neuron
            neuron := CreateNeuron(params.inActivationFunc)

            // ... and answers set
            currentAnswers := *answers
            currentAnswers.skipped = make([]bool, len(answers.skipped))
            copy(currentAnswers.skipped, answers.skipped)
            currentAnswers.skipped[i] = false

            // If correct answers write result
            if neuron.Train(&currentAnswers, params.inShift, 0) {
                params.mutex.Lock()
                // If same compact answers, but younger neuron
                if !(*params.wasFound) || (neuron.age < params.outNeuron.age) {
                    *params.wasFound = true
                    *params.outNeuron = neuron
                    *params.outAnswers = currentAnswers
                }
                params.mutex.Unlock()
            }

            // Wait others in all goroutines
            params.calcWaiter.Done()
            <-(*params.ctx).Done()
            params.calcWaiter.Add(1)
            params.readyWaiter.Done()

            params.mutex.Lock()
            wasFound := *params.wasFound
            params.mutex.Unlock()

            if !wasFound {
                FindMinAnswersIteration(&currentAnswers, i, params)
            }
            params.calcWaiter.Done()
        }(i)
    }
}

func (n *Neuron) FindMinAnswers(answers *Answers, shift float64) Answers {
    beginAnswers := *answers
    beginAnswers.skipped = make([]bool, len(answers.skipped))
    for i := range answers.skipped {
        beginAnswers.skipped[i] = true
    }

    wasFound := false
    resultAnswers := *answers
    ctx, ctxCancel := context.WithCancel(context.Background())
    params := MinAnswerParams{
        inSkipped:        answers.skipped,
        inActivationFunc: n.activationFunc,
        inShift:          shift,

        wasFound:   &wasFound,
        outNeuron:  n,
        outAnswers: &resultAnswers,

        mutex:       &sync.Mutex{},
        calcWaiter:  &sync.WaitGroup{},
        readyWaiter: &sync.WaitGroup{},
        ctx:         &ctx,
    }

    FindMinAnswersIteration(&beginAnswers, 0, params)

    for i:= 1; i < 20; i++ {
        params.calcWaiter.Wait()
        if *params.wasFound {
            ctxCancel()
            break
        }

        params.mutex.Lock()
        ctxCancel()
        params.readyWaiter.Wait()
        ctx, ctxCancel = context.WithCancel(context.Background())
        *params.ctx = ctx
        params.mutex.Unlock()
    }

    return *params.outAnswers
}
