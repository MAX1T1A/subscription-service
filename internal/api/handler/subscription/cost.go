package subscription

import (
	"net/http"
	"time"

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
// @Param        start_period query string true  "Period start (MM-YYYY)"
// @Param        end_period   query string true  "Period end (MM-YYYY)"
// @Success      200 {object} model.CostResponse
// @Failure      400 {object} ErrorResponse
// @Router       /subscriptions/cost [get]
func (h *Handler) GetTotalCost(c *gin.Context) {
	startStr := c.Query("start_period")
	endStr := c.Query("end_period")

	if startStr == "" || endStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_period and end_period are required"})
		return
	}

	startPeriod, err := time.Parse("01-2006", startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_period format, expected MM-YYYY"})
		return
	}

	endPeriod, err := time.Parse("01-2006", endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_period format, expected MM-YYYY"})
		return
	}

	// make end_period exclusive (first of next month)
	endPeriod = endPeriod.AddDate(0, 1, 0)

	if !endPeriod.After(startPeriod) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end_period must be after start_period"})
		return
	}

	q := model.CostQuery{
		StartPeriod: startPeriod,
		EndPeriod:   endPeriod,
	}

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

	total, err := h.svc.GetTotalCost(c.Request.Context(), q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.CostResponse{TotalCost: total})
}
