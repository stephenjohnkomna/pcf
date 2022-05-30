package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"pfc/models"
	"pfc/models/dto.go"
)

func ValidateIban(iban string) dto.IbanResponse {
	var ibanResponse dto.IbanResponse
	return ibanResponse
}

func ConvertJsonToStruct(filename string) (models.CountryList, error) {
	// Open the jsonFile
	jsonfile, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Failed to open file ", err)
	}

	// initialize our CountryList
	countries := models.CountryList{}

	// read the opened jsonfile as a byte array.
	err = json.Unmarshal([]byte(jsonfile), &countries)
	if err != nil {
		fmt.Println(err)
	}
	return countries, err
}

func FindCountryByCode(CountryList models.CountryList, countryCode string) models.Country {
	// Iterate through every country within our country list and
	// compare each of the Country Code with the provided Country Code

	var country models.Country

	for i := 0; i < len(CountryList); i++ {
		if CountryList[i].Code == countryCode {
			country = CountryList[i]
			break
		}
	}
	return country
}

func MoveInitialCharacterBehind(iban string) {

}
