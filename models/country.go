package models

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CountryList []struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
