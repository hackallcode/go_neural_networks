package rnn

import (
    "fmt"
)

func MatrixToVector(matrix [][]int, w, h uint) []int {
    vector := make([]int, w*h)
    for j := uint(0); j < h; j++ {
        for i := uint(0); i < w; i++ {
            vector[h*i+j] = matrix[j][i]
        }
    }
    return vector
}

func BWLetterToVector(letter []string, w, h uint) []int {
    vector := make([]int, w*h)
    for j := uint(0); j < h; j++ {
        for i := uint(0); i < w; i++ {
            if letter[j][i] == BlackChar {
                vector[h*i+j] = 1
            } else {
                vector[h*i+j] = -1
            }
        }
    }
    return vector
}
func VectorToBWLetter(vector []int, w, h uint) []string {
    letter := make([]string, w*h)
    for j := uint(0); j < h; j++ {
        for i := uint(0); i < w; i++ {
            if vector[h*i+j] > 0 {
                letter[j] += string(BlackChar)
            } else {
                letter[j] += string(WhiteChar)
            }
        }
    }
    return letter
}

func PrintBWLetterByVector(vector []int, w, h uint) {
    fmt.Println("Result:")
    for j := uint(0); j < h; j++ {
        for i := uint(0); i < w; i++ {
            if vector[h*i+j] > 0 {
                fmt.Printf("%c", BlackChar)
            } else {
                fmt.Printf("%c", WhiteChar)
            }
        }
        fmt.Println()
    }
}
