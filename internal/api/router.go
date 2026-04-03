package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	paymentHandler "github.com/max1t1a/subscription-service/internal/api/handler/payment"
	subscriptionHandler "github.com/max1t1a/subscription-service/internal/api/handler/subscription"
	"github.com/max1t1a/subscription-service/internal/api/middleware"
)

func NewRouter(subH *subscriptionHandler.Handler, payH *paymentHandler.Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(middleware.Logger(), gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		subs := v1.Group("/subscriptions")
		{
			subs.POST("", subH.Create)
			subs.GET("", subH.List)
			subs.GET("/cost", subH.GetTotalCost)
			subs.GET("/:id", subH.GetByID)
			subs.PUT("/:id", subH.Update)
			subs.DELETE("/:id", subH.Delete)
		}

		v1.GET("/payments", payH.ListBySubscription)
	}

	return r
}
