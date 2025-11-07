package repository

import (
	"trafilea-prepare/db"
	"trafilea-prepare/model"
)

type SubscriptionRepository interface {
	GetByUserID(userID int) ([]model.Subscription, error)
}

type subscriptionRepository struct{}

func NewSubscriptionRepository() SubscriptionRepository {
	return &subscriptionRepository{}
}

func (r *subscriptionRepository) GetByUserID(userID int) ([]model.Subscription, error) {
	var result []model.Subscription

	for _, s := range db.Subscriptions {
		if s.UserID == userID {
			result = append(result, s)
		}
	}
	return result, nil
}
