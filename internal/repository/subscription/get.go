package subscription

import (
	"context"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*model.Subscription, error) {
	var s model.Subscription
	err := r.db.GetContext(ctx, &s, "SELECT * FROM subscriptions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
