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

type floatParams struct {
	learningSet *LearningSet
	answerFunc  AnswerFunc
	width       uint
	points      uint
	begin       float64
	shift       float64
	skip        bool
}

func subFloatLearningSet(params floatParams) {
	row := make([]float64, params.width+1)
	row[0] = 1
	for i := uint(1); i < params.width+1; i++ {
		row[i] = params.answerFunc([]float64{params.begin})
		params.begin += params.shift
	}
	answer := params.answerFunc([]float64{params.begin})
	params.begin += params.shift

	for r := uint(1); ; r++ {
		// Remember result
		params.learningSet.data = append(params.learningSet.data, Answer{input: row, answer: answer})
		params.learningSet.skipped = append(params.learningSet.skipped, params.skip)

		if r == params.points-params.width {
			break
		}

		// Next row
		newRow := make([]float64, params.width+1)
		newRow[0] = 1
		copy(newRow[1:], row[2:])
		newRow[len(newRow)-1] = answer
		row = newRow
		answer = params.answerFunc([]float64{params.begin})
		params.begin += params.shift
	}
}

// @param width: width of window
// @param begin: begin of interval
// @param end: end of interval
// @param points: number of learning points
func CreateFloatLearningSet(answerFunc AnswerFunc, width uint, begin float64, end float64, points uint) (learningSet LearningSet) {
	if width >= points {
		return
	}

	learningSet.varsCount = width + 1
	params := floatParams{
		learningSet: &learningSet,
		answerFunc:  answerFunc,
		width:       width,
		points:      points,
		shift:       (end - begin) / float64(points-1),
	}

	params.begin = begin
	params.skip = false
	subFloatLearningSet(params)

	params.begin = end
	params.skip = true
	subFloatLearningSet(params)
	return
}
