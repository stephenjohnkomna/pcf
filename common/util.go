package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"pfc/models"
	"pfc/models/dto.go"
	"strconv"
)

func GetListOfCountries(filename string) (models.CountryList, error) {
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

// Function to Look Letter for its corresponding value representation
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

// This function converts the Rearrange IBAN to Integer
func ConvertToInteger(rearrangeIban string) string {

	rerrangedIbanToIntegerValue := ""
	for _, value := range rearrangeIban {

		_, err := strconv.Atoi(string(value))
		if err != nil {
			toInt := LookUpTable(string(value))
			rerrangedIbanToIntegerValue = rerrangedIbanToIntegerValue + strconv.Itoa(toInt)
		} else {
			rerrangedIbanToIntegerValue = rerrangedIbanToIntegerValue + string(value)
		}

	}
	return rerrangedIbanToIntegerValue
}

// Validate that the Country Code is Valid
func ValidateCountryCode(code string) bool {
	countries, _ := GetListOfCountries("country.json")
	country := FindCountryByCode(countries, code)
	if country.Code == code {
		return true
	} else {
		return false
	}
}

func ValidateIban(iban string) dto.IbanResponse {
	// Initialize IbanResponse
	var ibanResponse dto.IbanResponse
	ibanResponse.IsValid = false
	ibanResponse.Message = "IBAN is NOT VALID"

	// • IBAN:		GB82 WEST 1234 5698 7654 32
	// Validate Country Code
	countryCode := iban[0:2]
	if ValidateCountryCode(countryCode) != true {
		return ibanResponse
	}

	// • Rearrange:		W E S T12345698765432 G B82
	// First 2 characters are Country Code and the Next 2 characters are Check digit to be moved to the end of the IBAN
	rearrangedIban := iban[4:len(iban)] + iban[0:4]

	// • Convert to integer:		3214282912345698765432161182
	integerIbanConversion := ConvertToInteger(rearrangedIban)

	// • Compute remainder:		3214282912345698765432161182	mod 97 = 1
	isValid := strconv.ParseInt(integerIbanConversion, 10, 64) % 97

	if isValid == true {
		ibanResponse.IsValid = true
		ibanResponse.Message = "IBAN is VALID"
	}
	return ibanResponse
}
