package subscription

import (
	"context"
	"fmt"

	"github.com/max1t1a/subscription-service/internal/model"
)

func (r *Repository) GetTotalCost(ctx context.Context, q model.CostQuery) (int, error) {
	// $1 = end_period (exclusive upper bound), $2 = start_period
	query := `
		SELECT COALESCE(SUM(
			price * (
				EXTRACT(YEAR FROM AGE(LEAST(end_date, $1), GREATEST(start_date, $2)))::int * 12 +
				EXTRACT(MONTH FROM AGE(LEAST(end_date, $1), GREATEST(start_date, $2)))::int
			)
		), 0)
		FROM subscriptions
		WHERE start_date < $1 AND end_date > $2`

	args := []interface{}{q.EndPeriod, q.StartPeriod}
	idx := 3

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

	var total int
	err := r.db.GetContext(ctx, &total, query, args...)
	return total, err
}
