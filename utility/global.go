package utility

import (
	"fmt"
	"time"
)

func AddZeroAndToString(number int) string {
	if number < 10 {
		return fmt.Sprintf("0%d", number)
	}
	return fmt.Sprintf("%d", number)

}

func ThisDay() string {
	year, month, day := time.Now().Date()

	stringThisDay := fmt.Sprintf("%v-%s-%s", year, AddZeroAndToString(int(month)), AddZeroAndToString(day))
	return stringThisDay
}
