package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	cities, prices = citiesAndPrices()
	citiesAndPricesMap = groupSlices(cities, prices)
	fmt.Println(citiesAndPricesMap)
}
	


func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 10

	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}

	return cities, prices
}


func groupSlices(keySlice[]string, valueSlice[int]) map[string][]int {

	var mapCities = make(map[string][]int)
	for i, key := range keySlice{
		mapCities[key] = append(mapCities[key], valueSlice[i])
	}
	return mapCities
}

