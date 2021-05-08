package main

import "time"

type User struct {
    ID             string `json:"id"`
    Name           string `json:"name"`
    Coins          int    `json:"coins"`
    NumTripsDriven int    `json:"num_trips_driven"`
    NumTripsRidden int    `json:"num_trips_ridden"`
    TotalDistance  float64 
`json:"total_distance"`
}

type Drive struct {
    DriverID       string  `json:"driver_id"`
    Time           string  `json:"time"`
    DriveID        int      `json:"drive_id"`
    SpaceAvailable int     `json:"space_available"`
    StartAddress   string  `json:"start_address"`
    EndAddress     string  `json:"end_address"`
    StartLat       float64 `json:"start_lat"`
    StartLng       float64 `json:"start_lng"`
    DestLat        float64 `json:"dest_lat"`
    DestLng        float64 `json:"dest_lng"`
}
type DriveRequest struct {
    DriveID  int `json:"drive_id"`
    RiderID  string `json:"rider_id"`
    Status   string `json:"status"`
}
