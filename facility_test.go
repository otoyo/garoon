package garoon

import (
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetFacilities(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/facilities.json")
	if err != nil {
		t.Fatal()
		return
	}

	httpmock.RegisterResponder("GET", "https://subdomain.cybozu.com/g/api/v1/schedule/facilities",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	v := url.Values{}

	_, err = c.GetFacilities(v)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}

func TestGetFacilityGroups(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/facility_groups.json")
	if err != nil {
		t.Fatal()
		return
	}

	httpmock.RegisterResponder("GET", "https://subdomain.cybozu.com/g/api/v1/schedule/facilityGroups",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	v := url.Values{}

	_, err = c.GetFacilityGroups(v)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}

func TestGetFacilitiesByFacilityGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/facilities.json")
	if err != nil {
		t.Fatal()
		return
	}

	httpmock.RegisterResponder("GET", "https://subdomain.cybozu.com/g/api/v1/schedule/facilityGroups/1/facilities",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	v := url.Values{}

	_, err = c.GetFacilitiesByFacilityGroup("1", v)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}
