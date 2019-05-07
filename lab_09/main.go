package main

import (
    "fmt"

    "lab_09/server"
)

func main() {
    // cl := cluster.CreateArea()
    // cl.AddPoint(cluster.Point{X: 143, Y: 213})
    // cl.AddPoint(cluster.Point{X: 180, Y: 220})
    // cl.AddPoint(cluster.Point{X: 183, Y: 249})
    // cl.AddPoint(cluster.Point{X: 271, Y: 253})
    // cl.AddPoint(cluster.Point{X: 226, Y: 253})
    // cl.AddPoint(cluster.Point{X: 315, Y: 275})
    // cl.AddPoint(cluster.Point{X: 266, Y: 297})
    // cl.AddCluster(cluster.Point{X: 159, Y: 238})
    // cl.AddCluster(cluster.Point{X: 270, Y: 278})
    // cl.Train(cluster.EuclideanDistance, 100)
    // cl.Print()

    err := server.Start("4000")
    if err != nil {
        fmt.Println(err.Error())
    }
}
