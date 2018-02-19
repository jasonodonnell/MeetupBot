package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jasonodonnell/MeetupBot/calendar"
	"github.com/jasonodonnell/MeetupBot/slack"
)

const apiURL = "https://www.googleapis.com/calendar/v3/calendars"

var (
	apiKey = getenv("API_KEY")
	hook   = getenv("SLACK_WEBHOOK")
	calID  = getenv("CAL_ID")
)

func getenv(name string) string {
	env := os.Getenv(name)
	if env == "" {
		log.Fatalf("Missing environment variable: %s", name)
	}
	return env
}

func main() {
	cal := calendar.NewCalendar(apiKey, calID, apiURL)
	slack := slack.NewClient(hook)

	start := time.Now()
	end := start.Add(time.Hour * 24 * 7)
	c, err := cal.UpcomingEvents(start.Format(time.RFC3339), end.Format(time.RFC3339))
	if err != nil {
		log.Fatalf("Could not retrieve events: %s", err)
	}

	if len(c.Items) > 0 {
		payload := "*Meetups This Week*\n\n"
		for _, event := range c.Items {
			t, _ := time.Parse(time.RFC3339, event.Start.DateTime)
			meetupTime := t.Format("Mon 3:04PM")
			payload += fmt.Sprintf("• _%s_ - %s\n", event.Summary, meetupTime)
		}
		payload += "\n*For more info visit* http://techlancaster.com/meetup"
		slack.Send(payload)
	} else {
		log.Println("No meetups..")
	}
}
