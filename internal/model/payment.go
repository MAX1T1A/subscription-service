package model

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	PaymentStatusSuccess PaymentStatus = "success"
	PaymentStatusFailed  PaymentStatus = "failed"
)

type Payment struct {
	ID             uuid.UUID     `json:"id" db:"id"`
	SubscriptionID uuid.UUID     `json:"subscription_id" db:"subscription_id"`
	Amount         int           `json:"amount" db:"amount"`
	Status         PaymentStatus `json:"status" db:"status"`
	PaidAt         time.Time     `json:"paid_at" db:"paid_at"`
}
