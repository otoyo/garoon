package garoon

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestAvailableTimes(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/available_times.json")
	if err != nil {
		t.Fatal()
		return
	}

	httpmock.RegisterResponder("POST", "https://subdomain.cybozu.com/g/api/v1/schedule/searchAvailableTimes",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	a := Attendee{}
	a.ID = "940"
	a.Type = "USER"

	d := DateTimePeriod{}
	d.Start = time.Date(2019, 1, 6, 9, 0, 0, 0, time.Local)
	d.End = time.Date(2019, 1, 6, 10, 30, 0, 0, time.Local)

	f := Facility{}
	f.ID = "45"

	p := AvailableTimeParameter{}
	p.TimeRanges = append(p.TimeRanges, d)
	p.Attendees = append(p.Attendees, a)
	p.Facilities = append(p.Facilities, f)
	p.TimeInterval = "30"
	p.FacilitySearchCondition = "OR"

	_, err = c.SearchAvailableTimes(&p)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}
