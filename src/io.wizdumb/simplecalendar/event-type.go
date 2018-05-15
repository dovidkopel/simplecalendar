package simplecalendar

import (
	"time"
	"strings"
)

type EventType struct {
	Name string
	Duration time.Duration
}

var SyncUp15m = EventType{"SyncUp15m", 15 * time.Minute}
var SyncUp30m = EventType{"SyncUp30m", 30 * time.Minute}
var HourMeeting = EventType{"HourMeeting", 1 * time.Hour}
var TwoHourMeeting = EventType{"TwoHourMeeting", 2 * time.Hour}

func AllEventTypes() []EventType {
	return []EventType{SyncUp15m, SyncUp30m, HourMeeting, TwoHourMeeting}
}

type NoEventTypeError struct{
	error
}
func (e NoEventTypeError) Error() string {
	return "No event type found!"
}

func GetEventType(tpe string) (EventType, error) {
	for _, et := range AllEventTypes() {
		if strings.EqualFold(et.Name, tpe) {
			return et, nil
		}
	}

	return EventType{}, NoEventTypeError{}
}