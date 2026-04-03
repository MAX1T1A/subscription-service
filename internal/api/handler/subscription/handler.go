package subscription

import (
	subsvc "github.com/max1t1a/subscription-service/internal/service/subscription"
)

type Handler struct {
	svc *subsvc.Service
}

func New(svc *subsvc.Service) *Handler {
	return &Handler{svc: svc}
}
