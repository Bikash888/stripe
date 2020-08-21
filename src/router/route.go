package router

import (
	"github.com/gin-gonic/gin"
	"stripe.com/s/src/controller"
)

func ChargeRouter() *gin.Engine {
	r := gin.Default()
	g1 := r.Group("/stripe")
	{
		g1.POST("payment", controller.Payment)
	}
	return r
}
