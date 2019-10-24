package data

type Aeroport struct {
	Name string
	Id   string
}
type Datas struct {
	Aeroports []Aeroport
}

type AirportData struct {
	Airportname string
	AirportId   string
	Times       []Time
}

type Time struct {
	Timestamp string
	Time      string
}
