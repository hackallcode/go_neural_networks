package rnn

import (
    "fmt"
)

const (
    BlackChar = '*'
    WhiteChar = ' '
)

type RNN struct {
    w      [][]int
    k      uint
    maxAge uint
}

func CreateRNN(x [][]int, k uint, maxAge uint) (result RNN) {
    result.k = k
    result.maxAge = maxAge

    result.w = make([][]int, k)
    for i := uint(0); i < k; i++ {
        result.w[i] = make([]int, k)
        for j := uint(0); j < k; j++ {
            // After this cell will be only symmetric values
            if i == j {
                break
            }

            // Sum of images
            for l := range x {
                result.w[i][j] += x[l][i] * x[l][j]
            }

            // Save same in symmetric cell
            result.w[j][i] = result.w[i][j]
        }
    }

    return
}

func (n *RNN) PrintWeights() {
    fmt.Println("Weights:")
    for i := uint(0); i < n.k; i++ {
        for j := uint(0); j < n.k; j++ {
            if j == n.k-1 {
                fmt.Printf("%2v\n", n.w[i][j])
            } else {
                fmt.Printf("%2v, ", n.w[i][j])
            }
        }
    }
}

func (n *RNN) Detect(x []int, sync bool, printMode uint8) []int {
    out := make([]int, n.k)
    for i := range x {
        out[i] = x[i]
    }

    var prevOut []int
    var checkOut []int
    if sync {
        prevOut = make([]int, n.k)
        checkOut = prevOut
    } else {
        prevOut = out
        checkOut = make([]int, n.k)
    }

    for age := uint(0); age < n.maxAge; age++ {
        if printMode >= 1 {
            fmt.Printf("%v\n", out)
        }

        if sync {
            out, prevOut = prevOut, out
            checkOut = prevOut
        } else {
            copy(checkOut, prevOut)
        }

        for k := uint(0); k < n.k; k++ {
            net := 0
            for j := uint(0); j < n.k; j++ {
                if j == k {
                    continue
                }
                net += n.w[j][k] * prevOut[j]
            }

            if net < 0 {
                out[k] = -1
            } else if net > 0 {
                out[k] = 1
            } else {
                out[k] = prevOut[k]
            }
        }

        hasDiff := false
        for i := uint(0); i < n.k; i++ {
            if out[i] != checkOut[i] {
                hasDiff = true
            }
        }
        if !hasDiff {
            break
        }
    }

    return out
}
