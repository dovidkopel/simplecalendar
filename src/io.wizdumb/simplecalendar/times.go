package simplecalendar

import (
	"log"
	"time"
)

// Time spans that are disabled across days
// Busy 1PM - 1:30PM (All)
// Available: 7:00AM - 7PM (Mon, Wed, Thurs) -->
//    0000-0829 [Monday, Wednesday, Thursday] BUSY
//    0701-2359 [Monday, Wednesday, Thursday] BUSY
// Available: 6AM - 4PM (Friday) - Early day -->
//    0000-0559 [Friday] BUSY - Pre
//    1601-2359 [Friday] BUSY - Post
// Available: 12PM - 10PM (Tuesday) - Late day
//    0000-1159 [Tuesday] BUSY - Pre
//    2201-2359 [Tuesday] BUSY - Post

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

	label  string
	days   []time.Weekday
	start  int
	end    int
	status Availability
}

var AllDays = []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}

func TimeBlockAllDays(label string, start int, end int, status Availability) TimeBlockSchedule {
	t := TimeBlockSchedule{
		label:  label,
		days:   AllDays,
		start:  start,
		end:    end,
		status: status,
	}

	return t
}

func TimeBlockDays(label string, start int, end int, days []time.Weekday, status Availability) TimeBlockSchedule {
	t := TimeBlockSchedule{
		label:  label,
		days:   days,
		start:  start,
		end:    end,
		status: status,
	}

	return t
}

type TimeSchedule struct {
	SchedulingPolicy

	Times []TimeBlockSchedule
}

func (b *TimeSchedule) IsAvailable(times EventTimes) Availability {
	var as []Availability
	for _, s := range b.Times {
		for _, d := range s.days {
			if times.Start.Weekday() == d || times.End.Weekday() == d {
				as = append(as, s.IsAvailable(times))
			}
		}
	}
	return simplify(as)
}

func (b *TimeBlockSchedule) IsAvailable(times EventTimes) Availability {
	log.Printf("Inputted: Start %d %d, End %d %d", times.Start.Hour(), times.Start.Minute(), times.End.Hour(), times.End.Minute())
	s := (times.Start.Hour() * 100) + times.Start.Minute()
	e := (times.End.Hour() * 100) + times.End.Minute()

	log.Printf("Start: %d >= to %d", s, b.start)
	log.Printf("End: %d <= to %d", e, b.end)

	if s >= b.start && e <= b.end {
		log.Println("Return: " + b.status.name)
		return b.status
	}

	log.Printf("Start: %d >= to %d", e, b.start)
	log.Printf("End: %d <= to %d", s, b.end)
	if e >= b.start && s <= b.end {
		log.Println("Return: " + b.status.name)
		return b.status
	}

	log.Println("Return: " + Unknown.name)
	return Unknown
}
