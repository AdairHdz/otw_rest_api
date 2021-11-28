package utility

import "time"

func ServerWeekDay() int {
	t := time.Now()
	weekday := t.Weekday()

	if int(weekday) == 0 {
		return 7
	}

	return int(weekday)
}