package controllers

import (
	"net/http"
	"pfc/service"

	"github.com/gin-gonic/gin"
)

func ValidateIban(context *gin.Context) {
	ibanResponse := service.ValidateIban(context)
	context.JSON(http.StatusOK, ibanResponse)
}
