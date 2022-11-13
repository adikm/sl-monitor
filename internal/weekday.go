package internal

type Weekday int8
type WeekdaysSum int8

const (
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 4
	Thursday  Weekday = 8
	Friday    Weekday = 16
	Saturday  Weekday = 32
	Sunday    Weekday = 64
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
