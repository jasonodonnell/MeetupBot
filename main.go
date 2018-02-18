package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const apiURL = "https://www.googleapis.com/calendar/v3/calendars"
const calID = "6l7e832ee9bemt1i9c42vltrug@group.calendar.google.com"

var (
	apiKey = getenv("API_KEY")
)

func getenv(name string) string {
	env := os.Getenv(name)
	if env == "" {
		fmt.Println("Missing environment variable: " + name)
		os.Exit(1)
	}
	return env
}

type Response struct {
	Items []struct {
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Location    string `json:"location"`
		Start       struct {
			DateTime string `json:"dateTime"`
		} `json:"start"`
	} `json:"items"`
}

type Calendar struct {
	APIKey string
	ID     string
	Start  string
	End    string
}

func (c *Calendar) URL() string {
	format := "%s/%s/events?timeMin=%s&timeMax=%s&key=%s"
	id := url.QueryEscape(c.ID)
	start := url.QueryEscape(c.Start)
	end := url.QueryEscape(c.End)
	return fmt.Sprintf(format, apiURL, id, start, end, c.APIKey)
}

func main() {
	start := time.Now()
	end := start.Add(time.Hour * 24 * 6)
	c := Calendar{
		APIKey: apiKey,
		ID:     calID,
		Start:  start.Format(time.RFC3339),
		End:    end.Format(time.RFC3339),
	}

	client := http.Client{Timeout: 5 * time.Second}
	r, err := client.Get(c.URL())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cal := Response{}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&cal)
	fmt.Println("Meetups This Week:")
	for _, event := range cal.Items {
		t, _ := time.Parse(time.RFC3339, event.Start.DateTime)
		meetupTime := t.Format("Mon 3:04PM")
		fmt.Printf("%s: %s :: %s\n", meetupTime, event.Summary, event.Location)
	}
}
