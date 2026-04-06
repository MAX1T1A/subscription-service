package subscription

import (
	"context"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
	subscriptionRepository "github.com/max1t1a/subscription-service/internal/repository/subscription"
)

type Service struct {
	repository *subscriptionRepository.Repository
}

func New(repository *subscriptionRepository.Repository) *Service {
	return &Service{repository: repository}
}

type Svc interface {
	Create(ctx context.Context, req model.CreateSubscriptionRequest) (*model.Subscription, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Subscription, error)
	List(ctx context.Context, filter model.SubscriptionFilter) ([]model.Subscription, error)
	Update(ctx context.Context, id uuid.UUID, req model.UpdateSubscriptionRequest) (*model.Subscription, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetTotalCost(ctx context.Context, q model.CostQuery) (int, error)
}
