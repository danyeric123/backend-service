package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	DB *sqlx.DB
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{DB: db}
}

// HealthCheck is a handler for health checks
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	logrus.Println("Health check")
	w.WriteHeader(http.StatusOK)
}
