package data

type Aeroport struct {
	Name string
	Id   string
}
type Datas struct {
	Aeroports []Aeroport
}

const (
	TEMPERATURE = "temp"
	PRESSURE    = "press"
	WIND        = "wind"
)

type AirportData struct {
	Airportname string
	AirportId   string
	Times       []Time
	Types       []string
	Startime    string
	EndTime     string
	Date        string
	Graph1Error bool
	Graph2Error bool
}

type Time struct {
	Timestamp string
	Time      string
}
