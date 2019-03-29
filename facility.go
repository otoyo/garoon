package garoon

import (
	"fmt"
	"net/url"
	"time"
)

type Facility struct {
	ID            int64  `json:"id,string,omitempty"`
	Name          string `json:"name,omitempty"`
	Code          string `json:"code,omitempty"`
	Notes         string `json:"notes,omitempty"`
	FacilityGroup string `json:"facilityGroup,omitempty"`
}

type FacilityGroup struct {
	ID                  int64  `json:"id,string"`
	Name                string `json:"name"`
	Code                string `json:"code"`
	Notes               string `json:"notes"`
	ParentFacilityGroup string `json:"parentFacilityGroup"`
	ChildFacilityGroups []struct {
		ID int64 `json:"id,string"`
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

func (c *Client) GetFacilitiesByFacilityGroup(facilityGroupID int64, values url.Values) (*FacilityPager, error) {
	path := fmt.Sprintf("schedule/facilityGroups/%d/facilities?%s", facilityGroupID, values.Encode())
	var pager FacilityPager
	if err := c.fetchResource("GET", path, nil, &pager); err != nil {
		return nil, err
	}
	return &pager, nil
}
