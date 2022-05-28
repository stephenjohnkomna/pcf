package api

import (
	"pfc/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()
	v1 := route.Group("/v1/api/iban")
	{
		v1.GET("validate/:iban", controllers.ValidateIban)
	}
	return route
}
