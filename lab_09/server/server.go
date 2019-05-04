package server

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"

    "lab_09/handlers"
    "lab_09/middleware"
)

func Start(port string) error {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.HomeHandler)

    apiRouter := router.PathPrefix("/api").Subrouter()
    apiRouter.Use(middleware.ApplyJsonContentType)

    apiRouter.HandleFunc("/", handlers.ApiHomeHandler)
    apiRouter.HandleFunc("/area", handlers.AddArea).Methods("POST")
    apiRouter.HandleFunc("/point", handlers.AddPoint).Methods("POST")
    apiRouter.HandleFunc("/cluster", handlers.AddCluster).Methods("POST")
    apiRouter.HandleFunc("/train", handlers.Train).Methods("POST")

    fs := http.FileServer(http.Dir("static"))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

    fmt.Println("Server listening at " + port)
    return http.ListenAndServe(":"+port, router)
}

func Stop() {
    fmt.Println("Stopping server...")
}
