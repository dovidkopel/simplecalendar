package simplecalendar

import "time"

type EventTimes struct {
	Start time.Time
	End   time.Time
	Zone  time.Location
}

func times(t time.Time, d time.Duration) EventTimes {
	return EventTimes{
		Start: t,
		End:   t.Add(d),
	}
}

type Event struct {
	Label     string
	Times     EventTimes
	Location  string
	Attendees []string
}

