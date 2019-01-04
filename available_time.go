package garoon

type AvailableTime struct {
	Start    Time     `json:"start"`
	End      Time     `json:"end"`
	Facility Facility `json:"facility"`
}

type AvailableTimeParameter struct {
	TimeRanges              []DateTimePeriod `json:"timeRanges"`
	TimeInterval            string           `json:"timeInterval"`
	Attendees               []Attendee       `json:"attendees,omitempty"`
	Facilities              []Facility       `json:"facilities,omitempty"`
	FacilitySearchCondition string           `json:"facilitySearchCondition"`
}

func (c *Client) SearchAvailableTimes(param *AvailableTimeParameter) (*AvailableTimePager, error) {
	path := "schedule/searchAvailableTimes"
	var pager AvailableTimePager
	if err := c.fetchResource("POST", path, param, &pager); err != nil {
		return nil, err
	}
	return &pager, nil
}
