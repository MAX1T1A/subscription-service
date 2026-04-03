package subscription

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

// Update godoc
// @Summary      Update subscription
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path string true "Subscription ID" format(uuid)
// @Param        body body model.UpdateSubscriptionRequest true "Fields to update"
// @Success      200 {object} model.Subscription
// @Failure      400 {object} ErrorResponse
// @Router       /subscriptions/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req model.UpdateSubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, err := h.svc.Update(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sub)
}
