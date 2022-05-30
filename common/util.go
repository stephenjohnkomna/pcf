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

func LookUpTable(letter string) int {
	m := map[string]int{
		"A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15,
		"G": 16, "H": 17, "I": 18, "J": 19, "K": 20, "L": 21,
		"M": 22, "N": 23, "O": 24, "P": 25, "Q": 26, "R": 27,
		"S": 28, "T": 29, "U": 30, "V": 31, "W": 32, "X": 33,
		"Y": 34, "Z": 35,
	}
	return m[letter]
}
