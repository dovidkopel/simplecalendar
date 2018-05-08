package simplecalendar

import (
	"log"
	"time"
)

type EnableDisableDays interface {
	SetEnabled(days ...time.Weekday)
	SetDisabled(days ...time.Weekday)
	UnsetDays(days ...time.Weekday)
}

type WeekSchedule struct {
	SchedulingPolicy
	EnableDisableDays

	days map[time.Weekday]*DaySchedule
}

func (week *WeekSchedule) SetEnabled(days ...time.Weekday) {
	for _, day := range days {
		week.days[day].SetEnabled()
	}
}

func (week *WeekSchedule) SetDisabled(days ...time.Weekday) {
	for _, day := range days {
		week.days[day].SetEnabled()
	}
}

func (week *WeekSchedule) UnsetDays(days ...time.Weekday) {
	for _, day := range days {
		delete(week.days, day)
	}
}

func (week *WeekSchedule) HasStart(times EventTimes) bool {
	_, exists := week.days[times.Start.Weekday()]
	return exists
}

func (week *WeekSchedule) HasEnd(times EventTimes) bool {
	_, exists := week.days[times.End.Weekday()]
	return exists
}

func (week *WeekSchedule) IsAvailable(times EventTimes) Availability {
	var start Availability
	var end Availability

	// There is a start value
	if week.HasStart(times) {
		log.Printf("%s is in the start time", times.Start.Weekday())
		// Is it enabled or disabled
		start = week.days[times.Start.Weekday()].IsAvailable(times)
	} else {
		start = Unknown
	}

	// There is a end value
	if week.HasEnd(times) {
		log.Printf("%s is in the end time", times.End.Weekday())
		// Is it enabled or disabled
		end = week.days[times.End.Weekday()].IsAvailable(times)
	} else {
		end = Unknown
	}

	return simplify([]Availability{start, end})
}

var DefaultBusinessWeek = WeekSchedule{
	days: map[time.Weekday]*DaySchedule{
		time.Sunday:    DisabledDaySchedule(time.Sunday),
		time.Monday:    EnabledDaySchedule(time.Monday),
		time.Tuesday:   EnabledDaySchedule(time.Tuesday),
		time.Wednesday: EnabledDaySchedule(time.Wednesday),
		time.Thursday:  EnabledDaySchedule(time.Thursday),
		time.Friday:    EnabledDaySchedule(time.Friday),
		time.Saturday:  EnabledDaySchedule(time.Saturday),
	},
}
