package subscription

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetByID godoc
// @Summary      Get subscription by ID
// @Tags         subscriptions
// @Produce      json
// @Param        id path string true "Subscription ID" format(uuid)
// @Success      200 {object} model.Subscription
// @Failure      400 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Router       /subscriptions/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	sub, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	c.JSON(http.StatusOK, sub)
}
