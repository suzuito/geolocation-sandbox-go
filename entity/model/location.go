package model

// Location ...
type Location struct {
	ID        string  `json:"id"`
	Seconds   int64   `json:"seconds"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
