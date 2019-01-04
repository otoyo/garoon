package garoon

import (
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestSearchUsers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/users.json")
	if err != nil {
		t.Fatal()
		return
	}

	httpmock.RegisterResponder("GET", "https://subdomain.cybozu.com/g/api/v1/base/users",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	v := url.Values{}

	_, err = c.SearchUsers(v)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}

func TestGetUsersByOrganization(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/users.json")
	if err != nil {
		t.Fatal()
		return
	}

	httpmock.RegisterResponder("GET", "https://subdomain.cybozu.com/g/api/v1/base/organizations/1/users",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	v := url.Values{}

	_, err = c.GetUsersByOrganization("1", v)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}
