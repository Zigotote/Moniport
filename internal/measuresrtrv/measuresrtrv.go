package measuresrtrv

import (
	redis "Moniport/internal/helpers/redis"
	"strconv"
	"strings"
	"time"
)

func GetAirports() []string {
	return redis.GetSet("airports")
}

func GetMeasures(airport string, measureType string) map[int64]float64 {
	return parseMeasures(redis.GetAllFromOrderedSet(airport + ":" + measureType))
}

func GetMeasuresInRange(airport string, measureType string, start, end time.Time) map[int64]float64 {
	return parseMeasures(redis.GetRangeFromOrderedSet(airport+":"+measureType, start.Unix(), end.Unix()))
}

func parseMeasures(measures map[int64]string) map[int64]float64 {
	parsedMeasures := make(map[int64]float64)
	for key, value := range measures {
		strValue := strings.Split(value, "_")[1]
		parsedMeasures[key], _ = strconv.ParseFloat(strValue, 64)
	}
	return parsedMeasures
}
