package subscription

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/max1t1a/subscription-service/internal/model"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

type SubscriptionRepository interface {
	Create(ctx context.Context, s *model.Subscription) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Subscription, error)
	List(ctx context.Context, filter model.SubscriptionFilter) ([]model.Subscription, error)
	Update(ctx context.Context, id uuid.UUID, req model.UpdateSubscriptionRequest) (*model.Subscription, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetTotalCost(ctx context.Context, q model.CostQuery) (int, error)
	GetExpiring(ctx context.Context, threshold string) ([]model.Subscription, error)
	Renew(ctx context.Context, id uuid.UUID, durationSeconds int) (*model.Subscription, error)
	Expire(ctx context.Context, id uuid.UUID) error
}
