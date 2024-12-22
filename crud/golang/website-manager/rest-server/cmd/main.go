package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "rest-server/internal/controller"
)

func main() {
    r := mux.NewRouter()

    // Define routes for CRUD operations
    r.HandleFunc("/api/websites", controller.CreateWebsite).Methods("POST")
    r.HandleFunc("/api/websites", controller.GetWebsites).Methods("GET")
    r.HandleFunc("/api/websites/{id}", controller.UpdateWebsite).Methods("PUT")
    r.HandleFunc("/api/websites/{id}", controller.DeleteWebsite).Methods("DELETE")

    // Start the server
    log.Println("Starting REST server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}