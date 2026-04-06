package subscription

import (
	"context"
	"fmt"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (r *Repository) List(ctx context.Context, filter model.SubscriptionFilter) ([]model.Subscription, error) {
	query := "SELECT id, service_name, price, user_id, start_date, end_date, auto_renew, status, created_at, updated_at FROM subscriptions WHERE 1=1"
	args := []interface{}{}
	idx := 1

	if filter.UserID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", idx)
		args = append(args, *filter.UserID)
		idx++
	}
	if filter.ServiceName != nil {
		query += fmt.Sprintf(" AND service_name = $%d", idx)
		args = append(args, *filter.ServiceName)
		idx++
	}
	if filter.Status != nil {
		query += fmt.Sprintf(" AND status = $%d", idx)
		args = append(args, *filter.Status)
		idx++
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", idx, idx+1)
	args = append(args, filter.Limit, filter.Offset)

	var subs []model.Subscription
	err := r.db.SelectContext(ctx, &subs, query, args...)
	if err != nil {
		return nil, err
	}
	return subs, nil
}
