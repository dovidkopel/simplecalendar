package simplecalendar

import (
	"time"
	"log"
)

// Time spans that are disabled across days
// Busy 1PM - 1:30PM (All)
// Available: 8:30AM - 7PM (Mon, Wed, Thurs) -->
//    0000-0829 [Sunday, Tuesday, Wednesday] BUSY
//    0701-2359 [Monday, Wednesday, Thursday] BUSY
// Available: 6AM - 5PM (Friday)
// Available: 10AM - 10PM (Tuesday)

// A time that is not bound to a day
// The only question to ask is whether a
// given time range is within the time boundary
// or not.

// 1PM-1:30PM shall be represented as:
// 1300 - 1330
type ContainsTime interface {
	Contains(st time.Time, et time.Time) bool
}

type TimeBlockSchedule struct {
	SchedulingPolicy
	ContainsTime

	days []time.Weekday
	start int
	end int
	status Availability
}

func TimeBlockAllDays(start int, end int, status Availability) TimeBlockSchedule {
	return TimeBlockSchedule{
		days: []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday},
		start: start,
		end: end,
		status: status,
	}
}

func TimeBlockDays(start int, end int, days []time.Weekday, status Availability) TimeBlockSchedule {
	return TimeBlockSchedule{
		days: days,
		start: start,
		end: end,
		status: status,
	}
}

type TimeSchedule struct {
	SchedulingPolicy

	Times []TimeBlockSchedule
}

func (b *TimeBlockSchedule) IsAvailable(times EventTimes) Availability {
	log.Printf("Start %d %d, End %d %d", times.Start.Hour(), times.Start.Minute(), times.End.Hour(), times.End.Minute())
	s := (times.Start.Hour() * 100) + times.Start.Minute()
	e := (times.End.Hour() * 100) + times.End.Minute()

	log.Printf("Start: %d compared to %d", s, b.start)
	log.Printf("End: %d compared to %d", e, b.end)


	if s >= b.start && e <= b.end {
		log.Println("BUSY")
		return Busy
	}
	return Unknown
}