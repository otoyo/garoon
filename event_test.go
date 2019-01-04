package garoon

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestFindEvent(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/event.json")
	if err != nil {
		t.Fatal()
		return
	}

	var event Event
	if err = json.Unmarshal(b, &event); err != nil {
		t.Fatal()
		return
	}

	httpmock.RegisterResponder("GET", "https://subdomain.cybozu.com/g/api/v1/schedule/events/123",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	_, err = c.FindEvent(event.ID)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}

func TestSearchEvents(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/events.json")
	if err != nil {
		t.Fatal()
		return
	}

	v := url.Values{}

	httpmock.RegisterResponder("GET", "https://subdomain.cybozu.com/g/api/v1/schedule/events",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	_, err = c.SearchEvents(v)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}

func TestCreateEvent(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/event.json")
	if err != nil {
		t.Fatal()
		return
	}

	var event Event
	if err = json.Unmarshal(b, &event); err != nil {
		t.Fatal()
		return
	}

	httpmock.RegisterResponder("POST", "https://subdomain.cybozu.com/g/api/v1/schedule/events",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	_, err = c.CreateEvent(&event)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}

func TestUpdateEvent(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/event.json")
	if err != nil {
		t.Fatal()
		return
	}

	var event Event
	if err = json.Unmarshal(b, &event); err != nil {
		t.Fatal()
		return
	}

	httpmock.RegisterResponder("PATCH", "https://subdomain.cybozu.com/g/api/v1/schedule/events/123",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	_, err = c.UpdateEvent(&event)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}

func TestDeleteEvent(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "https://subdomain.cybozu.com/g/api/v1/schedule/events/123",
		httpmock.NewStringResponder(200, string("")))

	c, err := NewClient("subdomain", "user", "password")

	if err = c.DeleteEvent("123"); err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}
