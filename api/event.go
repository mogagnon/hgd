package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Event struct {
	What string `json:"what"`
	When int64  `json:"when"`
	Tags string `json:"tags"`
}

func (client *HgClient) CreateEvent(event Event) error {
	url := client.GetEvent()
	var requestBody, _ = json.Marshal(event)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	defer response.Body.Close()

	apiResponse := Response{response}

	err = apiResponse.HandleApiError(fmt.Sprintf("error creating event %s", event.What))

	if err != nil {
		return fmt.Errorf("create event %s: %v", event.What, err)
	}

	return nil
}
