package handler

import (
	"encoding/json"
	"net/http"

	"server/internal/application"
	"server/internal/domain"
)

type LocationHandler struct {
	service *application.LocationService
}

func NewLocationHandler(service *application.LocationService) *LocationHandler {
	return &LocationHandler{service: service}
}

func (h *LocationHandler) EstimateLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.DeviceLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.Estimate(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}