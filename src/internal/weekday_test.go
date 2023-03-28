package internal

import (
	"reflect"
	"testing"
)

func TestWeekdaysSum_AsWeekdays(t *testing.T) {
	tests := []struct {
		name string
		sum  WeekdaysSum
		want []Weekday
	}{
		{"should return correct weekdays", WeekdaysSum(127), []Weekday{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}},
		{"should return correct weekdays", WeekdaysSum(48), []Weekday{Friday, Saturday}},
		{"should return correct weekdays", WeekdaysSum(37), []Weekday{Monday, Wednesday, Saturday}},
		{"should return correct weekdays", WeekdaysSum(8), []Weekday{Thursday}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sum.AsWeekdays(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsWeekdays() = %v, want %v", got, tt.want)
			}
		})
	}
}
