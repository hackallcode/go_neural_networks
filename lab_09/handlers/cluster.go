package handlers

import (
    "encoding/json"
    "net/http"

    "lab_09/models"
    "lab_09/storage"
)

const (
    defaultMaxAge = 100
    defaultDistFunc = 1
)

func AddArea(w http.ResponseWriter, r *http.Request) {
    id, err := storage.AddArea()
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
        return
    }

    models.SendJson(w, http.StatusOK, models.GetAddAreaAnswer(&models.AddAreaAnswerData{Id: id}))
}

func AddPoint(w http.ResponseWriter, r *http.Request) {
    inputData := models.AddPointData{}
    err := json.NewDecoder(r.Body).Decode(&inputData)
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
        return
    }

    err = r.Body.Close()
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
        return
    }

    area, err := storage.GetArea(inputData.Id)
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
        return
    }

    for p := range inputData.Points {
        area.AddPoint(inputData.Points[p])
    }
    models.SendJson(w, http.StatusOK, models.GetSuccessAnswer("ok"))
}

func AddCluster(w http.ResponseWriter, r *http.Request) {
    inputData := models.AddClusterData{}
    err := json.NewDecoder(r.Body).Decode(&inputData)
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
        return
    }

    err = r.Body.Close()
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
        return
    }

    area, err := storage.GetArea(inputData.Id)
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
        return
    }

    for c := range inputData.Clusters {
        area.AddCluster(inputData.Clusters[c])
    }
    models.SendJson(w, http.StatusOK, models.GetSuccessAnswer("ok"))
}

func Train(w http.ResponseWriter, r *http.Request) {
    inputData := models.TrainData{}
    err := json.NewDecoder(r.Body).Decode(&inputData)
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
        return
    }

    if inputData.MaxAge == 0 {
        inputData.MaxAge = defaultMaxAge
    }
    if inputData.DistFuncId == 0 {
        inputData.DistFuncId = defaultDistFunc
    }

    err = r.Body.Close()
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
        return
    }

    area, err := storage.GetArea(inputData.Id)
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
        return
    }

    distFunc, err := storage.GetDistFunc(inputData.DistFuncId)
    if err != nil {
        models.SendJson(w, http.StatusInternalServerError, models.GetErrorAnswer(err.Error()))
        return
    }

    finished := false
    if inputData.ByStep {
        finished = area.TrainStep(distFunc)
    } else {
        finished = area.Train(distFunc, inputData.MaxAge)
    }

    models.SendJson(w, http.StatusOK, models.GetTrainAnswer(&models.TrainAnswerData{
        Finished: finished,
        Clusters: area.GetClustersWithPoints(distFunc),
    }))
}
