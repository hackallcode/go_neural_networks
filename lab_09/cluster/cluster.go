package cluster

import (
    "fmt"
)

type Point struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type ClusterWithPoints struct {
    X      float64 `json:"x"`
    Y      float64 `json:"y"`
    Points []Point `json:"points"`
}

type Area struct {
    points   []Point
    clusters []Point
    age      uint
}

type avg struct {
    sum   Point
    count uint
}

func CreateArea() *Area {
    return &Area{
        points:   []Point{},
        clusters: []Point{},
    }
}

func (cl *Area) AddPoint(point Point) {
    cl.points = append(cl.points, point)
}

func (cl *Area) AddCluster(cluster Point) {
    cl.clusters = append(cl.clusters, cluster)
}

func (cl *Area) TrainStep(distFunc IDistance) bool {
    if len(cl.points) == 0 || len(cl.clusters) == 0 {
        return true
    }

    cl.age++
    newClusters := make([]avg, len(cl.clusters))

    for p := range cl.points {
        fmt.Printf("x_%v=(%v;%v); ", p+1, Round(cl.points[p].X, ResultAccuracy), Round(cl.points[p].Y, ResultAccuracy))
    }
    fmt.Println()

    for p := range cl.points {
        minDist := float64(-1)
        minCluster := -1
        for c := range cl.clusters {
            dist := distFunc(cl.points[p], cl.clusters[c])
            fmt.Printf("ρ_%v%v=%v; ", p+1, c+1, Round(dist, ResultAccuracy))
            if minCluster == -1 || dist < minDist {
                minDist = dist
                minCluster = c
            }
        }
        newClusters[minCluster].sum.X += cl.points[p].X
        newClusters[minCluster].sum.Y += cl.points[p].Y
        newClusters[minCluster].count++
    }
    fmt.Println()

    finished := true
    for c := range newClusters {
        oldX := cl.clusters[c].X
        oldY := cl.clusters[c].Y
        fmt.Printf("Y_%v=(%v;%v); ", c+1, Round(oldX, ResultAccuracy), Round(oldY, ResultAccuracy))

        if newClusters[c].count != 0 {
            cl.clusters[c].X = newClusters[c].sum.X / float64(newClusters[c].count)
            cl.clusters[c].Y = newClusters[c].sum.Y / float64(newClusters[c].count)
        }

        if (Round(cl.clusters[c].X-oldX, ResultAccuracy) != 0) || (Round(cl.clusters[c].Y-oldY, ResultAccuracy) != 0) {
            finished = false
        }
    }
    fmt.Println()

    return finished
}

func (cl *Area) Train(distFunc IDistance, maxAge uint) bool {
    for cl.age = 0; cl.age < maxAge; {
        if cl.TrainStep(distFunc) {
            return true
        }
    }
    return false
}

func (cl *Area) Clear() {
    cl.points = cl.points[:0]
    cl.clusters = cl.clusters[:0]
}

func (cl *Area) GetClusters() []Point {
    return cl.clusters
}

func (cl *Area) GetClustersWithPoints(distFunc IDistance) []ClusterWithPoints {
    if len(cl.clusters) == 0 {
        return []ClusterWithPoints{}
    }

    clusters := make([]ClusterWithPoints, len(cl.clusters))

    for c := range cl.clusters {
        clusters[c].X = cl.clusters[c].X
        clusters[c].Y = cl.clusters[c].Y
        clusters[c].Points = []Point{}
    }

    for p := range cl.points {
        minDist := float64(-1)
        minCluster := -1
        for c := range cl.clusters {
            dist := distFunc(cl.points[p], cl.clusters[c])
            if minCluster == -1 || dist < minDist {
                minDist = dist
                minCluster = c
            }
        }
        clusters[minCluster].Points = append(clusters[minCluster].Points, cl.points[p])
    }

    return clusters
}

func (cl *Area) Print() {
    for c := range cl.clusters {
        fmt.Printf("%v: (%v, %v)\n", c+1, Round(cl.clusters[c].X, ResultAccuracy), Round(cl.clusters[c].Y, ResultAccuracy))
    }
    fmt.Printf("Age: %v\n", cl.age)
}
