package cluster

import (
    "math"
)

const ResultAccuracy = 100

// 10 - 1 digit after addPoint
// 100 - 2 digits after addPoint
// ...
func Round(x float64, accuracy float64) float64 {
    return math.Round(x*accuracy) / accuracy
}
