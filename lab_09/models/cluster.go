package models

import (
    "lab_09/cluster"
)

type AddPointData struct {
    Id     int             `json:"id"`
    Points []cluster.Point `json:"points"`
}

type AddClusterData struct {
    Id       int             `json:"id"`
    Clusters []cluster.Point `json:"clusters"`
}

type TrainData struct {
    Id         int  `json:"id"`
    MaxAge     uint `json:"max_age"`
    ByStep     bool `json:"by_step"`
    DistFuncId int  `json:"dist_id"`
}

type ClearAreaData struct {
    Id       int             `json:"id"`
}

type AddAreaAnswerData struct {
    Id int `json:"id"`
}

type TrainAnswerData struct {
    Finished bool                        `json:"finished"`
    Clusters []cluster.ClusterWithPoints `json:"clusters"`
}
