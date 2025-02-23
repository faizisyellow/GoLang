package main

import (
	"fmt"
)

type City struct {
	country string
	capital string
	region  string
}

func New(country, city, region string) *City {
	return &City{
		country: country,
		capital: city,
		region:  region,
	}
}

func main() {
	cities := []City{
		*New("Usa", "Washington DC.", "america"),
		*New("Japan", "Tokyo", "asian"),
		*New("Indonesia", "Jakarta", "asian"),
		*New("Netherlands", "Amsterdam", "europe"),
	}

	citiesByCountry := make(map[string]City, 0)
	createNewCities(&citiesByCountry, &cities)

	citiesByRegionCountry := filter(&citiesByCountry, func(region string) bool {
		return region == "asian"
	})

	citiesCompare := filter(&citiesByCountry, func(region string) bool {
		return region == "Jakarta"
	})

	// ngopibang :=  (*citiesCompare)["Indonesia"]
	// ngopibang.capital = "Bali"

	// fmt.Println(ngopibang)
	(*citiesCompare)["Indonesia"] = City{
		capital: "Bali",
		country: (*citiesCompare)["Indonesia"].country,
		region:  (*citiesCompare)["Indonesia"].region,
	}

	fmt.Println(*citiesCompare)
	fmt.Println(*citiesByRegionCountry)

}

func createNewCities(citiesM *map[string]City, cities *[]City) {
	for _, v := range *cities {
		(*citiesM)[v.country] = v
	}
}

func filter(mapCitites *map[string]City, filterfn func(region string) bool) *map[string]City {
	newMapCities := make(map[string]City, len(*mapCitites))

	for _, v := range *mapCitites {
		// Calling function parameter
		if filterfn(v.region) {
			newMapCities[v.country] = v
		} else if filterfn(v.capital) {
			newMapCities[v.country] = v
		}
	}

	if len(newMapCities) == 0 {
		fmt.Println("City not found, try another region.")
		return &map[string]City{}
	}

	return &newMapCities
}
