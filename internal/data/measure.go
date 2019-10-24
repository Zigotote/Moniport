package data

import "time"

type Measure struct {
	IDSensor    string  `json:"idSensor"`
	IDAirport   string  `json:"idAirport"`
	MeasureType string  `json:"measure"`
	Value       float64 `json:"value"`
	Date        string  `json:"date"`
}

//a bouger
func MeasureDateFromTimestamp(date int64) string {
	layout := "2006-01-02-15-04-05"

	return time.Unix(date, 0).Format(layout)
}
