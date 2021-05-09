package main

import "time"

type User struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Coins          int     `json:"coins"`
	NumTripsDriven int     `json:"num_trips_driven"`
	NumTripsRidden int     `json:"num_trips_ridden"`
	TotalDistance  float64 `json:"total_distance"`
}

type Drive struct {
	DriveID        string    `json:"drive_id"`
	DriverID       string    `json:"driver_id"`
	Time           time.Time `json:"time"`
	SpaceAvailable int       `json:"space_available"`
	StartAddress   string    `json:"start_address"`
	DestAddress    string    `json:"dest_address"`
	StartLat       float64   `json:"start_lat"`
	StartLng       float64   `json:"start_lng"`
	DestLat        float64   `json:"dest_lat"`
	DestLng        float64   `json:"dest_lng"`
}
type DriveRequest struct {
	DriveID string `json:"drive_id"`
	RiderID string `json:"rider_id"`
	Status  string `json:"status"`
}

type PurchaseRequest struct {
	Item string `json:"item"`
}
