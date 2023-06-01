package pumps

import "time"

type response struct {
	Type string `json:"type"`
	Crs  struct {
		Type       string `json:"type"`
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
	} `json:"crs"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			StationID   string    `json:"station_id"`
			StationName string    `json:"station_name"`
			FuelID      string    `json:"fuel_id"`
			Price       string    `json:"price"`
			CountryID   int       `json:"country_id"`
			Datetime    time.Time `json:"datetime"`
		} `json:"properties,omitempty"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

type Station struct {
	ID            string
	Name          string
	Brand         string
	Address       string
	DieselPrice   string
	PetrolPrice   string
	ElectricPrice string
}

type Features []Feature

type Feature struct{}

type Property struct {
	StationID   string `json:"station_id"`
	StationName string `json:"station_name"`
	FuelID      string `json:"fuel_id"`
	Price       string `json:"price"`
	CountryID   string `json:"country_id"`
	Datetime    string `json:"datetime"`
}
