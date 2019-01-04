package garoon

import (
	"fmt"
	"net/url"
	"time"
)

type RepeatInfo struct {
	Type               string           `json:"type"`
	Period             DatePeriod       `json:"period"`
	Time               TimePeriod       `json:"time"`
	TimeZone           string           `json:"timeZone"`
	IsAllDay           bool             `json:"isAllDay,omitempty"`
	IsStartOnly        bool             `json:"isStartOnly,omitempty"`
	DayOfWeek          string           `json:"dayOfWeek,omitempty"`
	DayOfMonth         string           `json:"dayOfMonth,omitempty"`
	ExclusiveDateTimes []DateTimePeriod `json:"exclusiveDateTimes,omitempty"`
}

type TemporaryEventCandidate struct {
	End      Time     `json:"end,omitempty"`
	Start    Time     `json:"start,omitempty"`
	Facility Facility `json:"facility,omitempty"`
}

type AdditionalItem struct {
	Value string `json:"value"`
}

type AdditionalItems struct {
	Item AdditionalItem `json:"item"`
}

type Event struct {
	ID                       string                    `json:"id,omitempty"`
	Creator                  Creator                   `json:"creator,omitempty"`
	CreatedAt                time.Time                 `json:"createdAt,omitempty"`
	Updater                  Updater                   `json:"updater,omitempty"`
	UpdatedAt                time.Time                 `json:"updatedAt,omitempty"`
	EventType                string                    `json:"eventType"`
	EventMenu                string                    `json:"eventMenu,omitempty"`
	Subject                  string                    `json:"subject,omitempty"`
	Notes                    string                    `json:"notes,omitempty"`
	VisibilityType           string                    `json:"visibilityType,omitempty"`
	UseAttendanceCheck       bool                      `json:"useAttendanceCheck,omitempty"`
	CompanyInfo              CompanyInfo               `json:"companyInfo,omitempty"`
	Attachments              []Attachment              `json:"attachments,omitempty"`
	Start                    Time                      `json:"start"`
	End                      Time                      `json:"end,omitempty"`
	IsAllDay                 bool                      `json:"isAllDay,omitempty"`
	IsStartOnly              bool                      `json:"isStartOnly,omitempty"`
	OriginalStartTimeZone    string                    `json:"originalStartTimeZone,omitempty"`
	OriginalEndTimeZone      string                    `json:"originalEndTimeZone,omitempty"`
	Attendees                []Attendee                `json:"attendees,omitempty"`
	Watchers                 []Watcher                 `json:"watchers,omitempty"`
	Facilities               []Facility                `json:"facilities,omitempty"`
	FacilityUsingPurpose     string                    `json:"facilityUsingPurpose,omitempty"`
	FacilityReservationInfo  FacilityReservationInfo   `json:"facilityReservationInfo,omitempty"`
	FacilityUsageRequests    []FacilityUsageRequest    `json:"facilityUsageRequests,omitempty"`
	RepeatInfo               RepeatInfo                `json:"repeatInfo,omitempty"`
	TemporaryEventCandidates []TemporaryEventCandidate `json:"temporaryEventCandidates,omitempty"`
	AdditionalItems          AdditionalItems           `json:"additionalItems,omitempty"`
	RepeatID                 string                    `json:"repeatId,omitempty"`
}

func (c *Client) FindEvent(id string) (*Event, error) {
	path := fmt.Sprintf("schedule/events/%s", id)
	var event Event
	if err := c.fetchResource("GET", path, nil, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func (c *Client) SearchEvents(values url.Values) (*EventPager, error) {
	path := fmt.Sprintf("schedule/events?%s", values.Encode())
	var pager EventPager
	if err := c.fetchResource("GET", path, nil, &pager); err != nil {
		return nil, err
	}
	return &pager, nil
}

func (c *Client) CreateEvent(event *Event) (*Event, error) {
	path := "schedule/events"
	var newEvent Event
	if err := c.fetchResource("POST", path, event, &newEvent); err != nil {
		return nil, err
	}
	return &newEvent, nil
}

func (c *Client) UpdateEvent(event *Event) (*Event, error) {
	path := fmt.Sprintf("schedule/events/%s", event.ID)
	var newEvent Event
	if err := c.fetchResource("PATCH", path, event, &newEvent); err != nil {
		return nil, err
	}
	return &newEvent, nil
}

func (c *Client) DeleteEvent(id string) error {
	path := fmt.Sprintf("schedule/events/%s", id)
	if err := c.fetchResource("DELETE", path, nil, nil); err != nil {
		return err
	}
	return nil
}
