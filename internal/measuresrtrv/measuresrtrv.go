package measuresrtrv

import (
	data "Moniport/internal/data"
	redis "Moniport/internal/helpers/redis"
	"strconv"
	"strings"
	"time"
)

func GetAirports() []string {
	return redis.GetSet("airports")
}

func GetMeasures(airport string, measureType string) []data.Measure {
	measures := redis.GetAllFromOrderedSet(airport + ":" + measureType)
	return parseMeasures(airport, measureType, measures)
}

func GetMeasuresInRange(airport string, measureType string, start, end time.Time) []data.Measure {
	measures := redis.GetRangeFromOrderedSet(airport+":"+measureType, start.Unix(), end.Unix())
	return parseMeasures(airport, measureType, measures)
}

func parseMeasures(airport string, measureType string, measures map[int64]string) []data.Measure {
	var parsedMeasures []data.Measure
	for key, value := range measures {
		strValue := strings.Split(value, "_")[1]
		value, _ := strconv.ParseFloat(strValue, 64)
		measure := data.Measure{
			IDAirport:   airport,
			MeasureType: measureType,
			Value:       value,
			Date:        data.MeasureDateFromTimestamp(key),
		}
		parsedMeasures = append(parsedMeasures, measure)
	}
	return parsedMeasures
}
