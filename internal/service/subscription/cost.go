package subscription

import (
	"context"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (s *Service) GetTotalCost(ctx context.Context, q model.CostQuery) (int, error) {
	return s.repo.GetTotalCost(ctx, q)
}
