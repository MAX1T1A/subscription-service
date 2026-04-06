package subscription

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, req model.UpdateSubscriptionRequest) (*model.Subscription, error) {
	if req.EndDate != nil {
		if _, err := time.Parse("01-2006", *req.EndDate); err != nil {
			return nil, fmt.Errorf("invalid end_date format, expected MM-YYYY: %w", err)
		}
	}
	return s.repository.Update(ctx, id, req)
}
