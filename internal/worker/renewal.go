package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/max1t1a/subscription-service/internal/model"
	paymentRepository "github.com/max1t1a/subscription-service/internal/repository/payment"
	subscriptionRepository "github.com/max1t1a/subscription-service/internal/repository/subscription"
)

type RenewalWorker struct {
	subRepo   *subscriptionRepository.Repository
	payRepo   *paymentRepository.Repository
	interval  time.Duration
	threshold time.Duration
	logger    *zap.Logger
}

func NewRenewalWorker(
	subRepo *subscriptionRepository.Repository,
	payRepo *paymentRepository.Repository,
	interval, threshold time.Duration,
	logger *zap.Logger,
) *RenewalWorker {
	return &RenewalWorker{
		subRepo:   subRepo,
		payRepo:   payRepo,
		interval:  interval,
		threshold: threshold,
		logger:    logger,
	}
}

func (w *RenewalWorker) Start(ctx context.Context) {
	w.logger.Info("renewal worker started",
		zap.Duration("interval", w.interval),
		zap.Duration("threshold", w.threshold),
	)

	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("renewal worker stopped")
			return
		case <-ticker.C:
			w.processExpiring(ctx)
		}
	}
}

func (w *RenewalWorker) processExpiring(ctx context.Context) {
	thresholdStr := fmt.Sprintf("%d seconds", int(w.threshold.Seconds()))

	subs, err := w.subRepo.GetExpiring(ctx, thresholdStr)
	if err != nil {
		w.logger.Error("failed to get expiring subscriptions", zap.Error(err))
		return
	}

	if len(subs) == 0 {
		return
	}

	w.logger.Info("processing expiring subscriptions", zap.Int("count", len(subs)))

	for _, sub := range subs {
		if sub.AutoRenew {
			w.renewSubscription(ctx, sub)
		} else {
			w.expireSubscription(ctx, sub)
		}
	}
}

func (w *RenewalWorker) renewSubscription(ctx context.Context, sub model.Subscription) {
	durationSeconds := int(sub.EndDate.Sub(sub.StartDate).Seconds())

	payment := &model.Payment{
		ID:             uuid.New(),
		SubscriptionID: sub.ID,
		Amount:         sub.Price,
		Status:         model.PaymentStatusSuccess,
	}

	if err := w.payRepo.Create(ctx, payment); err != nil {
		w.logger.Error("failed to create renewal payment",
			zap.String("subscription_id", sub.ID.String()),
			zap.Error(err),
		)
		return
	}

	renewed, err := w.subRepo.Renew(ctx, sub.ID, durationSeconds)
	if err != nil {
		w.logger.Error("failed to renew subscription",
			zap.String("subscription_id", sub.ID.String()),
			zap.Error(err),
		)
		return
	}

	w.logger.Info("subscription renewed",
		zap.String("subscription_id", sub.ID.String()),
		zap.String("new_end_date", renewed.EndDate.Format("2006-01-02")),
		zap.String("payment_id", payment.ID.String()),
	)
}

func (w *RenewalWorker) expireSubscription(ctx context.Context, sub model.Subscription) {
	if err := w.subRepo.Expire(ctx, sub.ID); err != nil {
		w.logger.Error("failed to expire subscription",
			zap.String("subscription_id", sub.ID.String()),
			zap.Error(err),
		)
		return
	}

	w.logger.Info("subscription expired", zap.String("subscription_id", sub.ID.String()))
}
