package payment

import (
	paymentService "github.com/max1t1a/subscription-service/internal/service/payment"
)

type Handler struct {
	svc *paymentService.Service
}

func New(svc *paymentService.Service) *Handler {
	return &Handler{svc: svc}
}
