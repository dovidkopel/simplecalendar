package simplecalendar

import (
	"log"
	"testing"
	"time"
)

func lunch() TimeBlockSchedule {
	return TimeBlockAllDays("lunch", 1300, 1330, Busy)
}

func beforeHours() TimeBlockSchedule {
	ds := []time.Weekday{time.Monday, time.Wednesday, time.Thursday}
	return TimeBlockDays("before hours", 0, 659, ds, Busy)
}

func afterHours() TimeBlockSchedule {
	ds := []time.Weekday{time.Monday, time.Wednesday, time.Thursday}
	return TimeBlockDays("after hours", 1900, 2359, ds, Busy)
}

func prep() TimeSchedule {
	return TimeSchedule{
		Times: []TimeBlockSchedule{lunch(), beforeHours(), afterHours()},
	}
}

type TimeScenario struct {
	label string
	start string
	dur time.Duration
	expect Availability
}

func scenarios() []TimeScenario {
	return []TimeScenario{
		TimeScenario{"monday 4am", "2018-05-07T04:02:20+00:00", 15 * time.Minute, Busy},
		TimeScenario{"monday 6am", "2018-05-07T04:02:20+00:00", 15 * time.Minute, Busy},
		TimeScenario{"monday 7am", "2018-05-07T07:02:20+00:00", 15 * time.Minute, Available},
		TimeScenario{"monday 530pm", "2018-05-07T17:30:20+00:00", 15 * time.Minute, Available},
		TimeScenario{"monday 630pm", "2018-05-07T18:30:20+00:00", 45 * time.Minute, Busy},
		TimeScenario{"monday 705pm", "2018-05-07T19:05:20+00:00", 15 * time.Minute, Busy},
	}
}

func TestTimes(t *testing.T) {
	s := prep()
	Conf.SetTimes(s)

	for _, t := range scenarios() {
		st, _ := time.Parse(time.RFC3339, t.start)
		r := s.IsAvailable(times(st, t.dur))

		if r == t.expect {
			log.Println(t.label+" had the correct expectation: "+t.expect.name)
		} else {
			log.Fatalln(t.label+" had the INCORRECT expectation: "+t.expect.name)
		}
	}
}

func TestBlockAllDays(t *testing.T) {
	st22, _ := time.Parse(time.RFC3339, "2018-05-05T13:02:20+00:00")
	ee := times(st22, 15*time.Minute)


	ss := lunch()
	sr1e := ss.IsAvailable(ee)
	if sr1e != Busy {
		log.Fatal("There should be a break between 1300-1300")
	} else {
		log.Println("There is a break between 1300-1300")
	}

}
