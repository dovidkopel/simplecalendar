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
type Availability struct {
	string
	name string
}

var Available = Availability{
	name: "available",
}
var Busy = Availability{
	name: "busy",
}
var Unknown = Availability{
	name: "unknown",
}

func (a Availability) String() string {
	return a.name
}

func simplify(as []Availability) Availability {
	for _, a := range as {
		if a == Busy {
			return Busy
		} else if a == Unknown {
			if Availability(Conf.general.approach) == Busy {
				return Busy
			}
		}
	}
	return Available
}

func isAvailable(as []Availability) bool {
	if simplify(as) == Busy {
		return false
	}
	return true
}

type DefaultApproach Availability

var DefaultBusy = DefaultApproach(Busy)
var DefaultAvailable = DefaultApproach(Available)

type DefaultApproachSetter interface {
	SetDefaultIsAvailable()
	SetDefaultIsBusy()
}

type EnabledSetterGetter interface {
	SetEnabled()
	SetDisabled()
}

// Ultimate question needed to be answered
type SchedulingPolicy interface {
	IsAvailable(times EventTimes) Availability
}

type AvailabilityIterator interface {
	Next() EventTimes
}

type AvailabilityQuery struct {
	AvailabilityIterator

	min time.Time
	dur time.Duration

	max time.Time
}

func createIterator(min time.Time, dur time.Duration) AvailabilityQuery {
	return AvailabilityQuery{min: min, dur: dur, max: min.AddDate(0, 0, 7)}
}

func (q *AvailabilityQuery) Next() EventTimes {
	// Get any events within that time
	es := GetEvents(q.min, q.max)
	if len(es) > 0 {

	}
	return EventTimes{}
}

type GeneralSchedulingApproachConfig struct {
	DefaultApproachSetter

	approach DefaultApproach
}

func (c *GeneralSchedulingApproachConfig) SetDefaultIsAvailable() {
	c.approach = DefaultAvailable
}

func (c *GeneralSchedulingApproachConfig) SetDefaultIsBusy() {
	c.approach = DefaultBusy
}

type SchedulingConf struct {
	general GeneralSchedulingApproachConfig
	week    WeekSchedule
	times   TimeSchedule
}

func (s *SchedulingConf) SetWeek(w WeekSchedule) {
	s.week = w
}

func (s *SchedulingConf) SetTimes(t TimeSchedule) {
	s.times = t
}

var Conf = SchedulingConf{
	general: GeneralSchedulingApproachConfig{
		approach: DefaultAvailable,
	},
	week: DefaultBusinessWeek,
}

// Func to approve

// First the week
// Then times

func FindTimes(tpe EventType) []EventTimes {
	// Find how many slots in the next 7 days
	// Only let meetings be scheduled at 15 minute intervals
	// On a given calendar day first determine what are all of the valid
	// slots

}