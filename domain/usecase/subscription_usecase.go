package usecase

import "trafilea-prepare/domain/entity"

// SubscriptionUseCase defines the interface for subscription business logic
// This interface belongs to the domain layer
type SubscriptionUseCase interface {
	GetUserSubscriptions(userID int) ([]entity.Subscription, error)
}

