package garoon

import (
	"fmt"
	"net/url"
	"time"
)

type Facility struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type FacilityGroup struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Code                string `json:"code"`
	Notes               string `json:"notes"`
	ParentFacilityGroup string `json:"parentFacilityGroup"`
	ChildFacilityGroups []struct {
		ID string `json:"id"`
	} `json:"childFacilityGroups"`
}

type FacilityReservationInfo struct {
	AdditionalProp1 struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"additionalProp1"`
	AdditionalProp2 struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"additionalProp2"`
	AdditionalProp3 struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"additionalProp3"`
}

type FacilityUsageRequest struct {
	Status           string    `json:"status"`
	Facility         Facility  `json:"facility"`
	ApprovedBy       User      `json:"approvedBy"`
	ApprovedDateTime time.Time `json:"approvedDateTime"`
}

func (c *Client) GetFacilities(values url.Values) (*FacilityPager, error) {
	path := fmt.Sprintf("schedule/facilities?%s", values.Encode())
	var pager FacilityPager
	if err := c.fetchResource("GET", path, nil, &pager); err != nil {
		return nil, err
	}
	return &pager, nil
}

func (c *Client) GetFacilityGroups(values url.Values) (*FacilityGroupPager, error) {
	path := fmt.Sprintf("schedule/facilityGroups?%s", values.Encode())
	var pager FacilityGroupPager
	if err := c.fetchResource("GET", path, nil, &pager); err != nil {
		return nil, err
	}
	return &pager, nil
}

func (c *Client) GetFacilitiesByFacilityGroup(facilityGroupID string, values url.Values) (*FacilityPager, error) {
	path := fmt.Sprintf("schedule/facilityGroups/%s/facilities?%s", facilityGroupID, values.Encode())
	var pager FacilityPager
	if err := c.fetchResource("GET", path, nil, &pager); err != nil {
		return nil, err
	}
	return &pager, nil
}
