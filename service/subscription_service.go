package service

import (
	"errors"
	"trafilea-prepare/model"
	"trafilea-prepare/repository"
)

type SubscriptionService interface {
	GetUserSubscriptions(userID int) ([]model.Subscription, error)
}

type subscriptionService struct {
	repo repository.SubscriptionRepository
}

// GetUserSubscriptions implements SubscriptionService.
func (s *subscriptionService) GetUserSubscriptions(userID int) ([]model.Subscription, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}
	return s.repo.GetByUserID(userID)
}

func NewSubscriptionService(repo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{repo: repo}
}
