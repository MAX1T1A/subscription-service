package subscription

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM subscriptions WHERE id = $1", id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return fmt.Errorf("subscription not found")
	}
	return nil
}
