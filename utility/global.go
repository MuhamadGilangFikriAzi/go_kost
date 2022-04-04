package utility

import (
	"fmt"
	"time"
)

func ThisDay() string {
	var dayString string
	var monthString string
	year, month, day := time.Now().Date()
	if day < 10 {
		dayString = fmt.Sprintf("0%d", day)
	} else {
		dayString = fmt.Sprintf("%d", day)
	}

	if int(month) < 10 {
		monthString = fmt.Sprintf("0%d", int(month))
	} else {
		monthString = fmt.Sprintf("%d", int(month))
	}

	stringThisDay := fmt.Sprintf("%v-%s-%s", year, monthString, dayString)
	return stringThisDay
}
