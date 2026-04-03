package payment

import (
	"context"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
	paymentRepository "github.com/max1t1a/subscription-service/internal/repository/payment"
)

type Service struct {
	repo *paymentRepository.Repository
}

func New(repo *paymentRepository.Repository) *Service {
	return &Service{repo: repo}
}

type Svc interface {
	ListBySubscription(ctx context.Context, subID uuid.UUID) ([]model.Payment, error)
}
