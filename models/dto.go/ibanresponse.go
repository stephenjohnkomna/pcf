package dto

type IbanResponse struct {
	IsValid bool   `json:"isvalid"`
	Message string `json:"message"`
}
