package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"trafilea-prepare/service"
)

type SubscriptionController struct {
	service service.SubscriptionService
}

func NewSubscriptionController(s service.SubscriptionService) *SubscriptionController {
	return &SubscriptionController{service: s}
}

func (c *SubscriptionController) GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	subs, err := c.service.GetUserSubscriptions(userID)
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
