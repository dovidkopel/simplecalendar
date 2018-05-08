package simplecalendar

import (
	"time"
)

type DaySchedule struct {
	SchedulingPolicy
	EnabledSetterGetter

	day     time.Weekday
	enabled bool
}

func EnabledDaySchedule(day time.Weekday) *DaySchedule {
	return &DaySchedule{
		day:     day,
		enabled: true,
	}
}

func DisabledDaySchedule(day time.Weekday) *DaySchedule {
	return &DaySchedule{
		day:     day,
		enabled: false,
	}
}

func (day *DaySchedule) SetEnabled() {
	day.enabled = true
}

func (day *DaySchedule) SetDisabled() {
	day.enabled = false
}

func (day *DaySchedule) HasStart(times EventTimes) bool {
	return times.Start.Weekday() == day.day
}

func (day *DaySchedule) HasEnd(times EventTimes) bool {
	return times.End.Weekday() == day.day
}

func (day *DaySchedule) IsAvailable(times EventTimes) Availability {
	if day.enabled {
		return Available
	} else {
		return Busy
	}
}
