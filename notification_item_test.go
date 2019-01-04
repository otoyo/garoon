package garoon

import (
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetNotificationItems(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	b, err := ioutil.ReadFile("testdata/notification_items.json")
	if err != nil {
		t.Fatal()
		return
	}

	httpmock.RegisterResponder("GET", "https://subdomain.cybozu.com/g/api/v1/notification/items",
		httpmock.NewStringResponder(200, string(b)))

	c, err := NewClient("subdomain", "user", "password")

	v := url.Values{}

	_, err = c.GetNotificationItems(v)
	if err != nil {
		t.Errorf("Expected success, but error: %s", err)
		return
	}
}
