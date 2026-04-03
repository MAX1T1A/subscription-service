package payment

import (
	paysvc "github.com/max1t1a/subscription-service/internal/service/payment"
)

type Handler struct {
	svc *paysvc.Service
}

func New(svc *paysvc.Service) *Handler {
	return &Handler{svc: svc}
}
