package subscription

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (r *Repository) Update(ctx context.Context, id uuid.UUID, req model.UpdateSubscriptionRequest) (*model.Subscription, error) {
	sets := []string{}
	args := []interface{}{}
	idx := 1

	if req.ServiceName != nil {
		sets = append(sets, fmt.Sprintf("service_name = $%d", idx))
		args = append(args, *req.ServiceName)
		idx++
	}
	if req.Price != nil {
		sets = append(sets, fmt.Sprintf("price = $%d", idx))
		args = append(args, *req.Price)
		idx++
	}
	if req.EndDate != nil {
		sets = append(sets, fmt.Sprintf("end_date = $%d", idx))
		args = append(args, *req.EndDate)
		idx++
	}
	if req.AutoRenew != nil {
		sets = append(sets, fmt.Sprintf("auto_renew = $%d", idx))
		args = append(args, *req.AutoRenew)
		idx++
	}

	if len(sets) == 0 {
		return r.GetByID(ctx, id)
	}

	sets = append(sets, "updated_at = now()")
	args = append(args, id)

	query := fmt.Sprintf(
		"UPDATE subscriptions SET %s WHERE id = $%d RETURNING *",
		strings.Join(sets, ", "), idx,
	)

	var s model.Subscription
	err := r.db.GetContext(ctx, &s, query, args...)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
