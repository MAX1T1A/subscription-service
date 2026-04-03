package payment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ErrorResponse represents an error returned by the API.
type ErrorResponse struct {
	Error string `json:"error" example:"something went wrong"`
}

// ListBySubscription godoc
// @Summary      List payments by subscription
// @Tags         payments
// @Produce      json
// @Param        subscription_id query string true "Subscription ID" format(uuid)
// @Success      200 {array} model.Payment
// @Failure      400 {object} ErrorResponse
// @Router       /payments [get]
func (h *Handler) ListBySubscription(c *gin.Context) {
	subID, err := uuid.Parse(c.Query("subscription_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing subscription_id"})
		return
	}

	payments, err := h.svc.ListBySubscription(c.Request.Context(), subID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}
