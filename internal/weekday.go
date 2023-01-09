package internal

import "time"

type Weekday int8
type WeekdaysSum int8

const (
	Monday Weekday = 1 << iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (sum WeekdaysSum) AsWeekdays() []Weekday {
	weekdays := []Weekday{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}

	var retVal []Weekday

	sumValue := sum.asInt()

	for _, weekday := range weekdays {
		day := weekday.asInt()
		if day&sumValue == day {
			retVal = append(retVal, weekday)
		}
	}
	return retVal
}

func (day Weekday) asInt() int8 {
	return int8(day)
}

func (sum WeekdaysSum) asInt() int8 {
	return int8(sum)
}

func TodayWeekday() Weekday {
	return daysMapping[time.Now().Weekday()]
}

var daysMapping = []Weekday{time.Monday: Monday, time.Tuesday: Tuesday,
	time.Wednesday: Wednesday, time.Thursday: Thursday,
	time.Friday: Friday, time.Saturday: Saturday, time.Sunday: Sunday}
