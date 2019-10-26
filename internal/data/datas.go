package data

type Aeroport struct {
	Name string
	Id   string
}
type Datas struct {
	Aeroports []Aeroport
}

type AirportData struct {
	Airportname  string
	AirportId    string
	Types        []MeasureType
	Startime     string
	EndTime      string
	GraphData    []int
	GraphDates   []string
	SelectedType string
	Date         string
	Graph1Error  bool
	Graph2Error  bool
	Moyennes     [3]int
}

type Time struct {
	Timestamp string
	Time      string
}
