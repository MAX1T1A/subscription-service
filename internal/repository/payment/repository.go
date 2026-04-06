package payment

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

type PaymentRepository interface {
	Create(ctx context.Context, p *model.Payment) error
	ListBySubscription(ctx context.Context, subID uuid.UUID, limit, offset int) ([]model.Payment, error)
}
