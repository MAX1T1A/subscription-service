package payment

import (
	"context"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
	paymentRepository "github.com/max1t1a/subscription-service/internal/repository/payment"
)

type Service struct {
	repository *paymentRepository.Repository
}

func New(repository *paymentRepository.Repository) *Service {
	return &Service{repository: repository}
}

type Svc interface {
	ListBySubscription(ctx context.Context, subID uuid.UUID, limit, offset int) ([]model.Payment, error)
}
