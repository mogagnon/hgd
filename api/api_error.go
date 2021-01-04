package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

type Response struct {
	*http.Response
}

func (response *Response) HandleApiError(message string) error {
	if response.StatusCode != 200 {
		var apiError Error
		json.NewDecoder(response.Body).Decode(&apiError)

		return fmt.Errorf("%s: %d %s", message, response.StatusCode, apiError.Message)
	}

	return nil
}
