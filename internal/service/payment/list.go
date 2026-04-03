package payment

import (
	"context"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (s *Service) ListBySubscription(ctx context.Context, subID uuid.UUID) ([]model.Payment, error) {
	return s.repo.ListBySubscription(ctx, subID)
}
