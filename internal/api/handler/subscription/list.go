package subscription

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

// List godoc
// @Summary      List subscriptions
// @Tags         subscriptions
// @Produce      json
// @Param        user_id      query string false "Filter by user ID" format(uuid)
// @Param        service_name query string false "Filter by service name"
// @Param        status       query string false "Filter by status" Enums(active, expired)
// @Param        limit        query int    false "Page size (default 20)"
// @Param        offset       query int    false "Page offset (default 0)"
// @Success      200 {array} model.Subscription
// @Failure      400 {object} ErrorResponse
// @Router       /subscriptions [get]
func (h *Handler) List(c *gin.Context) {
	filter := model.SubscriptionFilter{
		Limit:  20,
		Offset: 0,
	}

	if v := c.Query("user_id"); v != "" {
		uid, err := uuid.Parse(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
			return
		}
		filter.UserID = &uid
	}

	if v := c.Query("service_name"); v != "" {
		filter.ServiceName = &v
	}

	if v := c.Query("status"); v != "" {
		st := model.SubscriptionStatus(v)
		if st != model.StatusActive && st != model.StatusExpired {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status, must be active or expired"})
			return
		}
		filter.Status = &st
	}

	if v := c.Query("limit"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}
		filter.Limit = n
	}

	if v := c.Query("offset"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
			return
		}
		filter.Offset = n
	}

	subs, err := h.svc.List(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subs)
}
