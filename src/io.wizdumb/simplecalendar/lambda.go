package simplecalendar

import (
	"fmt"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
	"log"
	"encoding/json"
	"encoding/base64"
)

type MyEvent struct {
	Name string `json:"name"`
}

func Main() {
	//simplecalendar.test()


	lambda.Start(handleRequest)
}

func testTime() {

}

func testDate() {
	// saturday
	st, _ := time.Parse(time.RFC3339, "2018-05-05T15:02:20+00:00")
	// sunday
	ed, _ := time.Parse(time.RFC3339, "2018-05-06T15:02:20+00:00")

	isAvail := DefaultBusinessWeek.IsAvailable(EventTimes{
		Start: st,
		End:   ed,
	})
	if isAvail == Available {
		log.Println("Is available!")
	} else {
		log.Println("Is NOT available!")
	}

	// saturday
	st1, _ := time.Parse(time.RFC3339, "2018-05-02T15:02:20+00:00")
	// sunday
	ed1, _ := time.Parse(time.RFC3339, "2018-05-03T15:02:20+00:00")

	isAvail1 := DefaultBusinessWeek.IsAvailable(EventTimes{
		Start: st1,
		End:   ed1,
	})
	if isAvail1 == Available {
		log.Println("Is available!")
	} else {
		log.Println("Is NOT available!")
	}
}

func testEventInsert() {
	CalendarInit()
	// Event spans from monday - tuesday
	st, err := time.Parse(time.RFC3339, "2018-05-21T12:00:00-04:00")
	if err != nil {
		log.Fatalf("Unable to init: %v", err)
	}
	log.Printf("Starting: %s\n", st.String())
	d := (time.Duration(5) * time.Hour) + (time.Duration(30) * time.Minute)
	fmt.Println(d.String())
	et := CreateEventTime(st, d)
	log.Printf("Ending: %s\n", et.End.String())
	event := CreateEvent(
		"my test",
		et,
		"6605 Deancroft Road Baltimore, MD 21209",
		[]string{"dovid+test1@dovidkopel.com", "dovid+test2@dovidkopel.com", "dovid+test3@dovidkopel.com"},
	)

	log.Printf("initial: %s", event)

	o, _ := json.Marshal(event)
	log.Printf("json: %s", o)
	o1 := base64.StdEncoding.EncodeToString(o)
	log.Printf("base64: %s", o1)
	b1, _ := base64.StdEncoding.DecodeString(o1)
	var o2 Event
	_ = json.Unmarshal(b1, &o2)

	log.Printf("unmarshalled: %s", o2)


	//oe := event.Insert()
	//log.Printf("Event created: %s\n", oe.HtmlLink)
}

func test() {
	CalendarInit()
	s := GetEvents(time.Now(), time.Now().AddDate(0, 0, 7))
	for _, ss := range s {
		log.Println(ss)
	}
	//testTime()
	//testDate()
	//testEventInsert()
	//testGetEvents()
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get event types
	if request.Path == "/" && request.HTTPMethod == "GET" {
		r, _ := json.Marshal(AllEventTypes())
		return events.APIGatewayProxyResponse{Body: string(r), StatusCode: 200}, nil
	} else if request.Resource == "/slots/{eventType}" && request.HTTPMethod == "GET" {
	// Event name / duration
		ets := request.PathParameters["eventType"]
		et, err := GetEventType(ets)
		fmt.Printf("The event type passed in is \"%s\n\"", et)

		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
		}

		etb, err := json.Marshal(et)

		// Get slots
		if err != nil {

		}

		return events.APIGatewayProxyResponse{Body: string(etb), StatusCode: 200}, nil



	}

	// List possible slots for a specific event type

	// Request a specific time slot for the event
	// Approve/Decline the appointment request

	return events.APIGatewayProxyResponse{Body: "Invalid operation!", StatusCode: 400}, nil
}