package neuron

import (
    "fmt"
    "math"
)

type AnswerFunc func([]uint8) uint8

type Answer struct {
    input  []uint8
    answer uint8
}

type Answers struct {
    varsCount uint
    skipped   []bool
    data      []Answer
}

func (a *Answers) PrintInfo() {
    for i, r := range a.data {
        if a.skipped[i] {
            fmt.Print("[ ]: ")
        } else {
            fmt.Print("[X]: ")
        }
        fmt.Print("F(")
        for i, v := range r.input[1:] {
            if i < len(r.input) - 2 {
                fmt.Printf("%v, ", v)
            } else {
                fmt.Printf("%v", v)
            }
        }
        fmt.Printf(") = %v\n", r.answer)
    }
}

func (a *Answers) Disable(index int) {
    a.skipped[index] = true
}

func (a *Answers) Enable(index int) {
    a.skipped[index] = false
}

// @param varsCount: number of variables
// @param shiftsCount: number of possible values
// @param shift: difference between neighboring values
func CreateAnswers(answerFunc AnswerFunc, varsCount uint, shiftsCount uint8, shift uint8) (result Answers) {
    rowsCount := uint(math.Pow(float64(shiftsCount), float64(varsCount)))

    result.varsCount = varsCount + 1
    result.skipped = make([]bool, rowsCount)

    row := make([]uint8, varsCount)
    for r := uint(0); r < rowsCount; r++ {

        // Current row with shifter input
        currRow := make([]uint8, varsCount+1)
        currRow[0] = 1
        copy(currRow[1:], row)

        // Remember result
        result.data = append(result.data, Answer{input: currRow, answer: answerFunc(row)})

        // Result next row
        for i := int(varsCount - 1); i >= 0; i-- {
            row[i] += shift
            if row[i] >= shiftsCount*shift {
                row[i] = 0
            } else {
                break
            }
        }

    }
    return
}
