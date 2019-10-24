package data

type Configuration struct {
	AdressBroker string `json:"adressBroker"`
	PortBroker   int    `json:"portBroker"`
	LevelQos     byte   `json:"levelQos"`
	IDSensor     int    `json:"idSensor"`
	IDAirport    string `json:"idAirport"`
}
