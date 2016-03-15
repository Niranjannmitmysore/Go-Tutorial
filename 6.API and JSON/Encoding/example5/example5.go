
// Sample program to show how to unmarshal a JSON document into
// a user defined struct type from a file.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	// buoyCondition contains information for an individual station.
	WindCondition struct {
		WindSpeed     float64 `json:"wind_speed_milehour"`
		WindDirection int     `json:"wind_direction_degnorth"`
		WindGust      float64 `json:"gust_wind_speed_milehour"`
	}

	// buoyLocation contains the buoys location.
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	// BuoyStation contains information for an individual station.
	RailwayStation struct {
		StationID string        `json:"station_id"`
		Name      string        `json:"name"`
		LocDesc   string        `json:"location_desc"`
		Condition WindCondition `json:"condition"`
		LocationDetails  Location  `json:"location"`
	}
)

// main is the entry point for the application.
func main() {
	// Open the file.
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Open File", err)
		return
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of buoy stations.
	var stations []RailwayStation
	err = json.NewDecoder(file).Decode(&stations)
	if err != nil {
		fmt.Println("Decode File", err)
		return
	}

	// Iterate over the slice and display
	// each station.
	for _, station := range stations {
		fmt.Printf("%+v\n\n", station)
    fmt.Println("description:",station.LocDesc)
	}
}
