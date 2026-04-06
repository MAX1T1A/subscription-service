package payment

import (
	"net/http"
	"strconv"

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
// @Param        subscription_id query string true  "Subscription ID" format(uuid)
// @Param        limit           query int    false "Page size (default 20)"
// @Param        offset          query int    false "Page offset (default 0)"
// @Success      200 {array} model.Payment
// @Failure      400 {object} ErrorResponse
// @Router       /payments [get]
func (h *Handler) ListBySubscription(c *gin.Context) {
	subID, err := uuid.Parse(c.Query("subscription_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing subscription_id"})
		return
	}

	limit := 20
	offset := 0

	if v := c.Query("limit"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}
		limit = n
	}

	if v := c.Query("offset"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
			return
		}
		offset = n
	}

	payments, err := h.svc.ListBySubscription(c.Request.Context(), subID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}
