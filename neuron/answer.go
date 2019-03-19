package neuron

import (
    "fmt"
    "math"
)

type AnswerFunc func([]float64) float64

type Answer struct {
    input  []float64
    answer float64
}

type LearningSet struct {
    varsCount uint
    skipped   []bool
    data      []Answer
}

func (a *LearningSet) PrintInfo() {
    for i, r := range a.data {
        if a.skipped[i] {
            fmt.Print("[ ]: ")
        } else {
            fmt.Print("[X]: ")
        }
        fmt.Print("F(")
        for i, v := range r.input[1:] {
            if i < len(r.input)-2 {
                fmt.Printf("%.5v, ", v)
            } else {
                fmt.Printf("%.5v", v)
            }
        }
        fmt.Printf(") = %.5v\n", r.answer)
    }
}

func (a *LearningSet) Disable(index int) {
    a.skipped[index] = true
}

func (a *LearningSet) Enable(index int) {
    a.skipped[index] = false
}

// @param varsCount: number of variables
// @param shiftsCount: number of possible values
// @param shift: difference between neighboring values
func CreateBoolLearningSet(answerFunc AnswerFunc, varsCount uint, shiftsCount uint8, shift uint8) (learningSet LearningSet) {
    rowsCount := uint(math.Pow(float64(shiftsCount), float64(varsCount)))

    learningSet.varsCount = varsCount + 1
    learningSet.skipped = make([]bool, rowsCount)

    row := make([]float64, varsCount)
    for r := uint(0); r < rowsCount; r++ {

        // Current row with shifter input
        currRow := make([]float64, varsCount+1)
        currRow[0] = 1
        copy(currRow[1:], row)

        // Remember result
        learningSet.data = append(learningSet.data, Answer{input: currRow, answer: answerFunc(row)})

        // Result next row
        for i := int(varsCount - 1); i >= 0; i-- {
            row[i] += float64(shift)
            if row[i] >= float64(shiftsCount)*float64(shift) {
                row[i] = 0
            } else {
                break
            }
        }

    }
    return
}

// @param varsCount: number of variables
// @param begin: begin of interval
// @param end: end of interval
// @param rowsCount: number of rows
func CreateFloatLearningSet(answerFunc AnswerFunc, varsCount uint, begin float64, end float64, rowsCount uint) (learningSet LearningSet) {
    shift := (end - begin) / float64(rowsCount)

    learningSet.varsCount = varsCount + 1
    learningSet.skipped = make([]bool, rowsCount)

    row := make([]float64, varsCount+1)
    row[0] = 1
    for i := uint(1); i < varsCount+1; i++ {
        row[i] = answerFunc([]float64{begin})
        begin += shift
    }
    learningSet.data = append(learningSet.data, Answer{input: row, answer: answerFunc([]float64{begin})})

    for r := varsCount; r < rowsCount; r++ {
        // Current row
        newRow := make([]float64, varsCount+1)
        newRow[0] = 1
        copy(newRow[1:], row[2:])
        newRow[len(newRow)-1] = answerFunc([]float64{begin})
        begin += shift

        // Remember result
        learningSet.data = append(learningSet.data, Answer{input: newRow, answer: answerFunc([]float64{begin})})
    }
    return
}
