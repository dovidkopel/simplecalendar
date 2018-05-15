package simplecalendar

import (
	"log"
	"time"

	"fmt"
	"google.golang.org/api/calendar/v3"
	googleapi "google.golang.org/api/googleapi"
)

var srv *calendar.Service
var calendarId string
var cals []string

func CalendarInit() *calendar.Service {
	log.Printf("Initializing calendar service")
	s, err := getService()

	if err != nil {
		log.Fatalf("Unable to init: %v", err)
	}
	srv = s

	ids, err := GetCalendars()
	if err != nil {
		fmt.Println(ids)
	} else {
		for _, v := range ids {
			cals = append(cals, v)
		}
		log.Printf("There are %d calendars", len(cals))
	}

	return srv
}

func GetCalendar() *calendar.Service {
	return srv
}

func convertAttendees(attendees []string) []*calendar.EventAttendee {
	var as []*calendar.EventAttendee
	for _, v := range attendees {
		log.Printf("Attendee email %s", v)
		as = append(as, &calendar.EventAttendee{
			Email: v,
		})
	}
	log.Printf("There are %d attendees", len(as))

	return as
}

type SendNotificationsOption struct{}

func (q SendNotificationsOption) Get() (string, string) {
	return "sendNotifications", "true"
}

func sendNotifications() googleapi.CallOption {
	return SendNotificationsOption{}
}

/**
Insert a an event created into the user's Google Calendar
*/
func (e *Event) Insert() *calendar.Event {
	ex := &calendar.Event{
		Summary:  e.Label,
		Location: e.Location,
		Start: &calendar.EventDateTime{
			DateTime: e.Times.Start.Format(time.RFC3339),
			TimeZone: e.Times.Zone.String(),
		},
		End: &calendar.EventDateTime{
			DateTime: e.Times.End.Format(time.RFC3339),
			TimeZone: e.Times.Zone.String(),
		},
		Attendees: convertAttendees(e.Attendees),
	}
	event, err := srv.Events.Insert("primary", ex).Do(
		sendNotifications(),
	)

	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	return event
}

func GetCalendars() ([]string, error) {
	var l []string

	res, err := srv.CalendarList.List().Fields("items/id").Do()

	if err != nil {
		log.Printf("Unable to retrieve list of calendars: %v", err)
		return nil, err
	}

	for _, v := range res.Items {
		var id string
		id = v.Id
		l = append(l, id)
	}

	return l, err
}

func GetEvents(min time.Time, max time.Time) []Event {
	events, err := GetCalendar().Events.List("primary").
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(min.Format(time.RFC3339)).
		TimeMax(max.Format(time.RFC3339)).
		OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
		return []Event{}
	} else {
		var as []Event
		for _, item := range events.Items {
			if len(item.Start.DateTime) > 0 && len(item.End.DateTime) > 0 {
				date, _ := time.Parse(time.RFC3339, item.Start.DateTime)
				endDate, _ := time.Parse(time.RFC3339, item.End.DateTime)
				ts := EventTimes{
					Start: date,
					End: endDate,
				}
				as = append(as, Event{
					Label: item.Summary,
					Times: ts,
				})
			}
		}
		return as
	}
}

func CreateEventTime(start time.Time, duration time.Duration) EventTimes {
	return EventTimes{
		Start: start,
		End:   start.Add(duration),
	}
}

func CreateEvent(label string, times EventTimes, location string, attendees []string) Event {
	return Event{
		Label:     label,
		Times:     times,
		Location:  location,
		Attendees: attendees,
	}
}
