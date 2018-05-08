package simplecalendar

import "time"

/**
Repeatable schedules
1. What is the default, should times be marked as "busy" or "available"
  if there is no rule or designation?
2. If no rules are present for the time specified yield to default.
3. Look at day blocks for enabled/disabled
4. Look at time blocks for enabled/disabled
*/

// Is available?
type DefaultApproach bool
var DefaultBusy = DefaultApproach(false)
var DefaultAvailable = DefaultApproach(true)

type Availability bool
type Available Availability
type Busy Availability
type Unknown Availability

type EnabledSetterGetter interface {
	SetEnabled()
	SetDisabled()
}

type EnableDisableDays interface {
	SetEnabled(days ...time.Weekday)
	SetDisabled(days ...time.Weekday)
	UnsetDays(days ...time.Weekday)
}

// Ultimate question needed to be answered
type SchedulingPolicy interface {
	HasStart(times EventTimes) bool
	HasEnd(times EventTimes) bool
	IsAvailable(times EventTimes) bool
}

type GeneralSchedulingApproachConfig struct {
	approach DefaultApproach
}

type SchedulingConf struct {
	general GeneralSchedulingApproachConfig

}

var conf = SchedulingConf{
	general: GeneralSchedulingApproachConfig{
		approach: DefaultAvailable,
	},
}

type DaySchedule struct {
	SchedulingPolicy
	EnabledSetterGetter

	day time.Weekday
	enabled bool
}

func EnabledDaySchedule(day time.Weekday) DaySchedule {
	return DaySchedule{
		day: day,
		enabled: true,
	}
}

func DisabledDaySchedule(day time.Weekday) DaySchedule {
	return DaySchedule{
		day: day,
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

func (day *DaySchedule) IsAvailable(times EventTimes) bool {
	return day.enabled
}

type WeekSchedule struct {
	SchedulingPolicy
	EnableDisableDays

	days map[time.Weekday]DaySchedule
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

func (week *WeekSchedule) IsAvailable(times EventTimes) bool {
	var start bool
	var end bool

	// There is a start value
	if week.HasStart(times) {
		// Is it enabled or disabled
		start = week.days[times.Start.Weekday()].IsAvailable(times)
	} else {
		// Use default
		start = bool(conf.general.approach)
	}

	// There is a end value
	if week.HasEnd(times) {
		// Is it enabled or disabled
		end = week.days[times.End.Weekday()].IsAvailable(times)
	} else {
		// Use default
		end = bool(conf.general.approach)
	}

	return start && end
}

var DefaultBusinessWeek = WeekSchedule{
	days: map[time.Weekday]DaySchedule {
		time.Sunday: DisabledDaySchedule(time.Sunday),
		time.Monday: EnabledDaySchedule(time.Monday),
		time.Tuesday: EnabledDaySchedule(time.Tuesday),
		time.Wednesday: EnabledDaySchedule(time.Wednesday),
		time.Thursday: EnabledDaySchedule(time.Thursday),
		time.Friday: EnabledDaySchedule(time.Friday),
		time.Saturday: EnabledDaySchedule(time.Saturday),
	},
}

func test() {

}
