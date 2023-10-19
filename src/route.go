package route

import (
	"time"
)

type OpeningHours struct {
	Day       string
	StartTime string
	EndTime   string
}

func NewOpeningHours(day string, startTime string, endTime string) *OpeningHours {
	return &OpeningHours{
		Day:       day,
		StartTime: startTime,
		EndTime:   endTime,
	}
}

// Definicion de la entidad Area de descanso
type RestArea struct {
	openingHours []OpeningHours
}

func (r RestArea) OpeningHours() []OpeningHours {
	return r.openingHours
}

func NewRestArea(openingHours []OpeningHours) *RestArea {
	return &RestArea{
		openingHours: openingHours,
	}
}

// Definicion de las Distancias entre areas de descanso
type RestAreaDistances struct {
	distances [][]uint16
}

func (r RestAreaDistances) Distance(origin uint8, destination uint8) uint16 {
	return r.distances[origin][destination]
}

func NewRestAreaDistances(distances [][]uint16) *RestAreaDistances {
	return &RestAreaDistances{
		distances: distances,
	}
}

// Definición de la entidad Ruta
type Route struct {
	id               string
	arrivalTime      time.Time
	restAreas        []RestArea
	timesToRestAreas []time.Duration
}

func (r Route) Id() string {
	return r.id
}

func (r Route) ArrivalTime() time.Time {
	return r.arrivalTime
}

func (r Route) RestAreas() []RestArea {
	return r.restAreas
}

func (r Route) TimesToRestAreas() []time.Duration {
	return r.timesToRestAreas
}

func NewRoute(id string, arrivalTime time.Time, restAreas []RestArea, timesToRestAreas []time.Duration) *Route {
	return &Route{
		id:               id,
		arrivalTime:      arrivalTime,
		restAreas:        restAreas,
		timesToRestAreas: timesToRestAreas,
	}
}
