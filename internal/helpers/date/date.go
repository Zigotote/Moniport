package date

import (
	"fmt"
	"time"
)

func ParseDate(date string) time.Time {
	layout := "2006-01-02-15-04-05"

	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return t
}

func GetTimestampFromDate(date time.Time) int64 {
	return date.Unix()
}

/*func GetTimestampFromDay(date string) int64 {
	layout := "2006-01-02"

	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return t.Unix()
}*/
