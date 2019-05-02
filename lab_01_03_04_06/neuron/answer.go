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
	isTest    []bool
	data      []Answer
}

func (a *LearningSet) PrintInfo() {
	for i, r := range a.data {
		if a.isTest[i] {
			fmt.Print("[ ]: ")
		} else {
			fmt.Print("[L]: ")
		}
		fmt.Print("F(")
		for i, v := range r.input[1:] {
			if i < len(r.input)-2 {
				fmt.Printf("%v, ", Round(v, ResultAccuracy))
			} else {
				fmt.Printf("%v", Round(v, ResultAccuracy))
			}
		}
		fmt.Printf(") = %v\n", Round(r.answer, ResultAccuracy))
	}
}

func (a *LearningSet) PrintResults() {
	for i, r := range a.data {
		fmt.Printf("%v", Round(r.answer, ResultAccuracy))
		if i < len(a.data)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func (a *LearningSet) Disable(index int) {
	a.isTest[index] = true
}

func (a *LearningSet) DisableByArray(indexes []int) {
	for _, i := range indexes {
		a.isTest[i] = true
	}
}

func (a *LearningSet) Enable(index int) {
	a.isTest[index] = false
}

// @param varsCount: number of variables
// @param shiftsCount: number of possible values
// @param shift: difference between neighboring values
func CreateBoolLearningSet(answerFunc AnswerFunc, varsCount uint, shiftsCount uint8, shift uint8) (learningSet LearningSet) {
	rowsCount := uint(math.Pow(float64(shiftsCount), float64(varsCount)))

	learningSet.varsCount = varsCount + 1
	learningSet.isTest = make([]bool, rowsCount)

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
// @param shiftsCount: number of possible values
// @param shift: difference between neighboring values
func CreateRBFBoolLearningSet(answerFunc AnswerFunc, varsCount uint, centers [][]float64) (learningSet LearningSet) {
	rbfCount := len(centers)
	rowsCount := uint(math.Pow(float64(2), float64(varsCount)))

	learningSet.varsCount = uint(rbfCount + 1)
	learningSet.isTest = make([]bool, rowsCount)

	row := make([]float64, varsCount)
	for r := uint(0); r < rowsCount; r++ {
		currInput := make([]float64, learningSet.varsCount)
		for j := range currInput {
			if j == 0 {
				currInput[j] = 1
				continue
			}

			currInput[j] = 0
			for i := uint(0); i < varsCount; i++ {
				currInput[j] += math.Pow(row[i] - centers[j - 1][i], 2)
			}
			currInput[j] = math.Exp(-currInput[j])
		}

		// Remember result
		learningSet.data = append(learningSet.data, Answer{input: currInput, answer: answerFunc(row)})

		// Result next row
		for i := int(varsCount - 1); i >= 0; i-- {
			row[i] += 1
			if row[i] >= float64(2) {
				row[i] = 0
			} else {
				break
			}
		}

	}
	return
}

// @param width: width of window
// @param begin: begin of interval
// @param end: end of interval
// @param points: number of learning points
func CreateFloatLearningSet(answerFunc AnswerFunc, width uint, begin float64, end float64, points uint) (learningSet LearningSet) {
	if width >= points {
		return
	}

	shift := (end - begin) / float64(points-1)
	learningSet.varsCount = width + 1

	row := make([]float64, width+1)
	row[0] = 1
	for i := uint(1); i < width+1; i++ {
		row[i] = answerFunc([]float64{begin})
		begin += shift
	}
	answer := answerFunc([]float64{begin})
	begin += shift

	for r := width; ; r++ {
		// Remember result
		learningSet.data = append(learningSet.data, Answer{input: row, answer: answer})
		if r >= points {
			learningSet.isTest = append(learningSet.isTest, true)
		} else {
			learningSet.isTest = append(learningSet.isTest, false)
		}

		if r == points*2-1 {
			break
		}

		// Next row
		newRow := make([]float64, width+1)
		newRow[0] = 1
		copy(newRow[1:], row[2:])
		newRow[len(newRow)-1] = answer
		row = newRow
		answer = answerFunc([]float64{begin})
		begin += shift
	}
	return
}
