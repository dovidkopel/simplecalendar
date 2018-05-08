package simplecalendar

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

func (s *SchedulingConf) SetTimes(t TimeSchedule) {
	s.times = t
}

var Conf = SchedulingConf{
	general: GeneralSchedulingApproachConfig{
		approach: DefaultAvailable,
	},
	week: DefaultBusinessWeek,
}
