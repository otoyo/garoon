package garoon

import (
	"fmt"
	"net/url"
	"time"
)

type NotificationItem struct {
	ModuleID  string    `json:"moduleId"`
	Creator   User      `json:"creator"`
	CreatedAt time.Time `json:"createdAt"`
	Operation string    `json:"operation"`
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Icon      string    `json:"icon"`
	IsRead    bool      `json:"isRead"`
}

func (c *Client) GetNotificationItems(values url.Values) (*NotificationItemPager, error) {
	path := fmt.Sprintf("notification/items?%s", values.Encode())
	var pager NotificationItemPager
	if err := c.fetchResource("GET", path, nil, &pager); err != nil {
		return nil, err
	}
	return &pager, nil
}
