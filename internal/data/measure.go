package data

type Measure struct {
	IDSensor    string  `json:"idSensor"`
	IDAirport   string  `json:"idAirport"`
	MeasureType string  `json:"measure"`
	Value       float64 `json:"value"`
	Date        string  `json:"date"`
}
