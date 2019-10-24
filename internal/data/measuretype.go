package data

type MeasureType string

const (
	TEMPERATURE MeasureType = "temp"
	PRESSURE    MeasureType = "press"
	WIND        MeasureType = "wind"
)

func (mt MeasureType) String() string {
	switch mt {
	case TEMPERATURE:
		return "temp"
	case PRESSURE:
		return "press"
	case WIND:
		return "wind"
	}
	return ""
}
