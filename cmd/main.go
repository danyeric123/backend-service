package main

import (
	"net/http"
	"os"

	"github.com/danyeric123/backend-service/db"
	"github.com/danyeric123/backend-service/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	db.Connect()

	h := handlers.NewHandler(db.DB)

	r := mux.NewRouter()
	r.HandleFunc("/healthz", h.HealthCheck).Methods("GET")
	logrus.Info("Starting server on :8080")
	logrus.Fatal(http.ListenAndServe(":8080", r))
}
