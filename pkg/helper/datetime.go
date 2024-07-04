package helper

import (
	"fmt"
	"time"
)

func ParseDateTime(dateTime string) time.Time {
	layout := "2006-01-02 15:04:05"

	// Parse the date-time string into a time.Time object in the specified location
	parsedTime, err := time.Parse(layout, dateTime)
	if err != nil {
		fmt.Println("Error parsing date-time with location:", err)
		panic(err)
	}

	return parsedTime
}
