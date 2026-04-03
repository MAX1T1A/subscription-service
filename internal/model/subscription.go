package model

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionStatus string

const (
	StatusActive  SubscriptionStatus = "active"
	StatusExpired SubscriptionStatus = "expired"
)

type Subscription struct {
	ID          uuid.UUID          `json:"id" db:"id"`
	ServiceName string             `json:"service_name" db:"service_name"`
	Price       int                `json:"price" db:"price"`
	UserID      uuid.UUID          `json:"user_id" db:"user_id"`
	StartDate   time.Time          `json:"start_date" db:"start_date"`
	EndDate     time.Time          `json:"end_date" db:"end_date"`
	AutoRenew   bool               `json:"auto_renew" db:"auto_renew"`
	Status      SubscriptionStatus `json:"status" db:"status"`
	CreatedAt   time.Time          `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" db:"updated_at"`
}
