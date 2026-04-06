package subscription

import (
	"context"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (r *Repository) GetExpiring(ctx context.Context, threshold string) ([]model.Subscription, error) {
	query := `
		SELECT id, service_name, price, user_id, start_date, end_date, auto_renew, status, created_at, updated_at FROM subscriptions
		WHERE status = 'active'
		  AND end_date <= now() + $1::interval
		ORDER BY end_date ASC`

	var subs []model.Subscription
	err := r.db.SelectContext(ctx, &subs, query, threshold)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (r *Repository) Renew(ctx context.Context, id uuid.UUID, durationSeconds int) (*model.Subscription, error) {
	query := `
		UPDATE subscriptions
		SET end_date = end_date + make_interval(secs => $1),
		    updated_at = now()
		WHERE id = $2
		RETURNING id, service_name, price, user_id, start_date, end_date, auto_renew, status, created_at, updated_at`

	var s model.Subscription
	err := r.db.GetContext(ctx, &s, query, durationSeconds, id)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *Repository) Expire(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE subscriptions SET status = 'expired', updated_at = now() WHERE id = $1", id)
	return err
}
