package subscription

import (
	"context"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (r *Repository) Create(ctx context.Context, s *model.Subscription) error {
	query := `
		INSERT INTO subscriptions (id, service_name, price, user_id, start_date, end_date, auto_renew, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING created_at, updated_at`
	return r.db.QueryRowContext(ctx, query,
		s.ID, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate, s.AutoRenew, s.Status,
	).Scan(&s.CreatedAt, &s.UpdatedAt)
}
