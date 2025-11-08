package repository

import "trafilea-prepare/domain/entity"

// SubscriptionRepository defines the interface for subscription data access
// This interface belongs to the domain layer and will be implemented by infrastructure
type SubscriptionRepository interface {
	GetByUserID(userID int) ([]entity.Subscription, error)
}
