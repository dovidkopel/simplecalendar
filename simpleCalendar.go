/**
 * @license
 * Copyright Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START calendar_quickstart]
package main

import (
	"fmt"
	"log"
	"time"

	"io.wizdumb/simplecalendar"
)

func main() {
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
		[]string {"dovid+test1@dovidkopel.com", "dovid+test2@dovidkopel.com", "dovid+test3@dovidkopel.com" },
	)
	oe := event.Insert()
	log.Printf("Event created: %s\n", oe.HtmlLink)

	//
	//
	//// Insert
	//event := &calendar.Event{
	//	Summary:     "Google I/O 2020",
	//	Location:    "800 Howard St., San Francisco, CA 94103",
	//	Description: "A chance to hear more about Google's developer products.",
	//	Start: &calendar.EventDateTime{
	//		DateTime: "2018-05-28T09:00:00-07:00",
	//		TimeZone: "America/Los_Angeles",
	//	},
	//	End: &calendar.EventDateTime{
	//		DateTime: "2018-05-28T17:00:00-07:00",
	//		TimeZone: "America/Los_Angeles",
	//	},
	//	Attendees: []*calendar.EventAttendee{},
	//}
	//
	//calendarId := "dovidkopel@gmail.com"
	//event, err = srv.Events.Insert(calendarId, event).Do()
	//if err != nil {
	//	log.Fatalf("Unable to create event. %v\n", err)
	//}
	//fmt.Printf("Event created: %s\n", event.HtmlLink)

	//t := time.Now().Format(time.RFC3339)
	//events, err := srv.Events.List("primary").ShowDeleted(false).
	//	SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	//if err != nil {
	//	log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	//}
	//fmt.Println("Upcoming events:")
	//if len(events.Items) == 0 {
	//	fmt.Println("No upcoming events found.")
	//} else {
	//	for _, item := range events.Items {
	//		date := item.Start.DateTime
	//		if date == "" {
	//			date = item.Start.Date
	//		}
	//
	//		endDate := item.End.DateTime
	//
	//		fmt.Printf("%v (%v -> %v)\n", item.Summary, date, endDate)
	//	}
	//}
}

// [END calendar_quickstart]
