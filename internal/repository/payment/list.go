package payment

import (
	"context"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (r *Repository) ListBySubscription(ctx context.Context, subID uuid.UUID) ([]model.Payment, error) {
	var payments []model.Payment
	err := r.db.SelectContext(ctx, &payments,
		"SELECT * FROM payments WHERE subscription_id = $1 ORDER BY paid_at DESC", subID)
	if err != nil {
		return nil, err
	}
	return payments, nil
}
