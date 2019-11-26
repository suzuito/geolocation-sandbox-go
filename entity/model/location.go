package model

// Location ...
type Location struct {
	Seconds   int64   `json:"seconds"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
