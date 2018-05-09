package simplecalendar

import (
	"log"
	"testing"
	"time"
	"fmt"
	"math"
)

func lunch() TimeBlockSchedule {
	return TimeBlockAllDays("lunch", 1300, 1330, Busy)
}

func beforeHoursMWTH() TimeBlockSchedule {
	ds := []time.Weekday{time.Monday, time.Wednesday, time.Thursday}
	return TimeBlockDays("monday, wednesday, thursday: before hours", 0, 659, ds, Busy)
}

func afterHoursMWTH() TimeBlockSchedule {
	ds := []time.Weekday{time.Monday, time.Wednesday, time.Thursday}
	return TimeBlockDays("monday, wednesday, thursday: after hours", 1900, 2359, ds, Busy)
}

func beforeHoursTuesday() TimeBlockSchedule {
	return TimeBlockDays("tuesday before hours", 0, 1159, []time.Weekday{time.Tuesday}, Busy)
}

func afterHoursTuesday() TimeBlockSchedule {
	return TimeBlockDays("tuesday after hours", 2201, 2359, []time.Weekday{time.Tuesday}, Busy)
}

func beforeHoursFriday() TimeBlockSchedule {
	return TimeBlockDays("friday before hours", 0, 559, []time.Weekday{time.Friday}, Busy)
}

func afterHoursFriday() TimeBlockSchedule {
	return TimeBlockDays("friday after hours", 1601, 2359, []time.Weekday{time.Friday}, Busy)
}

func prep() TimeSchedule {
	return TimeSchedule{
		Times: []TimeBlockSchedule{
			lunch(),
			beforeHoursMWTH(),
			afterHoursMWTH(),
			beforeHoursTuesday(),
			afterHoursTuesday(),
			beforeHoursFriday(),
			afterHoursFriday()},
	}
}

type TimeScenario struct {
	label  string
	start  string
	dur    time.Duration
	expect Availability
}

type PartScenario struct {
	hour int
	minute int
	dur time.Duration
	expect Availability
}

// 2018-05-06 = Sunday 0
// 2018-05-07 = Monday 1
// 2018-05-08 = Tuesday 2
// 2018-05-09 = Wednesday 3
// 2018-05-10 = Thursday 4
// 2018-05-11 = Friday 5
// 2018-05-12 = Saturday 6
var weekdays = map[time.Weekday]int {
	time.Sunday: 0,
	time.Monday: 1,
	time.Tuesday: 2,
	time.Wednesday: 3,
	time.Thursday: 4,
	time.Friday: 5,
	time.Saturday: 6,
}

func makeTime(weekday time.Weekday, h int, m int) time.Time {
	dd := time.Now()
	w := dd.Weekday()
	wi := weekdays[w]
	ww := weekdays[weekday]

	diff := int(math.Abs(float64(int(ww - wi))))
	//log.Printf("wi: %d, ww: %d, diff: %d\n", wi, ww, diff)
	if wi > ww {
		diff = diff * -1
	}
	dd = dd.AddDate(0, 0, diff)
	if dd.Weekday() != weekday {
		//log.Fatalf("The day of the week doesn't match %s %s\n", weekday.String(), dd.Weekday().String())
	}
	return time.Date(dd.Year(), dd.Month(), dd.Day(), h, m, 0, 0, time.UTC)
}

func monWedThurScenarios() []TimeScenario {
	var ts []TimeScenario
	for _, w := range []time.Weekday{time.Monday, time.Wednesday, time.Thursday} {
		for _, t := range []PartScenario{
			PartScenario{4, 0, 15*time.Minute, Busy},
			PartScenario{6, 0, 15*time.Minute, Busy},
			PartScenario{7, 0, 15*time.Minute, Available},
			PartScenario{17, 30, 15*time.Minute, Available},
			PartScenario{18, 30, 45*time.Minute, Busy},
			PartScenario{19, 05, 15*time.Minute, Busy},
		} {
			m := makeTime(w, t.hour, t.minute)
			ddt := m.Format(time.RFC3339)
			label := fmt.Sprintf("%s %s %d %d", w.String(), ddt, t.hour, t.minute)
			ts = append(ts, TimeScenario{label, ddt, t.dur, t.expect})
		}

	}

	return ts
}

func tuesdayScenarios() []TimeScenario {
	var ts []TimeScenario
	for _, t := range []PartScenario{
		PartScenario{11, 0, 15*time.Minute, Busy},
		PartScenario{12, 15, 15*time.Minute, Available},
		PartScenario{20, 0, 15*time.Minute, Available},
		PartScenario{22, 30, 15*time.Minute, Busy},
	} {
		m := makeTime(time.Tuesday, t.hour, t.minute)
		ddt := m.Format(time.RFC3339)
		label := fmt.Sprintf("%s %s %d %d", time.Tuesday.String(), ddt, t.hour, t.minute)
		ts = append(ts, TimeScenario{label, ddt, t.dur, t.expect})
	}

	return ts
}

func fridayScenarios() []TimeScenario {
	var ts []TimeScenario
	for _, t := range []PartScenario{
		PartScenario{4, 0, 15*time.Minute, Busy},
		PartScenario{6, 15, 15*time.Minute, Available},
		PartScenario{15, 0, 15*time.Minute, Available},
		PartScenario{17, 30, 15*time.Minute, Busy},
	} {
		m := makeTime(time.Friday, t.hour, t.minute)
		ddt := m.Format(time.RFC3339)
		label := fmt.Sprintf("%s %s %d %d", time.Friday.String(), ddt, t.hour, t.minute)
		ts = append(ts, TimeScenario{label, ddt, t.dur, t.expect})
	}

	return ts
}

func testScenario(s TimeSchedule, t TimeScenario) {
	log.Printf("Label: %s %s -> %s\n", t.label, t.start, t.expect.name)
	st, _ := time.Parse(time.RFC3339, t.start)
	r := s.IsAvailable(times(st, t.dur))

	if r == t.expect {
		log.Println(t.label + " had the correct expectation: " + t.expect.name)
	} else {
		log.Fatalln(t.label + " had the INCORRECT expectation: " + t.expect.name)
	}
}

func TestTimes(t *testing.T) {
	s := prep()
	Conf.SetTimes(s)

	//ss := scenarios()
	ss := append(monWedThurScenarios(), fridayScenarios()...)
	ss = append(ss, tuesdayScenarios()...)

	for _, t := range ss {
		testScenario(s, t)
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
