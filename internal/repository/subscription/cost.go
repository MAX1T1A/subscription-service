package subscription

import (
	"context"
	"fmt"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (r *Repository) GetTotalCost(ctx context.Context, q model.CostQuery) (int, error) {
	query := "SELECT COALESCE(SUM(price), 0) FROM subscriptions WHERE 1=1"
	args := []interface{}{}
	idx := 1

	if q.UserID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", idx)
		args = append(args, *q.UserID)
		idx++
	}
	if q.ServiceName != nil {
		query += fmt.Sprintf(" AND service_name = $%d", idx)
		args = append(args, *q.ServiceName)
		idx++
	}
	if q.StartPeriod != nil {
		query += fmt.Sprintf(" AND start_date >= $%d", idx)
		args = append(args, *q.StartPeriod)
		idx++
	}
	if q.EndPeriod != nil {
		query += fmt.Sprintf(" AND start_date <= $%d", idx)
		args = append(args, *q.EndPeriod)
		idx++
	}

	var total int
	err := r.db.GetContext(ctx, &total, query, args...)
	return total, err
}
