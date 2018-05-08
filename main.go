package main

import (
	"fmt"
	"log"
	"time"

	"context"
	"io.wizdumb/simplecalendar"
)

/**
Based on configuration enable remote user to select
events and they will be added to the user's calendar

--Event Types--
Display the event types that are available

--Free/Busy--
- What is the general daily schedule?

--Schedule Event--
- Given a selected EventType display the next available time slots
- If a date is specified begin the time slots at the date specified
- Else, find earliest time
*/

func testTime() {

}

func testDate() {
	// saturday
	st, _ := time.Parse(time.RFC3339, "2018-05-05T15:02:20+00:00")
	// sunday
	ed, _ := time.Parse(time.RFC3339, "2018-05-06T15:02:20+00:00")

	isAvail := simplecalendar.DefaultBusinessWeek.IsAvailable(simplecalendar.EventTimes{
		Start: st,
		End:   ed,
	})
	if isAvail == simplecalendar.Available {
		log.Println("Is available!")
	} else {
		log.Println("Is NOT available!")
	}

	// saturday
	st1, _ := time.Parse(time.RFC3339, "2018-05-02T15:02:20+00:00")
	// sunday
	ed1, _ := time.Parse(time.RFC3339, "2018-05-03T15:02:20+00:00")

	isAvail1 := simplecalendar.DefaultBusinessWeek.IsAvailable(simplecalendar.EventTimes{
		Start: st1,
		End:   ed1,
	})
	if isAvail1 == simplecalendar.Available {
		log.Println("Is available!")
	} else {
		log.Println("Is NOT available!")
	}
}

func testEventInsert() {
	simplecalendar.CalendarInit()
	// Event spans from monday - tuesday
	st, err := time.Parse(time.RFC3339, "2018-05-21T12:00:00-04:00")
	if err != nil {
		log.Fatalf("Unable to init: %v", err)
	}
	log.Printf("Starting: %s\n", st.String())
	d := (time.Duration(5) * time.Hour) + (time.Duration(30) * time.Minute)
	fmt.Println(d.String())
	et := simplecalendar.CreateEventTime(st, d)
	log.Printf("Ending: %s\n", et.End.String())
	event := simplecalendar.CreateEvent(
		"my test",
		et,
		"6605 Deancroft Road Baltimore, MD 21209",
		[]string{"dovid+test1@dovidkopel.com", "dovid+test2@dovidkopel.com", "dovid+test3@dovidkopel.com"},
	)
	oe := event.Insert()
	log.Printf("Event created: %s\n", oe.HtmlLink)
}

func testGetEvents() {
	t := time.Now().Format(time.RFC3339)
	events, err := simplecalendar.GetCalendar().Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}

			endDate := item.End.DateTime

			fmt.Printf("%v (%v -> %v)\n", item.Summary, date, endDate)
		}
	}
}

type MyEvent struct {
	Name string `json:"name"`
}

func test() {
	testTime()
	testDate()
	//testEventInsert()
	//testGetEvents()
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	test()

	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	test()
	//lambda.Start(HandleRequest)
}
