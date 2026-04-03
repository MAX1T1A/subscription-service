package model

import "github.com/google/uuid"

type CreateSubscriptionRequest struct {
	ServiceName string    `json:"service_name" binding:"required"`
	Price       int       `json:"price" binding:"required,gt=0"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	StartDate   string    `json:"start_date" binding:"required"`
	EndDate     *string   `json:"end_date"`
	AutoRenew   *bool     `json:"auto_renew"`
}

type UpdateSubscriptionRequest struct {
	ServiceName *string `json:"service_name"`
	Price       *int    `json:"price" binding:"omitempty,gt=0"`
	EndDate     *string `json:"end_date"`
	AutoRenew   *bool   `json:"auto_renew"`
}

type SubscriptionFilter struct {
	UserID      *uuid.UUID
	ServiceName *string
	Status      *SubscriptionStatus
}

type CostQuery struct {
	UserID      *uuid.UUID
	ServiceName *string
	StartPeriod *string
	EndPeriod   *string
}

type CostResponse struct {
	TotalCost int `json:"total_cost"`
}
