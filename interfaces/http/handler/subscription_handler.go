package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"trafilea-prepare/domain/usecase"
)

type SubscriptionHandler struct {
	useCase usecase.SubscriptionUseCase
}

// NewSubscriptionHandler creates a new subscription HTTP handler
func NewSubscriptionHandler(useCase usecase.SubscriptionUseCase) *SubscriptionHandler {
	return &SubscriptionHandler{useCase: useCase}
}

// GetSubscriptions handles GET /subscriptions requests
func (h *SubscriptionHandler) GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	subs, err := h.useCase.GetUserSubscriptions(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(subs) == 0 {
		http.Error(w, "No subscriptions found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(subs)
}

