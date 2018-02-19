package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type slack struct {
	webhook string
}

type payload struct {
	Username string `json:"username"`
	Icon     string `json:"icon_emoji"`
	Text     string `json:"text"`
}

// NewClient returns the slack struct for sending notifications.
func NewClient(hook string) *slack {
	return &slack{
		webhook: hook,
	}
}

// Send takes a body string and sends it to the configured webhook.
func (s *slack) Send(body string) (err error) {
	p := payload{
		Text: body,
	}

	pJSON, err := json.Marshal(p)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", s.webhook, bytes.NewBuffer(pJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}
