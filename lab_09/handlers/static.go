package handlers

import (
    "io/ioutil"
    "net/http"

    "lab_09/models"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
    text, err := ioutil.ReadFile("static/index.html")
    if err != nil {
        models.SendText(w, http.StatusOK, "Index page not found!")
    }
    models.SendText(w, http.StatusOK, string(text))
}

func ApiHomeHandler(w http.ResponseWriter, _ *http.Request) {
    models.SendJson(w, http.StatusOK, models.GetSuccessAnswer("Main page!"))
}
