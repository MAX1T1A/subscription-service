package subscription

import (
	"context"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (s *Service) List(ctx context.Context, filter model.SubscriptionFilter) ([]model.Subscription, error) {
	return s.repository.List(ctx, filter)
}
