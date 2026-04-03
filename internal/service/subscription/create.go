package subscription

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (s *Service) Create(ctx context.Context, req model.CreateSubscriptionRequest) (*model.Subscription, error) {
	startDate, err := time.Parse("01-2006", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start_date format, expected MM-YYYY: %w", err)
	}

	var endDate time.Time
	if req.EndDate != nil {
		endDate, err = time.Parse("01-2006", *req.EndDate)
		if err != nil {
			return nil, fmt.Errorf("invalid end_date format, expected MM-YYYY: %w", err)
		}
	} else {
		endDate = startDate.AddDate(0, 1, 0)
	}

	if endDate.Before(startDate) {
		return nil, fmt.Errorf("end_date must be after start_date")
	}

	autoRenew := false
	if req.AutoRenew != nil {
		autoRenew = *req.AutoRenew
	}

	sub := &model.Subscription{
		ID:          uuid.New(),
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   startDate,
		EndDate:     endDate,
		AutoRenew:   autoRenew,
		Status:      model.StatusActive,
	}

	if err := s.repo.Create(ctx, sub); err != nil {
		return nil, fmt.Errorf("failed to create subscription: %w", err)
	}

	return sub, nil
}
