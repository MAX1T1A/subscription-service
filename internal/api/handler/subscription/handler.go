package subscription

import (
	subscriptionService "github.com/max1t1a/subscription-service/internal/service/subscription"
)

type Handler struct {
	svc *subscriptionService.Service
}

func New(svc *subscriptionService.Service) *Handler {
	return &Handler{svc: svc}
}
