package garoon

import (
	"fmt"
	"net/url"
)

type Organization struct {
	ID                 int64  `json:"id,string"`
	Name               string `json:"name"`
	Code               string `json:"code"`
	ParentOrganization string `json:"parentOrganization"`
	ChildOrganizations []struct {
		ID int64 `json:"id,string"`
	} `json:"childOrganizations"`
}

func (c *Client) GetOrganizations(values url.Values) (*OrganizationPager, error) {
	path := fmt.Sprintf("base/organizations?%s", values.Encode())
	var pager OrganizationPager
	if err := c.fetchResource("GET", path, nil, &pager); err != nil {
		return nil, err
	}
	return &pager, nil
}
