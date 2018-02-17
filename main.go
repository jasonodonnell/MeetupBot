package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		Start:  string(start.Format(time.RFC3339)),
		End:    string(end.Format(time.RFC3339)),
	}

	client := http.Client{Timeout: 5 * time.Second}
	r, err := client.Get(c.URL())
	if err != nil {
		log.Panic(err)
	}

	cal := Response{}

	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&cal)
	for _, event := range cal.Items {
		fmt.Println("Meetups This Week:")
		fmt.Printf("%s :: %s\n", event.Summary, event.Location)
		// fmt.Println(event.Description)
		// fmt.Println(event.Location)
		// t, _ := time.Parse(time.RFC3339, event.End.DateTime)
		// local := t.Local()
		// fmt.Println(local)
	}

}
