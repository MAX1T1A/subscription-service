package subscription

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/max1t1a/subscription-service/internal/model"
)

// GetTotalCost godoc
// @Summary      Get total subscription cost for a period
// @Tags         subscriptions
// @Produce      json
// @Param        user_id      query string false "Filter by user ID" format(uuid)
// @Param        service_name query string false "Filter by service name"
// @Param        start_period query string false "Period start (MM-YYYY)"
// @Param        end_period   query string false "Period end (MM-YYYY)"
// @Success      200 {object} model.CostResponse
// @Failure      400 {object} ErrorResponse
// @Router       /subscriptions/cost [get]
func (h *Handler) GetTotalCost(c *gin.Context) {
	q := model.CostQuery{}

	if v := c.Query("user_id"); v != "" {
		uid, err := uuid.Parse(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
			return
		}
		q.UserID = &uid
	}

	if v := c.Query("service_name"); v != "" {
		q.ServiceName = &v
	}
	if v := c.Query("start_period"); v != "" {
		q.StartPeriod = &v
	}
	if v := c.Query("end_period"); v != "" {
		q.EndPeriod = &v
	}

	total, err := h.svc.GetTotalCost(c.Request.Context(), q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.CostResponse{TotalCost: total})
}
