package subscription

import (
	"errors"
	"trafilea-prepare/domain/entity"
	"trafilea-prepare/domain/repository"
	"trafilea-prepare/domain/usecase"
)

type subscriptionUseCase struct {
	repo repository.SubscriptionRepository
}

// NewSubscriptionUseCase creates a new subscription use case
func NewSubscriptionUseCase(repo repository.SubscriptionRepository) usecase.SubscriptionUseCase {
	return &subscriptionUseCase{repo: repo}
}

// GetUserSubscriptions implements usecase.SubscriptionUseCase
func (s *subscriptionUseCase) GetUserSubscriptions(userID int) ([]entity.Subscription, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}
	return s.repo.GetByUserID(userID)
}

