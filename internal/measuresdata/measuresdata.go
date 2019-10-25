package measuresdata

import (
	data "Moniport/internal/data"
	"Moniport/internal/helpers/date"
	"Moniport/internal/helpers/redis"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Connect() {
	redis.Connect()
}

func Disconnect() {
	redis.CloseConnection()
}

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
			Date:        date.GetStringFromDate(date.GetDateFromTimestamp(key)),
		}
		parsedMeasures = append(parsedMeasures, measure)
	}
	return parsedMeasures
}

func SendMeasure(m data.Measure) {
	setKey := m.IDAirport + ":" + m.MeasureType

	redis.AddToSet("airports", m.IDAirport)

	setValue := fmt.Sprintf("%d_%.2f", getNewIdMeasure(), m.Value)

	setTimestamp := date.ParseDate(m.Date)

	fmt.Println(setTimestamp)

	redis.AddToOrdSet(setKey, setValue, date.GetTimestampFromDate(setTimestamp))
}

func getNewIdMeasure() int {
	if redis.KeyExists("currIdMeasure") {
		redis.IncrKey("currIdMeasure")
	} else {
		redis.SendData("currIdMeasure", "0")
	}

	return redis.GetDataInt("currIdMeasure")
}
