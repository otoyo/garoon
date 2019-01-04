package garoon

type UserPager struct {
	Users   []User `json:users`
	HasNext bool   `json:"hasNext"`
}

type OrganizationPager struct {
	Organizations []Organization `json:organizations`
	HasNext       bool           `json:"hasNext"`
}

type EventPager struct {
	Events  []Event `json:events`
	HasNext bool    `json:"hasNext"`
}

type AvailableTimePager struct {
	AvailableTimes []AvailableTime `json:"availableTimes"`
}

type FacilityPager struct {
	Facilities []Facility `json:"facilities"`
	HasNext    bool       `json:"hasNext"`
}

type FacilityGroupPager struct {
	FacilityGroups []FacilityGroup `json:"facilityGroups"`
	HasNext        bool            `json:"hasNext"`
}

type NotificationItemPager struct {
	NotificationItems []NotificationItem `json:"items"`
	HasNext           bool               `json:"hasNext"`
}
