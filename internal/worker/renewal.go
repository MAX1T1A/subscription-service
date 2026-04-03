package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
	paymentRepository "github.com/max1t1a/subscription-service/internal/repository/payment"
	subscriptionRepository "github.com/max1t1a/subscription-service/internal/repository/subscription"
)

type RenewalWorker struct {
	subRepo   *subscriptionRepository.Repository
	payRepo   *paymentRepository.Repository
	interval  time.Duration
	threshold time.Duration
}

func NewRenewalWorker(
	subRepo *subscriptionRepository.Repository,
	payRepo *paymentRepository.Repository,
	interval, threshold time.Duration,
) *RenewalWorker {
	return &RenewalWorker{
		subRepo:   subRepo,
		payRepo:   payRepo,
		interval:  interval,
		threshold: threshold,
	}
}

func (w *RenewalWorker) Start(ctx context.Context) {
	log.Printf("renewal worker started (interval: %s, threshold: %s)", w.interval, w.threshold)

	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("renewal worker stopped")
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
		log.Printf("failed to get expiring subscriptions: %v", err)
		return
	}

	if len(subs) == 0 {
		return
	}

	log.Printf("processing %d expiring subscription(s)", len(subs))

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
		log.Printf("failed to create renewal payment for subscription %s: %v", sub.ID, err)
		return
	}

	renewed, err := w.subRepo.Renew(ctx, sub.ID, durationSeconds)
	if err != nil {
		log.Printf("failed to renew subscription %s: %v", sub.ID, err)
		return
	}

	log.Printf("subscription %s renewed until %s (payment %s)", sub.ID, renewed.EndDate.Format("2006-01-02"), payment.ID)
}

func (w *RenewalWorker) expireSubscription(ctx context.Context, sub model.Subscription) {
	if err := w.subRepo.Expire(ctx, sub.ID); err != nil {
		log.Printf("failed to expire subscription %s: %v", sub.ID, err)
		return
	}

	log.Printf("subscription %s expired", sub.ID)
}
