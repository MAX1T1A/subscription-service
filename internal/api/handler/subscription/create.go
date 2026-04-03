package subscription

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/max1t1a/subscription-service/internal/model"
)

// Create godoc
// @Summary      Create subscription
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        body body model.CreateSubscriptionRequest true "Subscription data"
// @Success      201 {object} model.Subscription
// @Failure      400 {object} ErrorResponse
// @Router       /subscriptions [post]
func (h *Handler) Create(c *gin.Context) {
	var req model.CreateSubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, err := h.svc.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sub)
}
