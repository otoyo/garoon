package garoon

import (
	"fmt"
	"net/url"
)

type User struct {
	ID   string `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}

type Creator User
type Updater User
type Watcher User

type AttendanceResponse struct {
	Status  string `json:"status"`
	Comment string `json:"comment,omitempty"`
}

type Attendee struct {
	User
	AttendanceResponse AttendanceResponse `json:"attendanceResponse,omitempty"`
}

func (c *Client) SearchUsers(values url.Values) (*UserPager, error) {
	path := fmt.Sprintf("base/users?%s", values.Encode())
	var pager UserPager
	if err := c.fetchResource("GET", path, nil, &pager); err != nil {
		return nil, err
	}
	return &pager, nil
}

func (c *Client) GetUsersByOrganization(organizationID string, values url.Values) (*UserPager, error) {
	path := fmt.Sprintf("base/organizations/%s/users?%s", organizationID, values.Encode())
	var pager UserPager
	if err := c.fetchResource("GET", path, nil, &pager); err != nil {
		return nil, err
	}
	return &pager, nil
}
