package payment

import (
	"context"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (r *Repository) Create(ctx context.Context, p *model.Payment) error {
	query := `
		INSERT INTO payments (id, subscription_id, amount, status)
		VALUES ($1, $2, $3, $4)
		RETURNING paid_at`
	return r.db.QueryRowContext(ctx, query,
		p.ID, p.SubscriptionID, p.Amount, p.Status,
	).Scan(&p.PaidAt)
}
