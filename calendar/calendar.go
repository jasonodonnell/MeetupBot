package calendar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const apiURL = "https://www.googleapis.com/calendar/v3/calendars"

type calendar struct {
	key string
	id  string
}

// Events represents calendar events.
type Events struct {
	Items []struct {
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Location    string `json:"location"`
		Start       struct {
			DateTime string `json:"dateTime"`
		} `json:"start"`
	} `json:"items"`
}

// NewCalendar the calendar struct for getting events.
func NewCalendar(key, id string) *calendar {
	return &calendar{
		key: key,
		id:  id,
	}
}

// UpcomingEvents takes a start/end time and returns a list of events.
func (c *calendar) UpcomingEvents(start, end string) (*Events, error) {
	client := http.Client{Timeout: 5 * time.Second}
	r, err := client.Get(c.formatURL(start, end))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	e := Events{}
	json.NewDecoder(r.Body).Decode(&e)
	return &e, nil
}

func (c *calendar) formatURL(start, end string) string {
	format := "%s/%s/events?timeMin=%s&timeMax=%s&key=%s"
	id := url.QueryEscape(c.id)
	s := url.QueryEscape(start)
	e := url.QueryEscape(end)
	return fmt.Sprintf(format, apiURL, id, s, e, c.key)
}
