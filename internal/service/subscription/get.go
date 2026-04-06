package subscription

import (
	"context"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*model.Subscription, error) {
	return s.repository.GetByID(ctx, id)
}
