package service

import (
	"pfc/common"
	"pfc/models/dto.go"

	"github.com/gin-gonic/gin"
)

func ValidateIban(context *gin.Context) dto.IbanResponse {
	var ibanResponse dto.IbanResponse
	iban := context.Params.ByName("iban")
	ibanResponse = common.ValidateIban(iban)
	ibanResponse.IsValid = true
	ibanResponse.Message = "Validation Successful"
	return ibanResponse
}
