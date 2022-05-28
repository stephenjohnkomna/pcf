package service

import (
	"pfc/common"
	models "pfc/models/responsemodel.go"

	"github.com/gin-gonic/gin"
)

func ValidateIban(context *gin.Context) models.IbanResponse {
	var ibanResponse models.IbanResponse
	iban := context.Params.ByName("iban")
	ibanResponse = common.ValidateIban(iban)
	ibanResponse.IsValid = true
	ibanResponse.Message = "Validation Successful"
	return ibanResponse
}
