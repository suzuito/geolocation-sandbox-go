package router

import "github.com/suzuito/geolocation-sandbox-go/entity/model"

type location struct {
	ID        string  `json:"id"`
	Seconds   int64   `json:"seconds"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func newLocation(l *model.Location) *location {
	return &location{
		ID:        l.ID,
		Seconds:   l.Seconds,
		Longitude: l.Longitude,
		Latitude:  l.Latitude,
	}
}

// locations ...
type locations struct {
	Partition string      `json:"partition"`
	Locations []*location `json:"locations"`
}

func newLocations(partition string, ls []*model.Location) *locations {
	ret := []*location{}
	for _, l := range ls {
		ret = append(ret, newLocation(l))
	}
	return &locations{
		Partition: partition,
		Locations: ret,
	}
}
