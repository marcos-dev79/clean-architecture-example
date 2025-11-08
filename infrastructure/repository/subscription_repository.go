package repository

import (
	"trafilea-prepare/domain/entity"
	"trafilea-prepare/domain/repository"
	"trafilea-prepare/infrastructure/db"
)

type subscriptionRepository struct{}

// NewSubscriptionRepository creates a new subscription repository implementation
func NewSubscriptionRepository() repository.SubscriptionRepository {
	return &subscriptionRepository{}
}

// GetByUserID implements repository.SubscriptionRepository
func (r *subscriptionRepository) GetByUserID(userID int) ([]entity.Subscription, error) {
	var result []entity.Subscription

	for _, s := range db.Subscriptions {
		if s.UserID == userID {
			result = append(result, entity.Subscription{
				ID:      s.ID,
				UserID:  s.UserID,
				Service: s.Service,
				Status:  s.Status,
			})
		}
	}
	return result, nil
}

